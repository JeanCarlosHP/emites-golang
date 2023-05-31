package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/jeancarloshp/emites/domain"
	httpService "github.com/jeancarloshp/emites/internal/http_service"
	organizationService "github.com/jeancarloshp/emites/service"
)

func StartIntegration(organizationService organizationService.OrganizationService, httpService httpService.HttpService) error {
	log.Println("Integração iniciada!")
	nfes, err := organizationService.OUseCase.GetAllNfes(httpService)
	if err != nil {
		return err
	}

	jsonToSap, err := convertBodyToSapFormat(nfes)
	if err != nil {
		log.Fatal(err)
	}

	err = sendToSap(jsonToSap, httpService)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Finalizado!")

	return nil
}

func convertBodyToSapFormat(nfes []domain.Nfe) ([]byte, error) {
	jsonToSap, err := json.Marshal(nfes)
	if err != nil {
		panic("Error during marshaling - " + err.Error())
	}
	jsonToSap = []byte(fmt.Sprintf(`{"I_NFE": %s}`, jsonToSap))

	return jsonToSap, nil
}

func sendToSap(body []byte, httpService httpService.HttpService) error {
	log.Println("Enviado para o SAP: " + string(body))

	httpClient, err := httpService.NewHttpClient()
	if err != nil {
		log.Fatalf("Error during the http client creation.\n%s", err)
	}

	sapEmitesUrl := os.Getenv("sap.url.base") + os.Getenv("sap.emites.function")
	sapAuthorization := "Basic " + os.Getenv("sap.authorization")

	request, err := httpService.Post(sapEmitesUrl, bytes.NewReader(body))
	if err != nil {
		panic("Erro ao criar a request.")
	}
	request.Header.Set("Authorization", sapAuthorization)

	response, err := httpClient.Do(request)
	if err != nil {
		panic("Erro ao executar a request.\n" + err.Error())
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	if response.StatusCode != 200 {
		panic("Erro na requisição - " + strconv.Itoa(response.StatusCode) + " - " + sapEmitesUrl + " - Retorno: " + string(responseBody))
	}

	log.Println("Retorno do SAP: ", string(responseBody))

	return nil
}

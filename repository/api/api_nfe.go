package api

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/jeancarloshp/emites/domain"
	httpService "github.com/jeancarloshp/emites/internal/http_service"
	nfe "github.com/jeancarloshp/emites/model"
)

type apiNfeRepository struct {
	httpService *httpService.HttpService
}

func NewApiNfeRepository(httpService *httpService.HttpService) domain.NfeRepository {
	return &apiNfeRepository{
		httpService,
	}
}

func (a *apiNfeRepository) GetAllNfeByOrganizationId(organizationId int, httpService httpService.HttpService) ([]domain.Nfe, error) {
	client, err := httpService.NewHttpClient()
	if err != nil {
		panic("Não foi possível criar o Http Client.")
	}

	yesterdayDate := time.Now().AddDate(0, 0, -1)

	layout := "2006/01/02"
	yesterdayDateFormatted := yesterdayDate.Format(layout)

	emitesUrl := os.Getenv("emites.url.base") + "organizations/" + strconv.Itoa(organizationId) + "/nfe?starting_at=" + yesterdayDateFormatted
	emitesToken := "Bearer " + os.Getenv("emites.token")

	nfeList := nfe.NfeJsonList{}
	nfeListFormatted := []domain.Nfe{}
	page := 1

	// Do While
	for ok := true; ok; ok = (len(nfeList.Nfe) != 0) {
		nfeList, err = fetchNfe(httpService, client, emitesUrl, emitesToken, page, &nfeListFormatted)
		page++
	}

	return nfeListFormatted, err
}

func fetchNfe(httpService httpService.HttpService, client *http.Client, emitesUrl string, emitesToken string, page int, nfeListFormatted *[]domain.Nfe) (nfe.NfeJsonList, error) {
	request, err := httpService.Get(emitesUrl + "&page=" + fmt.Sprint(page))
	if err != nil {
		panic("Erro ao criar a request.")
	}
	request.Header.Set("Authorization", emitesToken)

	response, err := client.Do(request)
	if err != nil {
		panic("Erro ao executar a request: " + err.Error())
	}
	if response.StatusCode != 200 {
		log.Fatal("Erro na requisição - " + strconv.Itoa(response.StatusCode) + " - " + emitesUrl)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		panic("Erro ao ler dados do payload.")
	}

	nfeList := nfe.NfeJsonList{}

	err = json.Unmarshal(data, &nfeList)
	if len(nfeList.Nfe) == 0 {
		return nfeList, err
	} else {
		layout := "2006-01-02"

		for _, s := range nfeList.Nfe {
			if s.Data.Xml != "" {
				nfeXml := nfe.Xml{}
				err := xml.Unmarshal([]byte(s.Data.Xml), &nfeXml)
				if err != nil {
					panic(err.Error())
				}

				*nfeListFormatted = append(*nfeListFormatted, domain.Nfe{
					ChaveNfe: nfeXml.ProtNFe.InfProt.ChaveNfe,
					ValorNfe: nfeXml.NFe.InfNFe.Total.ICMSTot.ValorNfe,
					Cfop:     nfeXml.NFe.InfNFe.Det.Prod.Cfop,
					DataNfe:  nfeXml.NFe.InfNFe.Ide.DataEmissao.Format(layout),
				})
			}
		}
	}

	return nfeList, err
}

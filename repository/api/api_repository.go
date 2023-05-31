package api

import (
	"encoding/json"
	"io"
	"os"
	"strconv"

	"github.com/jeancarloshp/emites/domain"
	httpService "github.com/jeancarloshp/emites/internal/http_service"
)

type apiOrganizationRepository struct {
	httpService *httpService.HttpService
}

func NewApiOrganizationRepository(httpService *httpService.HttpService) domain.OrganizationRepository {
	return &apiOrganizationRepository{
		httpService,
	}
}

func (a *apiOrganizationRepository) GetAllOrganizations(httpService httpService.HttpService) ([]domain.Organization, error) {
	client, err := httpService.NewHttpClient()
	if err != nil {
		panic("Não foi possível criar o Http Client.")
	}

	emitesOrganizationsUrl := os.Getenv("emites.url.base") + "organizations"
	emitesToken := "Bearer " + os.Getenv("emites.token")

	request, err := httpService.Get(emitesOrganizationsUrl)
	if err != nil {
		panic("Erro ao criar a request.")
	}
	request.Header.Set("Authorization", emitesToken)

	response, err := client.Do(request)
	if err != nil {
		panic("Erro ao executar a request.")
	}
	if response.StatusCode != 200 {
		panic("Erro na requisição - " + strconv.Itoa(response.StatusCode) + " - " + emitesOrganizationsUrl)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		panic("Erro ao ler dados do payload.")
	}

	organizations := domain.Organizations{}
	err = json.Unmarshal(data, &organizations)

	return organizations.Organization, err
}

func (a *apiOrganizationRepository) GetAllNfes(httpService httpService.HttpService) ([]domain.Nfe, error) {
	return []domain.Nfe{}, nil
}

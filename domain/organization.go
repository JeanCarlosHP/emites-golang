package domain

import (
	httpService "github.com/jeancarloshp/emites/internal/http_service"
)

type (
	Organizations struct {
		Organization []Organization `json:"organizations"`
	}

	Organization struct {
		ID int `json:"id"`
	}

	OrganizationUseCase interface {
		GetAllOrganizations(httpService httpService.HttpService) ([]Organization, error)
		GetAllNfes(httpService httpService.HttpService) ([]Nfe, error)
	}

	OrganizationRepository interface {
		GetAllOrganizations(httpService httpService.HttpService) ([]Organization, error)
		GetAllNfes(httpService httpService.HttpService) ([]Nfe, error)
	}
)

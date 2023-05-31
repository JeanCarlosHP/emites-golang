package domain

import httpService "github.com/jeancarloshp/emites/internal/http_service"

type (
	NfeList struct {
		NfeList []Nfe `json:"nfe"`
	}

	Nfe struct {
		ChaveNfe string `json:"CHAVE_NFE"`
		ValorNfe string `json:"VALOR_NFE"`
		Cfop     string `json:"CFOP"`
		DataNfe  string `json:"DATA_NFE"`
	}

	NfeRepository interface {
		GetAllNfeByOrganizationId(organizationId int, httpService httpService.HttpService) ([]Nfe, error)
	}
)

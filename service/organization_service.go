package http

import (
	"github.com/jeancarloshp/emites/domain"
	httpService "github.com/jeancarloshp/emites/internal/http_service"
)

type OrganizationService struct {
	OUseCase domain.OrganizationUseCase
}

func NewOrganizationService(ou domain.OrganizationUseCase) *OrganizationService {
	return &OrganizationService{
		OUseCase: ou,
	}
}

func (o *OrganizationService) GetAllNfes(httpCliente httpService.HttpService) ([]domain.Nfe, error) {
	return nil, nil
}

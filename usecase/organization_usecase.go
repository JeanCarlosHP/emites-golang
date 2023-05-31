package usecase

import (
	"github.com/jeancarloshp/emites/domain"
	httpservice "github.com/jeancarloshp/emites/internal/http_service"
)

type organizationUseCase struct {
	organizationRepo domain.OrganizationRepository
	nfeRepo          domain.NfeRepository
}

func NewOrganizationUseCase(o domain.OrganizationRepository, n domain.NfeRepository) domain.OrganizationUseCase {
	return &organizationUseCase{
		organizationRepo: o,
		nfeRepo:          n,
	}
}

func (o *organizationUseCase) GetAllOrganizations(httpService httpservice.HttpService) ([]domain.Organization, error) {
	return o.organizationRepo.GetAllOrganizations(httpService)
}

func (o *organizationUseCase) GetAllNfes(httpService httpservice.HttpService) ([]domain.Nfe, error) {
	organizations, err := o.organizationRepo.GetAllOrganizations(httpService)
	if err != nil {
		panic(err)
	}

	nfeList := []domain.Nfe{}

	for _, s := range organizations {
		nfeByOrganizations, err := o.nfeRepo.GetAllNfeByOrganizationId(s.ID, httpService)
		if err != nil {
			panic(err)
		}

		nfeList = append(nfeList, nfeByOrganizations...)
	}

	return nfeList, nil
}

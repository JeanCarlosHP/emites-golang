package usecase

import (
	"testing"

	"github.com/jeancarloshp/emites/domain"
	"github.com/jeancarloshp/emites/domain/mocks"
	httpService "github.com/jeancarloshp/emites/internal/http_service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAllOrganizations(t *testing.T) {
	mockOrganizationRepo := new(mocks.MockOrganizationRepository)
	mockOrganization := domain.Organization{
		ID: 1,
	}

	mockOrganizationList := make([]domain.Organization, 0)
	mockOrganizationList = append(mockOrganizationList, mockOrganization)

	t.Run("success", func(t *testing.T) {
		mockOrganizationRepo.On("GetAllOrganizations", mock.Anything).Return(mockOrganizationList, nil).Once()

		mockNfe := domain.Nfe{
			ChaveNfe: "123456789",
			ValorNfe: "10.0",
			Cfop:     "1234",
			DataNfe:  "2023-05-31",
		}

		mockNfeList := make([]domain.Nfe, 0)
		mockNfeList = append(mockNfeList, mockNfe)

		mockNfeRepo := new(mocks.MockNfeRepository)
		mockNfeRepo.On("GetAllNfeByOrganizationId", mock.AnythingOfType("int"), mock.Anything).Return(mockNfeList, nil)
		u := NewOrganizationUseCase(mockOrganizationRepo, mockNfeRepo)
		httpService := httpService.NewHttpService()

		nfeList, err := u.GetAllNfes(*httpService)
		assert.Len(t, nfeList, len(mockNfeList))
		assert.NoError(t, err)

		mockOrganizationRepo.AssertExpectations(t)
		mockNfeRepo.AssertExpectations(t)
	})
}

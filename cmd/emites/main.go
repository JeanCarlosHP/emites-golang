package main

import (
	"os"

	"github.com/jeancarloshp/emites/internal/dotenv"
	httpservice "github.com/jeancarloshp/emites/internal/http_service"
	"github.com/jeancarloshp/emites/internal/integration"

	_repository "github.com/jeancarloshp/emites/repository/api"
	_service "github.com/jeancarloshp/emites/service"
	_useCase "github.com/jeancarloshp/emites/usecase"
)

func main() {
	httpService := httpservice.NewHttpService()

	organizationRepository := _repository.NewApiOrganizationRepository(httpService)
	nfeRepository := _repository.NewApiNfeRepository(httpService)

	organizationUseCase := _useCase.NewOrganizationUseCase(organizationRepository, nfeRepository)
	organizationService := _service.NewOrganizationService(organizationUseCase)

	err := integration.StartIntegration(*organizationService, *httpService)
	if err != nil {
		panic("Erro durante a execução da integração.\n" + err.Error())
	}
}

func init() {
	goDotEnv := dotenv.NewGoDotEnv()
	err := goDotEnv.Load()
	if err != nil {
		panic("File .env not loaded.\n" + err.Error())
	}

	env := os.Getenv("emites.url.base")
	if env == "" {
		panic("Environment variable 'emites.url.base' not found.")
	}

	env = os.Getenv("emites.token")
	if env == "" {
		panic("Environment variable 'emites.token' not found.")
	}

	env = os.Getenv("sap.authorization")
	if env == "" {
		panic("Environment variable 'sap.authorization' not found.")
	}

	env = os.Getenv("sap.url.base")
	if env == "" {
		panic("Environment variable 'sap.url.base' not found.")
	}

	env = os.Getenv("sap.emites.function")
	if env == "" {
		panic("Environment variable 'sap.emites.function' not found.")
	}
}

package httpservice

import (
	"io"
	"net/http"
)

type HttpService struct {
	HttpController
}

func NewHttpService() *HttpService {
	return &HttpService{}
}

func (h *HttpService) NewHttpClient() (httpClient *http.Client, err error) {
	return &http.Client{}, nil
}

func (h *HttpService) Post(url string, payload io.Reader) (httpRequest *http.Request, err error) {
	return http.NewRequest(http.MethodPost, url, payload)
}

func (h *HttpService) Get(url string) (httpRequest *http.Request, err error) {
	return http.NewRequest(http.MethodGet, url, nil)
}

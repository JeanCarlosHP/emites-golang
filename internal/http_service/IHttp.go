package httpservice

type HttpController interface {
	NewHttpClient()

	Get()
	Post()

	SetHeaders()
	SetBody()
}

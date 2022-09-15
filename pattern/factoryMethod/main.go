package factoryMethod

type Response interface {
	Status() int
}

type NotFoundResponse struct {
	status int
	body   string
}

func NewNotFoundResponse() *NotFoundResponse {
	return &NotFoundResponse{status: 404}
}

func (nfr *NotFoundResponse) Status() int {
	return nfr.status
}

type NoContentResponse struct {
	status int
}

func NewNoContentResponse() *NoContentResponse {
	return &NoContentResponse{status: 204}
}

func (ncr *NoContentResponse) Status() int {
	return ncr.status
}

// Factory
type ResponseFactory interface {
	NewResponse() Response
}

type NotFoundResponseFactory struct {
	resp *NotFoundResponse
}

func NewNotFoundResponseFactory() *NotFoundResponseFactory {
	return &NotFoundResponseFactory{resp: NewNotFoundResponse()}
}

func (nfrf *NotFoundResponseFactory) NewResponse() Response {
	return nfrf.resp
}

type NoContentResponseFactory struct {
	resp *NoContentResponse
}

func NewNoContentResponseFactory() *NoContentResponseFactory {
	return &NoContentResponseFactory{resp: NewNoContentResponse()}
}

func (ncrf *NoContentResponseFactory) NewResponse() Response {
	return ncrf.resp
}

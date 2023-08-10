package responses

import (
	"go-tdd-clean/12/shared"
	"net/http"
	"reflect"
)

type Response struct {
	Errors     []string `json:"errors"`
	Data       any      `json:"data"`
	StatusCode int      `json:"-"`
}

func BuildResponse(output shared.Output, r *http.Request) Response {
	return Response{
		Errors:     output.GetErrors(),
		Data:       output.GetData(),
		StatusCode: buildHTTPStatusCode(output, r.Method),
	}
}

func buildHTTPStatusCode(output shared.Output, verb string) int {
	domainCode := output.GetCode()
	switch domainCode {
	case shared.DomainCodeSuccess:
		if verb == http.MethodPost {
			return http.StatusCreated
		}
		if verb == http.MethodGet && (output.GetData() == nil || reflect.ValueOf(output.GetData()).IsNil()) {
			return http.StatusNoContent
		}
		return http.StatusOK
	case shared.DomainCodeInvalidInput:
		return http.StatusBadRequest
	case shared.DomainCodeInvalidEntity:
		return http.StatusInternalServerError
	case shared.DomainCodeInternalError:
		return http.StatusInternalServerError
	case shared.DomainCodeNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}

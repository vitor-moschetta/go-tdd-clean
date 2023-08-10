package shared

import (
	"net/http"
	"reflect"
)

type Response struct {
	Errors     []string `json:"errors"`
	Data       any      `json:"data"`
	StatusCode int      `json:"-"`
}

func BuildResponse(output Output, r *http.Request) Response {
	return Response{
		Errors:     output.GetErrors(),
		Data:       output.GetData(),
		StatusCode: buildHTTPStatusCode(output, r.Method),
	}
}

func buildHTTPStatusCode(output Output, verb string) int {
	domainCode := output.GetCode()
	switch domainCode {
	case DomainCodeSuccess:
		if verb == http.MethodPost {
			return http.StatusCreated
		}
		if verb == http.MethodGet && (output.GetData() == nil || reflect.ValueOf(output.GetData()).IsNil()) {
			return http.StatusNoContent
		}
		return http.StatusOK
	case DomainCodeInvalidInput:
		return http.StatusBadRequest
	case DomainCodeInvalidEntity:
		return http.StatusInternalServerError
	case DomainCodeInternalError:
		return http.StatusInternalServerError
	case DomainCodeNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}

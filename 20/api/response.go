package api

import (
	"go-tdd-clean/20/product"
	"net/http"
	"reflect"
)

type Response struct {
	Errors     []string `json:"errors"`
	Data       any      `json:"data"`
	StatusCode int      `json:"-"`
}

func buildResponse(output product.Output, r *http.Request) Response {
	return Response{
		Errors:     output.GetErrors(),
		Data:       output.GetData(),
		StatusCode: buildHTTPStatusCode(output, r.Method),
	}
}

func buildHTTPStatusCode(output product.Output, verb string) int {
	domainCode := output.GetCode()
	switch domainCode {
	case product.DomainCodeSuccess:
		if verb == http.MethodPost {
			return http.StatusCreated
		}
		if verb == http.MethodGet && (output.GetData() == nil || reflect.ValueOf(output.GetData()).IsNil()) {
			return http.StatusNoContent
		}
		return http.StatusOK
	case product.DomainCodeInvalidInput:
		return http.StatusBadRequest
	case product.DomainCodeInvalidEntity:
		return http.StatusInternalServerError
	case product.DomainCodeInternalError:
		return http.StatusInternalServerError
	case product.DomainCodeNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}

package controllers

import (
	"go-tdd-clean/12/shared"
	"net/http"
	"reflect"
)

type VerbType int

const (
	VerbTypeGet    VerbType = 1
	VerbTypePost   VerbType = 2
	VerbTypePut    VerbType = 3
	VerbTypeDelete VerbType = 4
)

func BuildHttpStatusCode(output shared.Output, verb string) int {
	domainCode := output.GetCode()
	return domainCodeToHttpStatusCode(output, domainCode, verb)
}

func domainCodeToHttpStatusCode(output shared.Output, domainCode shared.DomainCode, verb string) int {
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

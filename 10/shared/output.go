package shared

import "strings"

type DomainCode int

const (
	DomainCodeSuccess DomainCode = 1 + iota
	DomainCodeInvalidInput
	DomainCodeInvalidEntity
	DomainCodeInternalError
	DomainCodeNotFound
)

// Deixamos as propriedades privadas, e só podemos acessá-las através dos métodos, o que nos permite controlar melhor o fluxo e a validação dos dados.
// Dejamos las propiedades privadas, y solo podemos acceder a ellas a través de los métodos, lo que nos permite controlar mejor el flujo y la validación de los datos.
type Output struct {
	errors []string
	code   DomainCode
	data   any
}

func (r *Output) SetError(code DomainCode, err error) {
	r.setCode(code)
	r.errors = append(r.errors, err.Error())
}

func (r *Output) SetErrorMsg(code DomainCode, message string) {
	r.setCode(code)
	r.errors = append(r.errors, message)
}

func (r *Output) SetOk(data any) {
	r.code = DomainCodeSuccess
	r.data = data
}

func (r *Output) GetError() string {
	return strings.Join(r.errors, ", ")
}

func (r *Output) GetErrors() []string {
	return r.errors
}

func (r *Output) GetCode() DomainCode {
	return r.code
}

func (r *Output) GetData() any {
	return r.data
}

func (r *Output) setCode(code DomainCode) {
	if !isValidDomainCode(code) {
		panic("invalid domain code")
	}
	r.code = code
}

func isValidDomainCode(code DomainCode) bool {
	return code == DomainCodeSuccess ||
		code == DomainCodeInvalidInput ||
		code == DomainCodeInvalidEntity ||
		code == DomainCodeInternalError ||
		code == DomainCodeNotFound
}

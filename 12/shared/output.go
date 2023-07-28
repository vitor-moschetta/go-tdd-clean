package shared

type DomainCode int

const (
	DomainCodeInvalidInput  DomainCode = 1
	DomainCodeInvalidEntity DomainCode = 2
	DomainCodeInternalError DomainCode = 3
	DomainCodeSuccess       DomainCode = 4
	DomainCodeNotFound      DomainCode = 5
)

// Deixamos as propriedades privadas, e só podemos acessá-las através dos métodos, o que nos permite controlar melhor o fluxo e a validação dos dados.
// Dejamos las propiedades privadas, y solo podemos acceder a ellas a través de los métodos, lo que nos permite controlar mejor el flujo y la validación de los datos.
type Output struct {
	errors []string
	code   DomainCode
	data   any
}

func (r *Output) SetError(code DomainCode, err string) {
	r.setCode(code)
	r.errors = append(r.errors, err)
}

func (r *Output) SetErrors(code DomainCode, errs []string) {
	r.setCode(code)
	r.errors = append(r.errors, errs...)
}

func (r *Output) SetOk() {
	r.code = DomainCodeSuccess
}

func (r *Output) SetOkWithData(data any) {
	r.code = DomainCodeSuccess
	r.data = data
}

func (r *Output) HasError() bool {
	return len(r.errors) > 0
}

func (r *Output) GetErrors() []string {
	return r.errors
}

func (r *Output) GetCode() DomainCode {
	return r.code
}

func (r *Output) setCode(code DomainCode) {
	if !isValidDomainCode(code) {
		panic("invalid domain code")
	}
	r.code = code
}

func (r *Output) GetData() any {
	return r.data
}

func isValidDomainCode(code DomainCode) bool {
	return code == DomainCodeInvalidInput ||
		code == DomainCodeInvalidEntity ||
		code == DomainCodeInternalError ||
		code == DomainCodeSuccess ||
		code == DomainCodeNotFound
}

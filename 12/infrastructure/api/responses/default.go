package responses

import "go-tdd-clean/12/shared"

type Response struct {
	Errors []string `json:"errors"`
	Data   any      `json:"data"`
}

func OutputToResponse(output shared.Output) Response {
	return Response{
		Errors: output.GetErrors(),
		Data:   output.GetData(),
	}
}

func ItemToResponse(item any, err string) Response {
	return Response{
		Errors: []string{err},
		Data:   item,
	}
}

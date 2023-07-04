package shared

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOutput_SetError(t *testing.T) {
	output := Output{}
	output.SetError(DomainCodeInvalidInput, "error")

	assert.Equal(t, DomainCodeInvalidInput, output.GetCode())
}

func TestOutput_SetErrors(t *testing.T) {
	output := Output{}
	output.SetErrors(DomainCodeInvalidInput, []string{"error1", "error2"})

	assert.Equal(t, DomainCodeInvalidInput, output.GetCode())
	assert.Equal(t, 2, len(output.GetErrors()))
}

func TestOutput_SetOk(t *testing.T) {
	output := Output{}
	output.SetOk()

	assert.Equal(t, DomainCodeSuccess, output.GetCode())
}

func TestOutput_HasError(t *testing.T) {
	output := Output{}
	output.SetError(DomainCodeInvalidInput, "error")

	assert.True(t, output.HasError())
}

func TestOutput_GetErrors(t *testing.T) {
	output := Output{}
	output.SetErrors(DomainCodeInvalidInput, []string{"error1", "error2"})

	assert.Equal(t, 2, len(output.GetErrors()))
}

func TestOutput_SetErrorWithInvalidDomainCode(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(string); ok {
				assert.Equal(t, "invalid domain code", err)
			}
		}
	}()

	output := Output{}
	output.SetError(999, "error")
}

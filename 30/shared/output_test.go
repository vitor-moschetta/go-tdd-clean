package shared

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOutput_SetInvalidDomainCode(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(string); ok {
				assert.Equal(t, "invalid domain code", err)
			}
		}
	}()

	output := Output{}
	output.SetError(999, errors.New("error"))
}

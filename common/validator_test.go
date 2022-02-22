package common

import (
	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestStruct struct {
	ABC string `json:"abc" validate:"required"`
}

func TestCustomValidator_Validate(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		v := &CustomValidator{
			Validator: validator.New(),
		}
		assert.Error(t, v.Validate(TestStruct{}))
	})

	t.Run("success", func(t *testing.T) {
		v := &CustomValidator{
			Validator: validator.New(),
		}
		assert.Nil(t, v.Validate(TestStruct{ABC: "test"}))
	})
}

package handlers

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleError(t *testing.T) {
	errMsg := "Some error: "
	err := errors.New("Internal server error!")
	errStatusCode := 500
	r := handleError(err, errMsg, errStatusCode)
	assert.Equal(t, errStatusCode, r.Code)
	assert.Equal(t, errMsg+err.Error(), r.Message)
}

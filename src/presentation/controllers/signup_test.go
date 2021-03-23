package controllers

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignUpController(t *testing.T) {

	t.Run("Should return 400 if no name provided", func(internal *testing.T) {
		var sut = NewSignUpController()

		var httpRequest = SignUpRequest{
			Email:                "any email",
			Password:             "any password",
			PasswordConfirmation: "any password confirmation",
		}

		var httpResponse = sut.handle(httpRequest)
		assert.Equal(internal, http.StatusBadRequest, httpResponse.StatusCode)
		assert.Equal(internal, errors.New("Missing param: name"), httpResponse.Body)
	})

	t.Run("Should return 400 if no email provided", func(internal *testing.T) {
		var sut = NewSignUpController()

		var httpRequest = SignUpRequest{
			Name:                 "any name",
			Password:             "any password",
			PasswordConfirmation: "any password confirmation",
		}

		var httpResponse = sut.handle(httpRequest)
		assert.Equal(internal, http.StatusBadRequest, httpResponse.StatusCode)
		assert.Equal(internal, errors.New("Missing param: email"), httpResponse.Body)
	})

}

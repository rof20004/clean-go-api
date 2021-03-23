package controllers

import (
	"errors"
	"net/http"
)

type SignUpController struct{}

type SignUpRequest struct {
	Name                 string `json:"name"`
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

type SignUpResponse struct {
	StatusCode int   `json:"status_code"`
	Body       error `json:"body"`
}

func NewSignUpController() SignUpController {
	return SignUpController{}
}

func (sc SignUpController) handle(httpRequest SignUpRequest) SignUpResponse {
	if httpRequest.Name == "" {
		return SignUpResponse{StatusCode: http.StatusBadRequest, Body: errors.New("Missing param: name")}
	}

	if httpRequest.Email == "" {
		return SignUpResponse{StatusCode: http.StatusBadRequest, Body: errors.New("Missing param: email")}
	}

	return SignUpResponse{}
}

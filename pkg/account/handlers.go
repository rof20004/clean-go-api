package account

import (
	"encoding/json"
	"io"
	"net/http"
)

type Handler struct {
	serv *Service
}

func NewHandler(serv *Service) Handler {
	return Handler{serv}
}

func (h Handler) Register(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, err.Error())
		return
	}

	var account Account
	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, err.Error())
		return
	}

	if err := h.serv.Register(&account); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, err.Error())
		return
	}

	b, err := json.Marshal(&account)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, err.Error())
		return
	}

	w.Header().Add("content-type", "application/json")
	w.Write(b)
}

func (h Handler) ListAccounts(w http.ResponseWriter, r *http.Request) {
	accounts, err := h.serv.ListAccounts()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, err.Error())
		return
	}

	b, err := json.Marshal(&accounts)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, err.Error())
		return
	}

	w.Header().Add("content-type", "application/json")
	w.Write(b)
}

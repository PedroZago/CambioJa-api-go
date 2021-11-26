package middlewares

import (
	"api/src/authentication"
	"api/src/response"
	"log"
	"net/http"
)

func Logger(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n%s %s %s", r.Method, r.RequestURI, r.Host)
		proximaFuncao(w, r)
	}
}

func Autenticar(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := authentication.ValidarToken(r); err != nil {
			response.Erro(w, http.StatusUnauthorized, err)
			return
		}
		proximaFuncao(w, r)
	}
}

package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/response"
	"api/src/security"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	corpoRequisicao, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Erro(w, http.StatusUnprocessableEntity, err)
	}

	var usuario models.Usuario
	if err = json.Unmarshal(corpoRequisicao, &usuario); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.ConectarBD()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioUsuario(db)
	usuarioRetorno, err := repositorio.BuscarPorEmail(usuario.Email)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.VerificarSenha(usuarioRetorno.Senha, usuario.Senha); err != nil {
		response.Erro(w, http.StatusUnauthorized, err)
		return
	}

	token, err := authentication.CriarToken(usuarioRetorno.UsuarioID)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	usuarioID := strconv.FormatUint(usuarioRetorno.UsuarioID, 10)

	response.JSON(w, http.StatusOK, models.DadosAutenticacao{UsuarioID: usuarioID, Token: token})
}

package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/response"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func CriarCambio(w http.ResponseWriter, r *http.Request) {
	var iof = 0.1
	var spread = 0.1
	var taxa = 0.1

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	corpoRequisicao, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var cambio models.Cambio
	if err = json.Unmarshal(corpoRequisicao, &cambio); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.ConectarBD()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorioMoeda := repositories.NovoRepositorioMoeda(db)
	cotacao, err := repositorioMoeda.PegarCotacao(cambio.MoedaID)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	cambio.ValorFinal = (cambio.ValorOrigem / cotacao) + (iof * (cambio.ValorOrigem / cotacao)) + (spread * (cambio.ValorOrigem / cotacao)) + (taxa * (cambio.ValorOrigem / cotacao))

	if err = cambio.Preparar(); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	repositorioCambio := repositories.NovoRepositorioCambio(db)
	cambio.CambioID, err = repositorioCambio.CriarCambio(cambio)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusCreated, cambio)
}

func BuscarTodosCambiosPorUsuarioEMoeda(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	corpoRequisicao, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var cambio models.Cambio
	if err = json.Unmarshal(corpoRequisicao, &cambio); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.ConectarBD()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioCambio(db)
	cambios, err := repositorio.BuscarTodosCambiosPorUsuarioEMoeda(cambio)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, cambios)
}

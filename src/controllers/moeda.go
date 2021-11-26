package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/response"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CriarMoeda(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	corpoRequisicao, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var moeda models.Moeda
	if err = json.Unmarshal(corpoRequisicao, &moeda); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = moeda.Preparar(); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.ConectarBD()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioMoeda(db)
	moeda.MoedaID, err = repositorio.CriarMoeda(moeda)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusCreated, moeda)
}

func AtualizarMoeda(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	parametros := mux.Vars(r)
	moedaID, err := strconv.ParseUint(parametros["moedaID"], 10, 64)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	corpoRequisicao, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var moeda models.Moeda
	if err = json.Unmarshal(corpoRequisicao, &moeda); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = moeda.Preparar(); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.ConectarBD()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioMoeda(db)
	if err = repositorio.AtualizarMoeda(moedaID, moeda); err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func DeletarMoeda(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	parametros := mux.Vars(r)
	moedaID, err := strconv.ParseUint(parametros["moedaID"], 10, 64)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.ConectarBD()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioMoeda(db)
	if err = repositorio.DeletarMoeda(moedaID); err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func BuscarMoedaPorID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	parametros := mux.Vars(r)
	moedaID, err := strconv.ParseUint(parametros["moedaID"], 10, 64)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.ConectarBD()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioMoeda(db)
	moeda, err := repositorio.BuscarMoedaPorID(moedaID)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, moeda)
}

func BuscarTodasMoedas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	db, err := database.ConectarBD()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioMoeda(db)
	moeda, err := repositorio.BuscarTodasMoedas()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, moeda)
}

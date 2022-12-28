package controllers

import (
	"encoding/json"
	//"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"../models"
	"../response"
	"../utils/formaterror"
	//"../auth"
)

func (server *Server) CreatePurchase(w http.ResponseWriter, r *http.Request) {


	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	purchase := models.Purchase{}
	err = json.Unmarshal(body, &purchase)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	
	purchaseCreated, err := purchase.SavePurchase(server.DB)

	if err != nil {

		formattedError := formaterror.FormatError(err.Error())

		response.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, purchaseCreated.ID))
	response.JSON(w, http.StatusCreated, purchaseCreated)
}

func (server *Server) GetPurchases(w http.ResponseWriter, r *http.Request) {

	purchase := models.Purchase{}

	purchases, err := purchase.FindAllPurchases(server.DB)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusOK, purchases)
}

func (server *Server) GetPurchaseById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)

	pid, err := strconv.ParseUint(vars["id"], 10, 32)

	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	purchase := models.Purchase{}

	purchaseGotten , err := purchase.FindPurchasesByID(server.DB, uint32(pid))
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	response.JSON(w, http.StatusOK, purchaseGotten)
}

func (server *Server) UpdatePurchase(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	purchase := models.Purchase{}

	err = json.Unmarshal(body, &purchase)

	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	purchaseProduct, err := purchase.UpdatePurchases(server.DB, uint32(pid))

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		response.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	response.JSON(w, http.StatusOK, purchaseProduct)
}

func (server *Server) DeletePurchase(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)

	purchase := models.Purchase{}

	pid,err := strconv.ParseUint(vars["id"], 10, 32)

	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	_, err = purchase.DeletePurchases(server.DB, uint32(pid))

	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", pid))
	response.JSON(w, http.StatusNoContent, "")
}
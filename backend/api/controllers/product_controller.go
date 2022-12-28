package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/Bernnardo-Pestana/Problem/api/models"
	"github.com/Bernnardo-Pestana/Problem/api/response"
	"github.com/Bernnardo-Pestana/Problem/api/utils/formaterror"

)

func (server *Server) CreateProduct(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	product := models.Product{}
	err = json.Unmarshal(body, &product)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	
	productCreated, err := product.SaveProduct(server.DB)

	if err != nil {

		formattedError := formaterror.FormatError(err.Error())

		response.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, productCreated.ID))
	response.JSON(w, http.StatusCreated, productCreated)
}

func (server *Server) GetProducts(w http.ResponseWriter, r *http.Request) {

	product := models.Product{}

	products, err := product.FindAllProducts(server.DB)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusOK, products)
}

func (server *Server) GetProductById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)

	pid, err := strconv.ParseUint(vars["id"], 10, 32)

	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	product := models.Product{}

	productGotten , err := product.FindProductByID(server.DB, uint32(pid))
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	response.JSON(w, http.StatusOK, productGotten)
}

func (server *Server) UpdateProduct(w http.ResponseWriter, r *http.Request){
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

	product := models.Product{}

	err = json.Unmarshal(body, &product)

	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	updateProduct, err := product.UpdateProduct(server.DB, uint32(pid))

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		response.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	response.JSON(w, http.StatusOK, updateProduct)
}

func (server *Server) DeleteProduct(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)

	product := models.Product{}

	pid,err := strconv.ParseUint(vars["id"], 10, 32)

	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	_, err = product.DeleteProduct(server.DB, uint32(pid))

	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", pid))
	response.JSON(w, http.StatusNoContent, "")
}
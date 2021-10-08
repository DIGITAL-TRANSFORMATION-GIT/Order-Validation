package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (h *HTTPHandler) GetStatusOfAllOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	retrievedData, err := h.database.RetrieveData("SELECT * FROM orders;")
	if err != nil {
		h.logger.ErrorLogger.Println("Can't retrieve data from SQL: ", err.Error())
		w.WriteHeader(http.StatusNotFound)
		return
	}
	response, err := h.retrieveOrders(retrievedData)
	if err != nil {
		h.logger.InfoLogger.Println("No order is present")
		w.WriteHeader(http.StatusNotFound)
	}
	resp, _ := json.Marshal(response)
	fmt.Fprint(w, string(resp))
}

func (h *HTTPHandler) GetStatusOfOrder(w http.ResponseWriter, r *http.Request) {
	request := mux.Vars(r)
	id := request["id"]
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	retrievedData, err := h.database.RetrieveData(fmt.Sprintf("SELECT * FROM orders where id = %s;", id))
	if err != nil {
		h.logger.ErrorLogger.Println("Can't retrieve data from SQL: ", err.Error())
		w.WriteHeader(http.StatusNotFound)
		return
	}
	response, err := h.retrieveOrders(retrievedData)
	if err != nil {
		h.logger.ErrorLogger.Println("Resource not found : ", id)
		w.WriteHeader(http.StatusNotFound)
	}
	resp, _ := json.Marshal(response)
	fmt.Fprint(w, string(resp))

}

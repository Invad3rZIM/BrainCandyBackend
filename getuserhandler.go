package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody map[string]interface{}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	//ensure all requisite json components are found
	if err := h.VerifyBody(requestBody, "uid", "pin"); err != nil {
		fmt.Fprintln(w, err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	uid := int(requestBody["uid"].(float64))

	u, err := h.GetUser(uid)

	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(err.Error())

		return
	}

	json.NewEncoder(w).Encode(&u)
	w.WriteHeader(http.StatusOK)
}

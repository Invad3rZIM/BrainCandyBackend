package main

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) NewUserHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody map[string]interface{}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	//adds user to cache
	u := h.UserCache.NewUser()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&u)
}

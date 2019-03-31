package main

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) HasOldMessagesHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(err.Error())

		return
	}

	//ensure all requisite json components are found
	if err := h.VerifyBody(requestBody, "uid", "pin"); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(err.Error())

		return
	}

	uid := int(requestBody["uid"].(float64))

	json.NewEncoder(w).Encode(h.CheckHasNewMessages(uid))
	w.WriteHeader(http.StatusOK)
}

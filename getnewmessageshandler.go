package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) GetNewMessagesHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody map[string]interface{}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	//ensure all requisite json components are found
	if err := h.VerifyBody(requestBody, "uid"); err != nil {
		json.NewEncoder(w).Encode(&err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	uid := int(requestBody["uid"].(float64))

	container := MessageContainer{
		Messages: h.GetNewMessages(uid),
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&container)
}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (h *Handler) SendMessageHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody map[string]interface{}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	//ensure all requisite json components are found
	if err := h.VerifyBody(requestBody, "sid", "rid", "secretid", "message", "num"); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	//sender and receiver ids
	sid := int(requestBody["sid"].(float64))
	rid := int(requestBody["rid"].(float64))
	secretID := int(requestBody["secretid"].(float64))
	mes := requestBody["message"].(string)
	number := int(requestBody["num"].(float64))

	m := Message{
		Num:       number,
		Body:      mes,
		To:        rid,
		From:      sid,
		SecretID:  secretID,
		TimeStamp: time.Now().Format("01-02-2006"),
	}

	h.PostMessage(&m)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&m)
}

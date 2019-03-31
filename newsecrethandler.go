package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func (h *Handler) NewSecretHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody map[string]interface{}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(err.Error())

		return
	}
	//ensure all requisite json components are found
	if err := h.VerifyBody(requestBody, "secret", "sid"); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	uid := int(requestBody["sid"].(float64))

	secret := requestBody["secret"].(string)

	s := Secret{
		SecretID:  h.GenSecretId(),
		SenderID:  uid,
		Body:      secret,
		TimeStamp: time.Now().Format("01-02-2006"),
	}

	h.SecretCache.CacheSecret(&s)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&s)
}

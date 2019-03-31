package main

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) GetSecretsHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	//ensure all requisite json components are found
	if err := h.VerifyBody(requestBody, "blacklist", "sid"); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	sid := int(requestBody["sid"].(float64))
	//bl = blacklist
	bbl := requestBody["blacklist"].([]interface{})

	//to hide the repeated secrets so we don't overfloat
	blacklist := make(map[int]struct{})

	for i := range bbl {
		blacklist[int(bbl[i].(float64))] = struct{}{}
	}

	var secrets []*Secret

	secrets = *h.GetSecrets(blacklist, sid, 10)

	cont := SecretContainer{
		Secrets: secrets,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&cont)
}

type SecretContainer struct {
	Secrets []*Secret
}

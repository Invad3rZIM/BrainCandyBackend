package main

import (
	"net/http"
)

func (h *Handler) InitMessageHandler(w http.ResponseWriter, r *http.Request) {
	/*	var requestBody map[string]interface{}

		if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(err.Error())

			return
		}

		//ensure all requisite json components are found
		if err := h.VerifyBody(requestBody, "uid", "pin", "bid", "body"); err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(err.Error())

			return
		}

		uid := int(requestBody["uid"].(float64))

		bid := int(requestBody["bid"].(float64))

		if to == uid {
			err := errors.New("error: can't send message to yourself!")
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(err.Error())

			return
		}

		body := requestBody["body"].(string)

		m := Message{
			Num:  0,
			Body: body,
			Bid:  bid,
			To:   h.BottleOrigin(bid),
			From: uid,
		}

		h.PostMessage(&m)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&m)*/
}

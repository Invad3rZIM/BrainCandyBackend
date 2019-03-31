package main

import (
	"errors"
	"fmt"
)

type Handler struct {
	*UserCache
	*SecretCache
	*MessageCache
}

func NewHandler() *Handler {
	h := Handler{
		NewUserCache(),
		NewSecretCache(),
		NewMessageCache(),
	}

	return &h
}

func (h *Handler) LoadTestData() {
	h.CacheUser(&User{
		Uid: 1,
	})

	h.CacheUser(&User{
		Uid: 2,
	})
}

//VerifyBody is a helper function to ensure all http requests contain the requisite fields returns error if fields missing
func (h *Handler) VerifyBody(body map[string]interface{}, str ...string) error {
	for _, s := range str {
		fmt.Println(s)
		if _, ok := body[s]; !ok {
			return errors.New("error: missing field: " + s)
		}
	}

	return nil
}

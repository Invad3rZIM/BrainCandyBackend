package main

import (
	"fmt"
	"time"
)

type SecretCache struct {
	Secrets map[int]*Secret //Geo + Global = this
}

func NewSecretCache() *SecretCache {
	return &SecretCache{
		Secrets: make(map[int]*Secret),
	}
}

func (sc *SecretCache) SecretExists(bid int) bool {
	_, ok := sc.Secrets[bid]

	return ok
}

//randomly generates 10 digit uid until a unique one is created
func (sc *SecretCache) GenSecretId() int {
	for {
		r := GenInt(6)

		if !sc.SecretExists(r) {
			return r
		}
	}
}

func (sc *SecretCache) GenTestSecrets(num int) {
	for i := 0; i < num; i++ {
		s := Secret{
			SecretID:  sc.GenSecretId(),
			SenderID:  i,
			Body:      fmt.Sprintf("%s %d", "this is a test thingy", i),
			TimeStamp: time.Now().Format("01-02-2006"),
		}

		sc.CacheSecret(&s)
	}
}

func (sc *SecretCache) CacheSecret(s *Secret) {
	sc.Secrets[s.SecretID] = s
}

func (sc *SecretCache) GetSecrets(blacklist map[int]struct{}, sid int, limit int) *[]*Secret {
	cache := []*Secret{}

	lim := 0

	for _, v := range sc.Secrets {
		if _, ok := blacklist[v.SecretID]; ok { //get rid of all the blacklist stuff
			continue
		}

		if v.SenderID == sid {
			continue
		}

		lim++
		cache = append(cache, v)

		if lim >= limit {
			break
		}
	}

	return &cache
}

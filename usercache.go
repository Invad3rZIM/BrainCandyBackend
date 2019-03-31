package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type UserCache struct {
	Users map[int]*User
}

func NewUserCache() *UserCache {
	return &UserCache{
		Users: make(map[int]*User),
	}
}

func (uc *UserCache) GetUser(uid int) (*User, error) {
	if u, ok := uc.Users[uid]; ok {
		return u, nil
	} else {
		return nil, errors.New(fmt.Sprintf("uid %d not found in userCache", uid))
	}
}

func (uc *UserCache) NewUser() *User {
	u := User{
		Uid: uc.GenUid(),
	}

	uc.CacheUser(&u)

	return &u
}

func (uc *UserCache) UserExists(uid int) bool {
	_, ok := uc.Users[uid]

	return ok
}

//randomly generates 10 digit uid until a unique one is created
func (uc *UserCache) GenUid() int {
	rand.Seed(time.Now().UTC().UnixNano())

	for {
		r := GenInt(6)

		if !uc.UserExists(r) {
			return r
		}
	}
}

func (uc *UserCache) CacheUser(u *User) {
	uc.Users[u.Uid] = u
}

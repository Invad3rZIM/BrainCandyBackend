package main

type FriendCache struct {
	FriendCache   map[int][]*Friendship
	NewFriendKeys map[int]bool
}

type Friendship struct {
	Uid       int
	FriendKey int
}

func NewFriendCache() *FriendCache {
	return &FriendCache{
		FriendCache:   make(map[int][]*Friendship),
		NewFriendKeys: make(map[int]bool),
	}
}

func (fc *FriendCache) GiveFriendKey(to, from, key int) {
	f := Friendship{
		Uid:       from,
		FriendKey: key,
	}

	//if the person doesn't have a friend mailbox yet
	if cache, found := fc.FriendCache[to]; !found || cache == nil {
		fc.FriendCache[to] = []*Friendship{}
	}

	fc.FriendCache[to] = append(fc.FriendCache[to], &f)
	fc.NewFriendKeys[to] = true
}

func (fc *FriendCache) GetFriendKeys(uid int) []*Friendship {
	if friendships, found := fc.FriendCache[uid]; found {
		delete(fc.FriendCache, uid)
		fc.FriendCache[uid] = []*Friendship{}

		return friendships
	} else {
		fc.FriendCache[uid] = []*Friendship{}
		return fc.FriendCache[uid]
	}
}

func (fc *FriendCache) HasNewFriendKeys(uid int) bool {
	if hasNew, found := fc.NewFriendKeys[uid]; !found {
		fc.NewFriendKeys[uid] = false
		return false
	} else {
		return hasNew
	}
}

package main

type MessageCache struct {
	NewMessages    map[int]*[]*Message
	HasNewMessages map[int]bool
	OldMessages    map[int]*[]*Message
	HasOldMessages map[int]bool
}

func NewMessageCache() *MessageCache {
	return &MessageCache{
		NewMessages:    make(map[int]*[]*Message),
		HasNewMessages: make(map[int]bool),
		OldMessages:    make(map[int]*[]*Message),
		HasOldMessages: make(map[int]bool),
	}
}

//adds new mails
func (mc *MessageCache) PostMessage(message *Message) {

	//create them if they're nil
	if _, ok := mc.NewMessages[message.To]; !ok {
		mc.NewMessages[message.To] = &[]*Message{}
	}

	if _, ok := mc.OldMessages[message.To]; !ok {
		mc.OldMessages[message.To] = &[]*Message{}
	}

	mbox := *mc.NewMessages[message.To]
	mbox = append(mbox, message)
	mc.NewMessages[message.To] = &mbox

	mc.HasNewMessages[message.To] = true
}

//copies all new messages to old mailbox, returns all the new messages that were just copied
func (mc *MessageCache) GetNewMessages(to int) []*Message {
	m := []*Message{}

	//create them if they're nil
	if _, ok := mc.NewMessages[to]; !ok {
		mc.NewMessages[to] = &[]*Message{}
	}

	if _, ok := mc.OldMessages[to]; !ok {
		mc.OldMessages[to] = &[]*Message{}
	}

	new := *mc.NewMessages[to]
	old := *mc.OldMessages[to]
	old = append(old, new...)
	mc.OldMessages[to] = &old
	mc.NewMessages[to] = &m

	mc.HasNewMessages[to] = false
	mc.HasOldMessages[to] = len(*mc.OldMessages[to]) > 0

	return new
}

//copies all new messages to old mailbox, returns all the new messages that were just copied
func (mc *MessageCache) GetOldMessages(to int) []*Message {
	old := *mc.OldMessages[to]
	mc.HasOldMessages[to] = true
	return old
}

func (mc *MessageCache) CheckHasNewMessages(uid int) bool {
	if hasNew, found := mc.HasNewMessages[uid]; !found {
		return false
	} else {
		return hasNew
	}
}

func (mc *MessageCache) CheckHasOldMessages(uid int) bool {
	if hasOld, found := mc.HasOldMessages[uid]; !found {
		return false
	} else {
		return hasOld
	}
}

package models

import "sync"

type OnlineUsers struct {
	users sync.Map
}

func (o *OnlineUsers) Add(username string) {
	o.users.Store(username, nil)
}

func (o *OnlineUsers) Delete(username string) {
	o.users.Delete(username)
}

func (o *OnlineUsers) GetAll() []string {
	users := make([]string, 0)
	o.users.Range(func(key, value any) bool {
		users = append(users, key.(string))
		return true
	})

	return users
}

func NewOnlineUsers() *OnlineUsers {
	return &OnlineUsers{
		users: sync.Map{},
	}
}

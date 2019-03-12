package proxy

import "fmt"

type User struct {
	ID int32
}

type UserFinder interface {
	FindUser(id string) (User, error)
}

type UserList []User

func (u UserList) FindUser(id int32) (User, error) {
	for i := 0; i < len(u); i++ {
		if u[i].ID == id {
			return u[i], nil
		}
	}
	return User{}, fmt.Errorf("User %d was not found\n", id)
}

type UserListProxy struct {
	MockedDatabase      *UserList
	StackCache          UserList
	StackSize           int
	LastSearchUsedCache bool
}

func (proxy *UserListProxy) FindUser(id int32) (User, error) {
	user, err := proxy.StackCache.FindUser(id)
	if err == nil {
		proxy.LastSearchUsedCache = true
		return user, nil
	}
	user, err = proxy.MockedDatabase.FindUser(id)
	if err == nil {
		proxy.LastSearchUsedCache = false
		proxy.addUserToStack(user)
		return user, nil
	}
	return user, err
}

func (proxy *UserListProxy) addUserToStack(user User) {
	if len(proxy.StackCache) >= proxy.StackSize {
		proxy.StackCache = append(proxy.StackCache[1:], user)
	} else {
		proxy.StackCache = append(proxy.StackCache, user)
	}
}

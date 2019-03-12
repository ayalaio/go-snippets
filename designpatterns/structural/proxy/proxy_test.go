package proxy

import (
	"math/rand"
	"testing"
)

func Test_UserListProxy(t *testing.T) {
	mockedDatabase := UserList{}

	rand.Seed(235463456)
	for i := 0; i < 1000000; i++ {
		n := rand.Int31()
		mockedDatabase = append(mockedDatabase, User{ID: n})
	}

	proxy := UserListProxy{
		MockedDatabase: &mockedDatabase,
		StackSize:      2,
		StackCache:     UserList{},
	}
	knownIDs := [3]int32{
		mockedDatabase[3].ID,
		mockedDatabase[4].ID,
		mockedDatabase[5].ID,
	}
	t.Run("FindUser - Empty Cache", func(t *testing.T) {
		user, err := proxy.FindUser(knownIDs[0])
		if err != nil {
			t.Fatal(err)
		}
		if user.ID != knownIDs[0] {
			t.Error("Returned user is not the expected")
		}
		if len(proxy.StackCache) != 1 {
			t.Error("After 1 search, The size of cache should be one")
		}
		if proxy.LastSearchUsedCache == true {
			t.Error("The cache is not used on the first time")
		}
	})
	t.Run("FindUser - One user, ask for same", func(t *testing.T) {
		user, err := proxy.FindUser(knownIDs[0])
		if err != nil {
			t.Fatal(err)
		}
		if user.ID != knownIDs[0] {
			t.Error("Returned user is not the expected")
		}
		if len(proxy.StackCache) != 1 {
			t.Error("After 1 search, The size of cache still be one")
		}
		if proxy.LastSearchUsedCache != true {
			t.Error("The cache should be used for a second search of same")
		}
	})
	t.Run("FindUser - Inserting 3, overflowing the stack", func(t *testing.T) {
		user, err := proxy.FindUser(knownIDs[0])
		if err != nil {
			t.Fatal(err)
		}
		proxy.FindUser(knownIDs[1])
		if proxy.LastSearchUsedCache {
			t.Error("Shouldn't have used cache, knownIDs[1] is not there")
		}

		proxy.FindUser(knownIDs[2])
		if proxy.LastSearchUsedCache {
			t.Error("Shouldn't have used cache, knownIDs[2] is not there")
		}
		for i := 0; i < len(proxy.StackCache); i++ {
			if proxy.StackCache[i].ID == user.ID {
				t.Error("knownIDs[0] shouldnt be on the cache anymore")
			}
		}
		if len(proxy.StackCache) != 2 {
			t.Error("When inserting 3 users. cache should not grow more than 2 users")
		}

	})

}

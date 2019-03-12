package main

import (
	"log"
	"time"
)

func users() []User {
	return []User{
		User{Name: "David", Email: "david@rodriguez"},
		User{Name: "Diego", Email: "diego@anaya"},
		User{Name: "Daniel", Email: "daniel@rangel"},
	}
}

type User struct {
	Name  string
	Email string
}

type SearchWorker struct {
	Users []User
	Out   chan<- User
}

func (w SearchWorker) Search(email string) {
	users := w.Users
	for _, user := range users {
		if email == user.Email {
			w.Out <- user
		}
	}
}

func presenter(ch <-chan User) {
	for user := range ch {
		log.Printf("Found %s with email %s", user.Name, user.Email)
	}
}

func main() {
	log.Println("Hello")

	pipe := make(chan User)

	worker := SearchWorker{Users: users(), Out: pipe}
	go worker.Search("david@rodriguez")
	go presenter(pipe)

	for {
		time.Sleep(time.Duration(300) * time.Second)
	}

}

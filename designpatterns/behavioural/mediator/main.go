package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) Receive(m Message) {
	fmt.Printf("Received! Reading... \n")
	m.Read()
}

func (self *Person) ComposeMessage(subject, to string) Message {
	fmt.Println("Creating Msg")
	return Message{
		Subject: subject,
		From:    self.Name,
		To:      to,
	}
}

type Message struct {
	Subject string
	To      string
	From    string
}

func (m Message) Read() {
	fmt.Printf("%s \nfrom %s, to %s\n", m.Subject, m.From, m.To)
}

type Mediator struct {
	Customers []Person
}

func (m *Mediator) Send(msg Message) {
	for _, c := range m.Customers {
		if c.Name == msg.To {
			c.Receive(msg)
		}
	}
}

func main() {
	p1 := Person{Name: "Juan"}
	p2 := Person{Name: "Alex"}
	mailman := Mediator{
		Customers: []Person{p1, p2},
	}
	message := p1.ComposeMessage("Sample Message", "Alex")
	mailman.Send(message)

}

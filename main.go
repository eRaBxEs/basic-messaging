package main

import (
	"fmt"
	"time"
)

// Defining structs for objects that includes datamodel, messages

// DataModel
type DataModel struct {
	ID      int
	Content string
}

// Message
type Message struct {
	Event   string
	Payload DataModel
}

// CRUD operations
//
//	Create
func Create(data DataModel, messages chan<- Message) {

	// Add data to the database ...

	// Send a create event message
	messages <- Message{Event: "CREATE", Payload: data}
	time.Sleep(2 * time.Second)
	fmt.Println("Create done")
}

// Update
func Update(id int, newData DataModel, messages chan<- Message) {

	// update data in the database ...

	// send an update event message
	messages <- Message{Event: "UPDATE", Payload: newData}
	time.Sleep(2 * time.Second)
	fmt.Println("Update done")
}

// Update
func Delete(id int, messages chan<- Message) {

	// delete data from the database ...

	// send an update event message
	messages <- Message{Event: "DELETE", Payload: DataModel{ID: id}}
	time.Sleep(2 * time.Second)
	fmt.Println("Delete done")
}

// / Now to the Producer func which simulates the CREATE, UPDATE and DELETE
func Producer(messages chan<- Message) {

	// simulation of the data creation
	// add time sleep for proper simulation
	Create(DataModel{ID: 1, Content: "First Message"}, messages)
	time.Sleep(2 * time.Second)

	// simulation of the data update
	Update(1, DataModel{ID: 1, Content: "Updated Message"}, messages)
	time.Sleep(2 * time.Second)

	// simulation of the data deletion
	Delete(1, messages)
	close(messages)
}

// Consumer reads messages from the channel and processes them
func Consumer(messages <-chan Message) {
	for msg := range messages {
		fmt.Println("Received message:", msg)
	}
}

func main() {
	// create a channel for messages
	messages := make(chan Message)

	// start the producer go routine which in turn starts the consumer goroutine
	go Producer(messages)

	go Consumer(messages)

	// Wait to enable Producer and Consumer to finish
	time.Sleep(10 * time.Second)

}

package server

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func Run() {
	fmt.Println("Main server running")
	wg.Add(2)
	go startAddressBook()
	go startMailServer()
	wg.Wait()
}

func startMailServer() {
	defer wg.Done()

}

func startAddressBook() {
	defer wg.Done()

}

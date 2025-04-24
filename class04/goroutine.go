package main

import (
	"fmt"
	"sync"
	"time"
)

type Message struct {
	Chats   []string
	Friends []string
}

func main() {
	now := time.Now()
	id, err := getUser("John Pork")
	if err != nil {
		panic(err)
	}
	fmt.Println(id)

	ch := make(chan *Message)
	var wg sync.WaitGroup
	wg.Add(2)

	go getChats(id, ch, &wg)
	go getFriends(id, ch, &wg)

	go func() {
		wg.Wait()
		close(ch)
	}()

	for msg := range ch {
		fmt.Println(msg)
	}
	fmt.Println(time.Since(now))
}

func getFriends(id string, ch chan<- *Message, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second * 1)
	ch <- &Message{
		Friends: []string{"joao", "mario"},
	}
}

func getChats(id string, ch chan<- *Message, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second * 3)
	ch <- &Message{
		Chats: []string{"I", "hate", "my life"},
	}
}

func getUser(user string) (string, error) {
	time.Sleep(time.Second * 1)
	return fmt.Sprintln(user), nil
}

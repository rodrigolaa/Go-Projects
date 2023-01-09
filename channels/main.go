package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.com",
		"https://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	//infinity loop

	for l := range c {
		go func(link string) {

			time.Sleep(5 * time.Second)
			checkLink(link, c) // this loop permits to check forever the status of the url, inserting the link again in the channel.

		}(l)

	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, " minght be down!")

		c <- link
		return
	}

	fmt.Println(link, " is Up!")
	c <- link
}

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	//1) All response
	//fmt.Println(resp)

	//2) Print only body of response
	// bs := make([]byte, 99999)

	// resp.Body.Read(bs)

	// fmt.Println(string(bs))

	//3) Short built-in way to print the body of response
	//io.Copy(os.Stdout, resp.Body)

	//4) Custom way to print the body response
	lw := logWriter{}

	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {

	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}

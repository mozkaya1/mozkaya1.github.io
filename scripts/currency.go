package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("https://code.musti.uk/coin.sh")
	if err != nil {
		fmt.Println(err)
	}

	res, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(res))

	// fmt.Printf("resp.Body: %v\n", res)
}

package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	code := flag.String("code", "", "Enter a code so the URL can be shortened. (Eg: code `cat` can be mapped to a URL, and the resulting one after the shortening would be `http://example.com/cat`)")
	link := flag.String("link", "", "This is the link that the code represents. For example, if the code is `cat`, it would redirect to the URL of a cat video")

	flag.Parse()

	if len(*code) <= 0 {
		fmt.Println("Code cannot be blank")
		return
	} else if len(*link) <= 0 {
		fmt.Println("Link cannot be blank")
		return
	}

	jsonData := map[string]string{"code": *code, "destinationURL": *link}

	jsonValue, err := json.Marshal(jsonData)

	if err != nil {
		panic(err)
	}

	request, err := http.NewRequest("POST", "http://localhost:5000/create", bytes.NewBuffer(jsonValue))
	if err != nil {
		panic(err)
	}
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		fmt.Printf("Error: %s\n", err)

	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

}

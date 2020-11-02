package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {

	resp, err := http.Get("https://swapi.dev/api/starships/9")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var result map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	for k, v := range result {
		fmt.Println(k, ": ", v)
	}

}

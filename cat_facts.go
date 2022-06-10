package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CatFact struct {
	Fact   string
	Length int
}

func main() {
	resp, err := http.Get("https://catfact.ninja/fact")
	if err != nil {
		fmt.Println("Error Get: ", err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error Read: ", err)
		return
	}

	var fact CatFact

	json.Unmarshal(body, &fact)

	fmt.Println(fact.Fact)
}

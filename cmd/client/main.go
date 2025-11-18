package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"task-api/internal/models"
)

func main() {
	req := models.CreatePersonRequest{
		Name: "Matheus Griza",
	}

	b, err := json.Marshal(req)

	if err != nil {
		panic(err)
	}

	resp, err := http.Post("http://localhost:8080/persons", "application/json", bytes.NewReader(b))

	if err != nil {
		panic(err)
	}

	if resp.StatusCode != http.StatusCreated {
		panic("Erro creating person")
	}
	var responseApi models.CreatePersonResponse
	if err := json.NewDecoder(resp.Body).Decode(&responseApi); err != nil {
		panic(err)
	}

	fmt.Println("New person created ", responseApi)
}

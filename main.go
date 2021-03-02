package main

import (
	"fmt"

	"github.com/wilgustavo/gostar/starship/adapters/csv"
	"github.com/wilgustavo/gostar/starship/adapters/rest"
)

func main() {
	rest := rest.NewStarshipRestDeserializer()
	lista, err := rest.ListSharships()
	if err != nil {
		fmt.Println("Error al obtener la lista", err)
	} else {
		repo := csv.NewStarshipCSVSerializer("prueba.csv")
		repo.SaveSharships(lista)
	}
}

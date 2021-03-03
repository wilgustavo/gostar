package rest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/wilgustavo/gostar/internal/starship"
)

type deserializer struct {
	url string
}
type shipJSON struct {
	Count   int    `json:"count"`
	Next    string `json:"next"`
	Results []struct {
		Name                 string `json:"name"`
		Model                string `json:"model"`
		Manufacturer         string `json:"manufacturer"`
		CostInCredits        string `json:"cost_in_credits"`
		Length               string `json:"length"`
		MaxAtmospheringSpeed string `json:"max_atmosphering_speed"`
		Crew                 string `json:"crew"`
		Passengers           string `json:"passengers"`
		CargoCapacity        string `json:"cargo_capacity"`
		Consumables          string `json:"consumables"`
		HyperdriveRating     string `json:"hyperdrive_rating"`
		MGLT                 string `json:"MGLT"`
		StarshipClass        string `json:"starship_class"`
		URL                  string `json:"url"`
	} `json:"results"`
}

const apiEndpoint = "https://swapi.dev/api/starships"

// NewStarshipRestDeserializer create a Deserializer
func NewStarshipRestDeserializer() starship.Deserializer {
	return &deserializer{url: apiEndpoint}
}

// ListSharships list a Sharships from API Rest
func (d *deserializer) ListSharships() ([]starship.Starship, error) {
	var datos *shipJSON

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	response, error := client.Get(d.url)
	if error != nil {
		return nil, error
	}

	contents, error := ioutil.ReadAll(response.Body)
	if error != nil {
		return nil, error
	}

	error = json.Unmarshal(contents, &datos)
	return mapShip(datos), error
}

func mapShip(datos *shipJSON) []starship.Starship {
	var lista []starship.Starship
	for _, item := range datos.Results {
		lista = append(lista, starship.Starship{
			Name:         item.Name,
			Model:        item.Model,
			Manufacturer: item.Manufacturer,
			Cost:         item.Manufacturer,
			Length:       item.Length,
		})
	}
	return lista
}

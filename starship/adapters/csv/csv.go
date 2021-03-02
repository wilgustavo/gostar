package csv

import (
	"encoding/csv"
	"os"

	"github.com/wilgustavo/gostar/starship"
)

type serializer struct {
	fileName string
}

// NewStarshipCSVSerializer create a CSV Serializer
func NewStarshipCSVSerializer(fileName string) starship.Serializer {
	return &serializer{fileName: fileName}
}

// SaveSharships save into CSV file
func (s *serializer) SaveSharships(list []starship.Starship) error {
	file, error := os.Create(s.fileName)
	defer file.Close()

	if error != nil {
		return error
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, starship := range list {
		error = writer.Write(starship.ToSlice())
		if error != nil {
			return error
		}
	}

	return error
}

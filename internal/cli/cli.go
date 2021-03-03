package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wilgustavo/gostar/internal/starship"
	"github.com/wilgustavo/gostar/internal/starship/adapters/csv"
	"github.com/wilgustavo/gostar/internal/starship/adapters/rest"
)

// CobraFn funcion para ejecutar un comando
type CobraFn func(cmd *cobra.Command, args []string)

const fileFlag = "file"

// ListarArchivoCmd inicializa comando para lisar elementos en un archivo
func ListarArchivoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "write",
		Short: "Write CSV file with data from API",
		Run:   writeToFile(),
	}

	cmd.Flags().StringP(fileFlag, "f", "", "file name")

	return cmd
}

func writeToFile() CobraFn {
	return func(cmd *cobra.Command, args []string) {
		filename, _ := cmd.Flags().GetString(fileFlag)
		if filename == "" {
			fmt.Println("error: missing file name")
			return
		}
		deserializer := rest.NewStarshipRestDeserializer()
		serializer := csv.NewStarshipCSVSerializer(filename)

		err := CreateFile(deserializer, serializer)
		if err != nil {
			fmt.Println(err)
		}
	}
}

// CreateFile read data from API then write it in a file
func CreateFile(deserializer starship.Deserializer, serializer starship.Serializer) error {
	lista, err := deserializer.ListSharships()
	if err != nil {
		return err
	}

	err = serializer.SaveSharships(lista)
	return err
}

package main

import (
	"github.com/spf13/cobra"
	"github.com/wilgustavo/gostar/internal/cli"
)

func main() {
	rootCmd := &cobra.Command{Use: "starship"}
	rootCmd.AddCommand(cli.ListarArchivoCmd())
	rootCmd.Execute()
}

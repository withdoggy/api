package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/withdoggy/api/petsmanagement/db"
	"github.com/withdoggy/api/petsmanagement/http"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Run api server",
	Run: func(cmd *cobra.Command, args []string) {
		d, err := db.NewFirestoreClient("withdoggy")
		defer d.Close()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		http.Run(d)
	},
}

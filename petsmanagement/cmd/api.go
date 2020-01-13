package cmd

import (
	"fmt"
	"os"

	"contrib.go.opencensus.io/exporter/stackdriver"
	"github.com/spf13/cobra"
	"github.com/withdoggy/api/petsmanagement/db"
	"github.com/withdoggy/api/petsmanagement/http"
	"go.opencensus.io/trace"
)

const projectID = "withdoggy"

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Run api server",
	Run: func(cmd *cobra.Command, args []string) {

		sd, err := stackdriver.NewExporter(stackdriver.Options{
			ProjectID: projectID,
		})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer sd.Flush()
		trace.RegisterExporter(sd)
		d, err := db.NewFirestoreClient(projectID)
		defer d.Close()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		http.Run(d)
	},
}

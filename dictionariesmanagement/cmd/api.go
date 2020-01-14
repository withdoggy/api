package cmd

import (
	"fmt"
	"os"
	"time"

	"contrib.go.opencensus.io/exporter/stackdriver"
	"github.com/spf13/cobra"
	"github.com/withdoggy/api/dictionariesmanagement/http"
	"go.opencensus.io/trace"
)

const projectID = "withdoggy"

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Run api server",
	Run: func(cmd *cobra.Command, args []string) {
		sd, err := stackdriver.NewExporter(stackdriver.Options{
			ProjectID:         projectID,
			MetricPrefix:      "dictionariesmanagement-api",
			ReportingInterval: 60 * time.Second,
		})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer sd.Flush()

		sd.StartMetricsExporter()
		defer sd.StopMetricsExporter()
		trace.RegisterExporter(sd)
		trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		http.Run()
	},
}

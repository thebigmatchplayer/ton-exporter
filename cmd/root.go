package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/thebigmatchplayer/ton-exporter/exporter"
	apiv2 "github.com/thebigmatchplayer/ton-exporter/ton-client/api-v2"
	apiv3 "github.com/thebigmatchplayer/ton-exporter/ton-client/api-v3"
)

var rootCmd = &cobra.Command{
	Use:   "ton-exporter",
	Short: "Prometheus exporter for TON masterchain info.",
	RunE:  TonExporterCmd,
}

var allowedAPIVersions = map[string]struct{}{
	"v2": {},
	"v3": {},
}

func init() {
	rootCmd.Flags().String("port", "4000", "port for serving mertics")
	rootCmd.Flags().String("api", "", "api version of ton center")
	rootCmd.Flags().String("rpc", "", "rpc endpoint url of TON")
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func TonExporterCmd(cmd *cobra.Command, args []string) error {
	port, _ := cmd.Flags().GetString("port")
	api, _ := cmd.Flags().GetString("api")
	rpcEndpoint, _ := cmd.Flags().GetString("rpc")

	_, validApi := allowedAPIVersions[api]
	if !validApi {
		return fmt.Errorf("invalid API version: %s. Use --help", api)
	}

	switch api {
	case "v3":
		metrics := apiv3.InitMetrics()
		go apiv3.ScrapeBlockNumber(metrics, rpcEndpoint)
	case "v2":
		metrics := apiv2.InitMetrics()
		go apiv2.ScrapeBlockNumber(metrics, rpcEndpoint)
	}

	exporter.StartHTTPServer(port)
	return nil
}

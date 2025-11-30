package exporter

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/thebigmatchplayer/ton-exporter/utils"
	"go.uber.org/zap"
)

func StartHTTPServer(port string) {
	http.Handle("/metrics", promhttp.Handler())
	utils.Logger.Info("Starting metrics server", zap.String("port", port))

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		utils.Logger.Fatal("HTTP server failed", zap.Error(err))
	}
}

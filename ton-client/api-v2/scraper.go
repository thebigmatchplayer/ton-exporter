package apiv2

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/thebigmatchplayer/ton-exporter/utils"
	"go.uber.org/zap"
)

func ScrapeBlockNumber(metrics *Metrics, rpcEndpoint string) {
	ticker := time.NewTicker(4 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		//send http get request to remote/local ton rpc endpoint
		resp, err := http.Get(rpcEndpoint)
		if err != nil {
			utils.Logger.Warn("Unable to reach TON RPC Endpoint")
		}

		if resp.StatusCode != http.StatusOK {
			utils.Logger.Warn(fmt.Sprintf("TON RPC Endpoint returned non-OK status: %d", resp.StatusCode))
		}

		var result masterchainInfo
		err = json.NewDecoder(resp.Body).Decode(&result)
		if err != nil {
			utils.Logger.Warn(fmt.Sprintf("Failed to decode response body: %v", err))
		}

		utils.Logger.Info("Latest Block Height of TON RPC api/v2", zap.Int("Seq No", int(result.Result.Last.Seqno)))
		metrics.v2_BlockHeight.Set(float64(result.Result.Last.Seqno))
		resp.Body.Close()
	}
}

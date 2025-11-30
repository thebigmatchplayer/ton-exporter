package apiv3

type masterchainInfo struct {
	Last struct {
		Seqno uint64 `json:"seqno"`
	} `json:"last"`
}

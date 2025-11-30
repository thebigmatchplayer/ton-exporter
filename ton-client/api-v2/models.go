package apiv2

type masterchainInfo struct {
	Result struct {
		Last struct {
			Seqno uint64 `json:"seqno"`
		} `json:"last"`
	} `json:"result"`
}

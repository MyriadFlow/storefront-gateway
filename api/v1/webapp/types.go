package webapp

type WebappPayload struct {
	SubgraphId string `json:"subgraphId,omitempty"`
}

type WebappResponse struct {
	SubgraphUrl string `json:"subgraphUrl,omitempty"`
}

package request

type SaveOrderRequest struct {
	Note    string `json:"note,omitempty"`
	Address string `json:"address,omitempty"`
}

type PutOrderRequest struct {
	Status int `json:"status,omitempty"`
}

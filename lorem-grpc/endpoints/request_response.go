package endpoints

// LoremRequest ...
type LoremRequest struct {
	Min int32
	Max int32
}

// LoremResponse ...
type LoremResponse struct {
	Message string `json:"message"`
}

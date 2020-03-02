package endpoints

// GenerateLoremRequest ...
type GenerateLoremRequest struct {
	RequestType string
	Min         int32
	Max         int32
}

// GenerateLoremResponse ...
type GenerateLoremResponse struct {
	Message string `json:"message"`
	Err     string `json:"err,omitempty"`
}

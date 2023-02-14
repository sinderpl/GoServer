package sampleroutes

type Sample struct {
	ID   int64  `json:"ID,omitempty"`
	Name string `json:"name,omitempty"`
}

type RequestBody struct {
	UserName   string `json:"userName,omitempty"`
	SampleName string `json:"sampleName,omitempty"`
}

type StatusResponse struct {
	Message string `json:"message,omitempty"`
	Sample  Sample `json:"sample,omitempty"`
}

type ErrorResponse struct {
	Title   string `json:"title"`
	Message string `json:"message"`
}

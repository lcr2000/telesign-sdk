package telesign

// Requester interface
type Requester interface {
	GetMethod() string
	GetURI() string
	GetPath() string
	GetBody() string
}

// MainResponse returned by telesign API
type MainResponse struct {
	StatusCode  int
	ResourceURI string         `json:"resource_uri"`
	ReferenceID string         `json:"reference_id"`
	Status      StatusResponse `json:"status"`
}

// StatusResponse returned by telesign API
type StatusResponse struct {
	Code        int    `json:"code"`
	UpdatedOn   string `json:"updated_on"`
	Description string `json:"description"`
}

// AdditionalInfo returned by telesign API
type AdditionalInfo struct {
	CodeEntered       string `json:"code_entered"`
	MessagePartsCount int    `json:"message_parts_count"`
}

// Error returned by telesign API
type Error struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

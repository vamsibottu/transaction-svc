package models

// Transaction holds the details of transaction
type Transaction struct {
	OrderIdentifier string `json:"order_identifier,omitempty"`
	Name            string `json:"name,omitempty"`
	Email           string `json:"email,omitempty"`
	ItemDescription string `json:"item_description,omitempty"`
	Quantity        int64  `json:"quantity,omitempty"`
}

// ErrorResponse is used to write back the error response with message
type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

package models

type CreateServerRequest struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Status string `json:"status"`
}

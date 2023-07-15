package models

type ListServerResponse struct {
	Servers []*Server `json:"servers"`
	Total   int       `json:"total"`
}

type CreateServerResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

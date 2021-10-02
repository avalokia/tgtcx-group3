package dictionary

import "time"

type Users struct {
	ID          int64
	Name        string
	Tier        string
	Location    string
	TopCategory string
}

type Coupons struct {
	ID        int64
	Name      string
	Category  string
	StartDate time.Time
	EndDate   time.Time
	Status    string
}

type APIResponse struct {
	Data  interface{} `json:"data"`
	Error string      `json:"error_message"`
}

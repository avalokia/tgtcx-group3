package dictionary

// import "time"

type Users struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Tier        string `json:"tier"`
	Location    string `json:"location"`
	TopCategory string `json:"top_category"`
}

type Coupons struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Category  string `json:"category"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Status    string `json:"status"`
}

type APIResponse struct {
	Data  interface{} `json:"data"`
	Error string      `json:"error_message"`
}

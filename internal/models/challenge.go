package models

type Challenge struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	TargetValue float64   `json:"target_value"`
	Progress    int       `json:"progress"`
	Comment     []Comment `json:"comments"`
}

type User struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Active bool   `json:"active"`
}

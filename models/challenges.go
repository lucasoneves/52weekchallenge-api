package models

type Challenge struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	TargetValue float64 `json:"target_value"`
	Progress    int     `json:"progress"`
}

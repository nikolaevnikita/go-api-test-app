package models

type Task struct {
	TID         string    `json:"tid"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
}

func (t Task) ID() string {
	return t.TID
}
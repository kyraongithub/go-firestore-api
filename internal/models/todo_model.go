package models

type Todo struct {
	ID        string `json:"id,omitempty"`
	Item      string `json:"item" firestore:"item"`
	Completed bool   `json:"completed" firestore:"completed"`
}

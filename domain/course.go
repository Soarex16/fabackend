package domain

// Course - represents courses constructed from exercises
type Course struct {
	Label       string   `json:"label,omitempty"`
	Description string   `json:"description,omitempty"`
	Exercises   []string `json:"exercises,omitempty"`
	Completed   bool     `json:"completed,omitempty"`
}

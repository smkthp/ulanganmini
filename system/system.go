package system

type Task struct {
	ID    int    `json:"taskID"`
	Title string `json:"title"`
}
type Choice struct {
	ID    int    `json:"choiceID"`
	Label string `json:"label"`
	Body  string `json:"body"`
}

type Question struct {
	Type   int    `json:"type"`
	Body   string `json:"body"`
	Task   *Task
	Choice *Choice
}

package system

type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
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

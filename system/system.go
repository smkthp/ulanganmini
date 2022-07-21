package system

type Task struct {
	ID    int
	Title string
}
type Choice struct {
	ID    int
	Label string
	Body  string
}

type Question struct {
	Type   int
	Body   string
	Task   *Task
	Choice *Choice
}

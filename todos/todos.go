package todos

type Status int

const (
	TODO Status = iota + 1
	INPROGRESS
	DONE
)

type Todo struct {
	Id    int
	Title string
	State Status
}

func New(title string, state Status) Todo {
	return Todo{
		Id:    5, //TODO: this should be generated randomly!
		Title: title,
		State: state,
	}
}

func (todo *Todo) Update(new_todo Todo) {
	todo.Title = new_todo.Title
	todo.State = new_todo.State
}

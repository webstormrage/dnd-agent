package domain

type Stack struct {
	Push   *Command     `json:"push"`
	Pop    *interface{} `json:"pop"`
	Target string       `json:"target"`
}

func Resolve(command *Command) {
	val := interface{}(true)
	command.Stack.Pop = &val
}

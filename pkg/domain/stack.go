package domain

type Stack struct {
	Push   *Command
	Pop    *interface{}
	Target string
}

func Resolve(command *Command) {
	val := interface{}(true)
	command.Stack.Pop = &val
}

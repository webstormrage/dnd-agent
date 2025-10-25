package handlers

import (
	"dnd-agent/pkg/domain"
	"fmt"
)

func Default(w *domain.World, command *domain.Command) {
	fmt.Printf("Неизвестная команда: %v\n", command.Procedure)
	domain.Resolve(command)
}

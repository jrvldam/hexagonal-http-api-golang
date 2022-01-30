package creating

import (
	"context"
	"errors"

	"github.com/jrvldam/hexagonal-http-api-golang/kit/command"
)

const CourseCommandType command.Type = "command.creating.course"

// CourseCommand is the command dispatched to create a new course
type CourseCommand struct {
	id       string
	name     string
	duration string
}

// NewCourseCommand creates a new CourseCommand
func NewCourseCommand(id, name, duration string) CourseCommand {
	return CourseCommand{
		id,
		name,
		duration,
	}
}

func (cc CourseCommand) Type() command.Type {
	return CourseCommandType
}

// CourseCommandHandler is the command handler responsible for creating courses
type CourseCommandHandler struct {
	service CourseService
}

// NewCourseCommandHandler initialize a new CourseCommandHandler
func NewCourseCommandHandler(service CourseService) CourseCommandHandler {
	return CourseCommandHandler{
		service,
	}
}

// Handle implements the command.Handler interface
func (cch CourseCommandHandler) Handle(ctx context.Context, cmd command.Command) error {
	createCourseCmd, ok := cmd.(CourseCommand)

	if !ok {
		return errors.New("Unexpected command")
	}

	return cch.service.CreateCourse(ctx, createCourseCmd.id, createCourseCmd.name, createCourseCmd.duration)
}

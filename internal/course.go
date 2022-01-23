package mooc

import "context"

type CourseRepository interface {
	Save(ctx context.Context, course Course) error
}

//go:generate mockery --case=snake -- outpkg=storagemocks --putput=platform/storage/storagemocks --name=CourseRepository

// Course is the data structure that represents a course
type Course struct {
	id       string
	name     string
	duration string
}

// NewCourse creates a new course
func NewCourse(id, name, duration string) Course {
	return Course{
		id,
		name,
		duration,
	}
}

func (c Course) ID() string {
	return c.id
}

func (c Course) Name() string {
	return c.name
}

func (c Course) Duration() string {
	return c.duration
}

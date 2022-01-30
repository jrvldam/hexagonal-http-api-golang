package creating

import (
	"context"

	mooc "github.com/jrvldam/hexagonal-http-api-golang/internal"
)

// CourseService is the default CourseService interface
// implementation returned by creating.NewCourseService
type CourseService struct {
	courseRepository mooc.CourseRepository
}

// NewCourseService returns the default Service interface implementation
func NewCourseService(courseRepository mooc.CourseRepository) CourseService {
	return CourseService{
		courseRepository,
	}
}

// CreateCourse implements the creating.CourseService interface
func (cs CourseService) CreateCourse(ctx context.Context, id, name, duration string) error {
	course, err := mooc.NewCourse(id, name, duration)

	if err != nil {
		return err
	}

	return cs.courseRepository.Save(ctx, course)
}

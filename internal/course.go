package mooc

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type CourseRepository interface {
	Save(ctx context.Context, course Course) error
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=CourseRepository

// Course is the data structure that represents a course
type Course struct {
	id       CourseID
	name     CourseName
	duration CourseDuration
}

// NewCourse creates a new course
func NewCourse(id, name, duration string) (Course, error) {
	idVO, err := NewCourseID(id)

	if err != nil {
		return Course{}, err
	}

	nameVO, err := NewCourseName(name)

	if err != nil {
		return Course{}, err
	}

	durationVO, err := NewCourseDuration(duration)

	if err != nil {
		return Course{}, err
	}

	return Course{
		id:       idVO,
		name:     nameVO,
		duration: durationVO,
	}, nil
}

func (c Course) ID() CourseID {
	return c.id
}

func (c Course) Name() CourseName {
	return c.name
}

func (c Course) Duration() CourseDuration {
	return c.duration
}

var ErrInvalidCourseID = errors.New("Invalid course ID")

type CourseID struct {
	value string
}

// NewCourseID instantiate the VO for CourseID
func NewCourseID(value string) (CourseID, error) {
	v, err := uuid.Parse(value)

	if err != nil {
		return CourseID{}, fmt.Errorf("%w: %s", ErrInvalidCourseID, value)
	}

	return CourseID{value: v.String()}, nil
}

// String type converts the CourseID into string
func (id CourseID) String() string {
	return id.value
}

var ErrEmptyCourseName = errors.New("The field Course Name can not be empty")

// CourseName represents the course name
type CourseName struct {
	value string
}

// NewCourseName instantiate VO for CourseName
func NewCourseName(value string) (CourseName, error) {
	if value == "" {
		return CourseName{}, ErrEmptyCourseName
	}

	return CourseName{value}, nil
}

// String type converts the CourseName into string
func (name CourseName) String() string {
	return name.value
}

var ErrEmptyDuration = errors.New("The field Duration can not be empty")

// CourseDuration represents the course duration
type CourseDuration struct {
	value string
}

func NewCourseDuration(value string) (CourseDuration, error) {
	if value == "" {
		return CourseDuration{}, ErrEmptyDuration
	}

	return CourseDuration{value}, nil
}

// String type converts the CourseDuration into string
func (duration CourseDuration) String() string {
	return duration.value
}

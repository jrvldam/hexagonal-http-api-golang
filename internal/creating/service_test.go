package creating

import (
	"context"
	"errors"
	"testing"

	mooc "github.com/jrvldam/hexagonal-http-api-golang/internal"
	"github.com/jrvldam/hexagonal-http-api-golang/internal/platform/storage/storagemocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_CourseService_CreateCourse_RepositoryError(t *testing.T) {
	// Given
	courseID, courseName, courseDuration := "37a0f027-15e6-47cc-a5d2-64183281087e", "Test course", "10 months"

	course, err := mooc.NewCourse(courseID, courseName, courseDuration)
	require.NoError(t, err)

	courseRepositoryMock := new(storagemocks.CourseRepository)
	courseRepositoryMock.On("Save", mock.Anything, course).Return(errors.New("Something unexpected happened"))

	courseService := NewCourseService(courseRepositoryMock)

	// When
	err = courseService.CreateCourse(context.Background(), courseID, courseName, courseDuration)

	// Then
	courseRepositoryMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_CourseService_CreateCourse_Succeed(t *testing.T) {
	// Given
	courseID, courseName, courseDuration := "37a0f027-15e6-47cc-a5d2-64183281087e", "Test course", "10 months"

	course, err := mooc.NewCourse(courseID, courseName, courseDuration)
	require.NoError(t, err)

	courseRepositoryMock := new(storagemocks.CourseRepository)
	courseRepositoryMock.On("Save", mock.Anything, course).Return(nil)

	courseService := NewCourseService(courseRepositoryMock)

	// When
	err = courseService.CreateCourse(context.Background(), courseID, courseName, courseDuration)

	// Then
	courseRepositoryMock.AssertExpectations(t)
	assert.NoError(t, err)
}

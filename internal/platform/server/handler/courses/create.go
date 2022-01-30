package courses

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	mooc "github.com/jrvldam/hexagonal-http-api-golang/internal"
	"github.com/jrvldam/hexagonal-http-api-golang/internal/creating"
)

type createRequest struct {
	ID       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Duration string `json:"duration" binding:"required"`
}

// CreateHandler returns an HTTP handler for courses creation
func CreateHandler(creatingCourseService creating.CourseService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req createRequest

		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		err := creatingCourseService.CreateCourse(c, req.ID, req.Name, req.Duration)

		if err != nil {
			switch {
			case errors.Is(err, mooc.ErrInvalidCourseID), errors.Is(err, mooc.ErrEmptyCourseName), errors.Is(err, mooc.ErrInvalidCourseID):
				c.JSON(http.StatusBadRequest, err.Error())
				return
			default:
				c.JSON(http.StatusInternalServerError, err.Error())
				return
			}
		}

		c.Status(http.StatusCreated)
	}
}

package courses

import (
	"net/http"

	"github.com/gin-gonic/gin"
	mooc "github.com/jrvldam/hexagonal-http-api-golang/internal"
)

type createRequest struct {
	ID       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Duration string `json:"duration" binding:"required"`
}

// CreateHandler returns an HTTP handler for courses creation
func CreateHandler(courseRepository mooc.CourseRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req createRequest

		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		course := mooc.NewCourse(req.ID, req.Name, req.Duration)

		if err := courseRepository.Save(c, course); err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.Status(http.StatusCreated)
	}
}

package courses

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jrvldam/hexagonal-http-api-golang/internal/creating"
	"github.com/jrvldam/hexagonal-http-api-golang/internal/platform/storage/storagemocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHandler_Create(t *testing.T) {
	repositoryMock := new(storagemocks.CourseRepository)
	repositoryMock.On("Save", mock.Anything, mock.Anything).Return(nil)

	createCourseSrv := creating.NewCourseService(repositoryMock)

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/courses", CreateHandler(createCourseSrv))

	t.Run("Given an invalid request it returns 400", func(t *testing.T) {
		createCourseReq := createRequest{
			ID:   "a2d3a622-6d78-4ec9-b0d5-224b2ec0bef6",
			Name: "Demo Course",
		}

		body, err := json.Marshal(createCourseReq)
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/courses", bytes.NewBuffer(body))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})

	t.Run("Given a valid request it returns 201", func(t *testing.T) {
		createCourseReq := createRequest{
			ID:       "8a1c5cdc-ba57-445a-994d-aa412d23723f",
			Name:     "Demo Course",
			Duration: "10 months",
		}

		json, err := json.Marshal(createCourseReq)
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/courses", bytes.NewBuffer(json))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusCreated, res.StatusCode)

	})
}

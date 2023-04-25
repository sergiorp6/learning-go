package postad

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	. "github.com/stretchr/testify/mock"
	. "github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_4/internal/ad/application/postad"
	"github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_4/internal/ad/domain"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPostAdHandler_Success(t *testing.T) {
	id := "1d606f23-1e10-4c72-8dcd-ba3aa1098da8"
	payload := httpPostAdRequest{
		"Test Ad",
		"This is a test ad",
		9.99,
	}
	mockService := NewMockPostAdService(false)
	expectedRequest := NewRequest(id, payload.Title, payload.Description, payload.Price)
	mockService.On("Execute", expectedRequest).Return(nil)
	router := gin.Default()
	router.PUT("/ads/:id", Handler(mockService))
	marshaledBody, _ := json.Marshal(payload)

	req, _ := http.NewRequest("PUT", "/ads/"+id, bytes.NewBuffer(marshaledBody))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	mockService.AssertExpectations(t)
	assert.Equal(t, http.StatusCreated, resp.Code)
}

func TestPostAdHandler_LargeTitle(t *testing.T) {
	id := "1d606f23-1e10-4c72-8dcd-ba3aa1098da8"
	payload := httpPostAdRequest{
		"Lorem ipsum dolor sit amet, consectetuer adipiscing",
		"This is a test ad",
		9.99,
	}
	mockService := NewMockPostAdService(true)
	expectedRequest := NewRequest(id, payload.Title, payload.Description, payload.Price)
	mockService.On("Execute", expectedRequest).Return(domain.ErrTitleTooLong)
	router := gin.Default()
	router.PUT("/ads/:id", Handler(mockService))
	marshaledBody, _ := json.Marshal(payload)

	req, _ := http.NewRequest("PUT", "/ads/"+id, bytes.NewBuffer(marshaledBody))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	mockService.AssertExpectations(t)
	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

type MockPostAdService struct {
	Mock
	returnError bool
}

func NewMockPostAdService(returnError bool) MockPostAdService {
	return MockPostAdService{returnError: returnError}
}

func (m MockPostAdService) Execute(request Request) error {
	args := m.Called(request)
	if m.returnError {
		return args.Error(0)
	}
	return nil
}

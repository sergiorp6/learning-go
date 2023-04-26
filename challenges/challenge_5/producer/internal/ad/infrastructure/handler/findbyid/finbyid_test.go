package findbyid

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	. "github.com/stretchr/testify/mock"
	. "github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_5/producer/internal/ad/application/findbyid"
	. "github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_5/producer/internal/ad/domain"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type MockFindByIdService struct {
	Mock
	filledWithData bool
}

func NewMockFindByIdService(filledWithData bool) MockFindByIdService {
	return MockFindByIdService{filledWithData: filledWithData}
}

func (m MockFindByIdService) Execute(request Request) (*Ad, error) {
	args := m.Called(request)
	if m.filledWithData {
		return args.Get(0).(*Ad), args.Error(1)
	}
	return nil, nil
}

func TestFindByIdHandler_Success(t *testing.T) {
	mockService := NewMockFindByIdService(true)
	expected, _ := NewAd(
		"1d606f23-1e10-4c72-8dcd-ba3aa1098da8",
		"Test Ad",
		"Test description",
		9.99,
		time.Now(),
	)
	mockService.On("Execute", NewRequest(expected.Id().String())).Return(&expected, nil)
	router := gin.Default()
	router.GET("/ddadsdfsdg/:id", Handler(mockService))

	req, _ := http.NewRequest("GET", "/ddadsdfsdg/"+expected.Id().String(), nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	mockService.AssertExpectations(t)
	assert.Equal(t, http.StatusOK, resp.Code)
	expectedJSON := `{"id":"` + expected.Id().String() + `","title":"Test Ad","description":"Test description","price":9.99}`
	assert.JSONEq(t, expectedJSON, resp.Body.String())
}

func TestFindByIdHandler_NotFound(t *testing.T) {
	mockService := NewMockFindByIdService(false)
	expected, _ := NewAd(
		"1d606f23-1e10-4c72-8dcd-ba3aa1098da8",
		"Test Ad",
		"Test description",
		9.99,
		time.Now(),
	)
	mockService.On("Execute", NewRequest(expected.Id().String())).Return(nil, nil)
	router := gin.Default()
	router.GET("/:id", Handler(mockService))

	req, _ := http.NewRequest("GET", "/"+expected.Id().String(), nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
}

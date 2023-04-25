package getadslist

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	. "github.com/stretchr/testify/mock"
	. "github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_4/internal/ad/application/getadslist"
	. "github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_4/internal/ad/domain"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type MockGetAdsListService struct {
	Mock
}

func NewMockGetAdsListService() MockGetAdsListService {
	return MockGetAdsListService{}
}

func (m MockGetAdsListService) Execute(request Request) []Ad {
	args := m.Called(request)

	return args.Get(0).([]Ad)
}

func TestGetAdsListHandler_Success(t *testing.T) {
	mockService := NewMockGetAdsListService()
	const numberOfElements = 1
	var idString = uuid.NewString()
	var expected []Ad
	for range [numberOfElements]struct{}{} {
		ad, _ := NewAd(
			idString,
			"A title",
			"A description",
			9.99,
			time.Now(),
		)
		expected = append(expected, ad)
	}
	mockService.On("Execute", NewRequest(numberOfElements)).Return(expected)
	router := gin.Default()
	router.GET("/ads", Handler(mockService))

	req, _ := http.NewRequest("GET", "/ads?limit=1", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	mockService.AssertExpectations(t)
	assert.Equal(t, http.StatusOK, resp.Code)
	expectedJSON := `{"data":[{"id":"` + idString + `","title":"A title","description":"A description","price":9.99}]}`
	assert.JSONEq(t, expectedJSON, resp.Body.String())
}

func TestGetAdsListHandler_BadRequest(t *testing.T) {
	mockService := NewMockGetAdsListService()
	router := gin.Default()
	router.GET("/ads", Handler(mockService))

	req, _ := http.NewRequest("GET", "/ads?limit=aRandomString", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
	expectedJSON := `{"error":"Limit must be a number"}`
	assert.JSONEq(t, expectedJSON, resp.Body.String())
}

package precode

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenRequestIsCorrectTheAnswerIsNotEmpty(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count="+strconv.Itoa(totalCount)+"&city=moscow", nil)
	// здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.NotEmpty(t, responseRecorder.Body.String())
}

func TestMainHandlerWhenCityIsWrong(t *testing.T) {
	wrongCity := "wrongcity"
	correctError := "wrong city value"
	req := httptest.NewRequest("GET", "/cafe?count=3&city="+wrongCity, nil)
	// здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	assert.Equal(t, correctError, responseRecorder.Body.String())
}

func TestMainHandlerWhenCountMoreThanTheNumberOfCafes(t *testing.T) {
	totalCount := 4
	bigCount := 100
	req := httptest.NewRequest("GET", "/cafe?count="+strconv.Itoa(bigCount)+"&city=moscow", nil)
	// здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code)
	cafes := strings.Split(responseRecorder.Body.String(), ",")
	assert.Len(t, cafes, totalCount)
}

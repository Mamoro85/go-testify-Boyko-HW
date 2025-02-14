package main

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
    "github.com/stretchr/testify/assert"
)

// Тест 1: Корректный запрос
func TestMainHandlerCorrectRequest(t *testing.T) {
    req, err := http.NewRequest("GET", "/cafe?city=moscow&count=2", nil)
    if err != nil {
        t.Fatal(err)
    }

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    assert.Equal(t, http.StatusOK, responseRecorder.Code) // Поменяли if на assert
    assert.NotEmpty(t, responseRecorder.Body.String()) // Проверяем, что тело ответа не пустое
}

// Тест 2: Город не поддерживается (проверка корректности ошибки)
func TestMainHandlerWrongCityValue(t *testing.T) {
    req, err := http.NewRequest("GET", "/cafe?city=london&count=2", nil)
    if err != nil {
        t.Fatal(err)
    }

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    assert.Equal(t, http.StatusBadRequest, responseRecorder.Code) // Поменяли if на assert
    assert.Equal(t, "wrong city value\n", responseRecorder.Body.String()) // Проверка на точный ответ
}

// Тест 3: Запрашивается больше кафе, чем есть в списке
func TestMainHandlerCountMoreThanTotal(t *testing.T) {
    req, err := http.NewRequest("GET", "/cafe?city=moscow&count=10", nil)
    if err != nil {
        t.Fatal(err)
    }

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    assert.Equal(t, http.StatusOK, responseRecorder.Code) // Поменяли if на assert
    expected := "Мир кофе,Сладкоежка,Кофе и завтраки,Сытый студент"
    assert.Equal(t, expected, strings.TrimSpace(responseRecorder.Body.String())) // Проверка на точное значение
}

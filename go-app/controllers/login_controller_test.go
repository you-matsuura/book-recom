package controller

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestPostSignUp(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()

	r := gin.Default()
	r.POST("/signup", PostSignUp)
	req, _ := http.NewRequest("POST", "/signup", strings.NewReader(""))
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, but got %v", w.Code)
	}

	expectedBody := `{"message":"セッションに保存完了"}`
	if w.Body.String() != expectedBody {
		t.Errorf("Expected body %v, but got %v", expectedBody, w.Body.String())
	}
}

// func TestGetSession(t *testing.T) {
// 	gin.SetMode(gin.TestMode)
// 	w := httptest.NewRecorder()

// 	r := gin.Default()
// 	r.GET("/session", GetSession)
// 	req, _ := http.NewRequest("GET", "/session", nil)
// 	r.ServeHTTP(w, req)

// 	if w.Code != http.StatusOK {
// 		t.Errorf("Expected status 200, but got %v", w.Code)
// 	}

// 	expectedBody := `{"redisValue":"セッションバリュー"}`
// 	if w.Body.String() != expectedBody {
// 		t.Errorf("Expected body %v, but got %v", expectedBody, w.Body.String())
// 	}
// }

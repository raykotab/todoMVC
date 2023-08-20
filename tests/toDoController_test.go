package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"
	"fmt"
	"path/filepath"
	"github.com/raykotab/todoMVC/controllers"
)

func TestCreateHandler (t *testing.T) {
	recorderMock := httptest.NewRecorder()
	requestMock := httptest.NewRequest("POST", "/create", nil)

	controllers.CreateHandler(recorderMock, requestMock)

	if recorderMock.Code != http.StatusSeeOther {
		t.Errorf("Expected status code %d but got %d", http.StatusSeeOther, recorderMock.Code)
	}
}

func TestIndexHandler(t *testing.T) {
    requestMock := httptest.NewRequest("GET", "/", nil)
    recorderMock := httptest.NewRecorder()

	baseDir, err := filepath.Abs("../")
	if err != nil {
		t.Fatal(err)
	}

	tmplPath := filepath.Join(baseDir, "views/templates/index.html")
	fmt.Println(tmplPath, "hola")
	controllers.IndexHandler(recorderMock, requestMock, tmplPath)
    // http.HandlerFunc(controllers.IndexHandler).ServeHTTP(recorderMock, requestMock)

    if recorderMock.Code != http.StatusOK {
        t.Errorf("Expected status code %d but got %d", http.StatusOK, recorderMock.Code)
    }

    expectedContent := "To Do List"
    responseBody := recorderMock.Body.String()
	fmt.Println(responseBody)
    if !strings.Contains(responseBody, expectedContent) {
        t.Errorf("Response body does not contain expected content: %s", expectedContent)
    }
}

func TestToggleCompleteHandler(t *testing.T) {
	recorderMock := httptest.NewRecorder()
	requestMock := httptest.NewRequest("POST", "/toggle-complete", nil)

	controllers.ToggleCompleteHandler(recorderMock, requestMock)

	if recorderMock.Code != http.StatusSeeOther {
		t.Errorf("Expected status code %d but got %d", http.StatusSeeOther, recorderMock.Code)
	}
}

func TestDeleteHandler(t *testing.T) {
	recorderMock := httptest.NewRecorder()
	requestMock := httptest.NewRequest("POST", "/delete", nil)

	controllers.DeleteHandler(recorderMock, requestMock)

	if recorderMock.Code != http.StatusSeeOther {
		t.Errorf("Expected status code %d but got %d", http.StatusSeeOther, recorderMock.Code)
	}
} 
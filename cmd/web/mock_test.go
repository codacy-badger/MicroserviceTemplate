package main

import (
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/archit-p/MicroserviceTemplate/pkg/models"
)

type SampleMockFail struct{}

func (s *SampleMockFail) Insert(content string) (string, error) {
	return "", errors.New("Database error")
}

func (s *SampleMockFail) Get(id string) (*models.Sample, error) {
	return nil, errors.New("Database error")
}

func (s *SampleMockFail) Update(id string, content string) (int64, error) {
	return 0, errors.New("Database error")
}

func (s *SampleMockFail) Delete(id string) (int64, error) {
	return 0, errors.New("Database error")
}

func NewMockFailApplication() *application {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile)

	mockModel := &SampleMockFail{}

	app := newApplication(infoLog, errorLog, mockModel)

	return app
}

type SampleMockPass struct{}

func (s *SampleMockPass) Insert(content string) (string, error) {
	return "1234", nil
}

func (s *SampleMockPass) Get(id string) (*models.Sample, error) {
	res := models.Sample{
		Content: "abcd",
		Created: time.Now(),
		Deleted: false,
	}

	return &res, nil
}

func (s *SampleMockPass) Update(id string, content string) (int64, error) {
	return 0, nil
}

func (s *SampleMockPass) Delete(id string) (int64, error) {
	return 0, nil
}

func NewMockPassApplication() *application {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile)

	mockModel := &SampleMockPass{}

	app := newApplication(errorLog, infoLog, mockModel)

	return app
}

func checkResponeCode(t *testing.T, received, expected int) {
	if received != expected {
		t.Errorf("Handler returned wrong status code: got %v want %v", received, expected)
	}
}

func executeRequest(req *http.Request, handler http.Handler) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	return rr
}

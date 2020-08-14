package main

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/archit-p/MicroserviceTemplate/pkg/dto"
)

func TestUnit_createSample(t *testing.T) {
	app := NewMockPassApplication()

	s := &dto.Sample{
		Content: "hello",
	}
	payload := new(bytes.Buffer)
	err := s.ToJSON(payload)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodPost, "/sample/create", payload)
	if err != nil {
		t.Fatal(err)
	}

	rr := executeRequest(req, app.routes())
	checkResponeCode(t, rr.Code, http.StatusOK)
}

func TestUnit_createSample_Empty(t *testing.T) {
	app := NewMockPassApplication()

	req, err := http.NewRequest(http.MethodPost, "/sample/create", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := executeRequest(req, app.routes())
	checkResponeCode(t, rr.Code, http.StatusBadRequest)
}

func TestUnit_getSample(t *testing.T) {
	app := NewMockPassApplication()

	req, err := http.NewRequest(http.MethodGet, "/sample/1234", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := executeRequest(req, app.routes())
	checkResponeCode(t, rr.Code, http.StatusOK)
}

func TestUnit_deleteSample(t *testing.T) {
	app := NewMockPassApplication()

	req, err := http.NewRequest(http.MethodDelete, "/sample/1234", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := executeRequest(req, app.routes())
	checkResponeCode(t, rr.Code, http.StatusOK)
}

func TestUnit_updateSample_empty(t *testing.T) {
	app := NewMockPassApplication()

	req, err := http.NewRequest(http.MethodPut, "/sample/1234", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := executeRequest(req, app.routes())
	checkResponeCode(t, rr.Code, http.StatusBadRequest)
}

func TestUnit_createSample_fail(t *testing.T) {
	app := NewMockFailApplication()

	s := &dto.Sample{
		Content: "hello",
	}
	payload := new(bytes.Buffer)
	err := s.ToJSON(payload)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodPost, "/sample/create", payload)
	if err != nil {
		t.Fatal(err)
	}

	rr := executeRequest(req, app.routes())
	checkResponeCode(t, rr.Code, http.StatusInternalServerError)
}

func TestUnit_getSample_fail(t *testing.T) {
	app := NewMockFailApplication()

	req, err := http.NewRequest(http.MethodGet, "/sample/1234", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := executeRequest(req, app.routes())
	checkResponeCode(t, rr.Code, http.StatusBadRequest)
}

func TestUnit_deleteSample_fail(t *testing.T) {
	app := NewMockFailApplication()

	req, err := http.NewRequest(http.MethodDelete, "/sample/1234", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := executeRequest(req, app.routes())
	checkResponeCode(t, rr.Code, http.StatusBadRequest)
}

func TestUnit_updateSample(t *testing.T) {
	app := NewMockPassApplication()

	s := &dto.Sample{
		Content: "hello",
	}
	payload := new(bytes.Buffer)
	err := s.ToJSON(payload)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodPut, "/sample/1234", payload)
	if err != nil {
		t.Fatal(err)
	}

	rr := executeRequest(req, app.routes())
	checkResponeCode(t, rr.Code, http.StatusOK)
}

func TestUnit_updateSample_fail(t *testing.T) {
	app := NewMockFailApplication()

	s := &dto.Sample{
		Content: "hello",
	}
	payload := new(bytes.Buffer)
	err := s.ToJSON(payload)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodPut, "/sample/1234", payload)
	if err != nil {
		t.Fatal(err)
	}

	rr := executeRequest(req, app.routes())
	checkResponeCode(t, rr.Code, http.StatusBadRequest)
}

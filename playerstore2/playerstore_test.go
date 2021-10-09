package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

type testcase struct {
	name       string
	score      int
	httpStatus int
}

var testcases []*testcase

func initTestcases() {
	for name, score := range defaultHttpHandler.store {
		testcases = append(testcases, &testcase{name, score, http.StatusOK})
	}
}

func TestGetHandler(t *testing.T) {
	testcases = append(testcases, &testcase{"UNKOWN", 0, http.StatusNotFound})
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			getPlayerScore(t, testcase.name, testcase.score, testcase.httpStatus)
		})
	}
}

func TestPostHandler(t *testing.T) {
	testcases = append(testcases, &testcase{"UNKOWN", 0, http.StatusOK})
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			postPlayerName(t, testcase.name, testcase.score, testcase.httpStatus)
		})
	}
}

func getPlayerScore(t *testing.T, name string, score int, status int) {
	t.Helper()
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	if err != nil {
		t.Error(err)
	}
	// use testhttp.ResponseRecorder as http.Response.Writer
	recorder := httptest.NewRecorder()
	defaultHttpHandler.ServeHTTP(recorder, req)
	body := recorder.Body.String()
	gotScore, err := strconv.Atoi(body)
	if err != nil {
		t.Errorf("got %v want %d", err, score)
	}
	if gotScore != score {
		t.Errorf("got %d want %d", gotScore, score)
	}
	if recorder.Code != status {
		t.Errorf("got %d want %d", recorder.Code, status)
	}
}

func postPlayerName(t *testing.T, name string, score int, status int) {
	t.Helper()
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	if err != nil {
		t.Log(err)
	}
	recoder := httptest.NewRecorder()
	defaultHttpHandler.ServeHTTP(recoder, req)
	body := recoder.Body.String()
	gotScore, err := strconv.Atoi(body)
	if err != nil {
		t.Errorf("got %v want %d", err, score)
	}
	t.Logf("got %d want %d", gotScore, score)
	if score++; gotScore != score {
		t.Errorf("got %d want %d", gotScore, score)
	}
	if recoder.Code != status {
		t.Errorf("got %d want %d", recoder.Code, status)
	}
}

func TestMain(m *testing.M) {
	initTestcases()
	m.Run()
}

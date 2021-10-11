package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"testing"
)

var testMux = NewPlayerServeMux()

type testcase struct {
	name       string
	score      int
	httpStatus int
}

func initTestcases() (testcases []*testcase) {
	for name, score := range testMux.store {
		testcases = append(testcases, &testcase{name, score, http.StatusOK})
	}
	return
}

func TestGetScoreHandler(t *testing.T) {
	testcases := append(initTestcases(), &testcase{"UNKOWN", 0, http.StatusNotFound})
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			getPlayerScore(t, testcase.name, testcase.score, testcase.httpStatus)
		})
	}
}

func TestPostScoreHandler(t *testing.T) {
	testcases := append(initTestcases(), &testcase{"UNKOWN", 0, http.StatusOK})
	for _, testcase := range testcases {
		t.Log(testcase)
		t.Run(testcase.name, func(t *testing.T) {
			postPlayerScore(t, testcase.name, testcase.score, testcase.httpStatus)
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
	testMux.ServeHTTP(recorder, req)
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

func postPlayerScore(t *testing.T, name string, score int, status int) {
	t.Helper()
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	if err != nil {
		t.Log(err)
	}
	recoder := httptest.NewRecorder()
	testMux.ServeHTTP(recoder, req)
	body := recoder.Body.String()
	gotScore, err := strconv.Atoi(body)
	if err != nil {
		t.Errorf("got %v want %d", err, score)
	}
	if score++; gotScore != score {
		t.Errorf("got %d want %d", gotScore, score)
	}
	if recoder.Code != status {
		t.Errorf("got %d want %d", recoder.Code, status)
	}
}

func TestPlayerLeague(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/league", nil)
	if err != nil {
		t.Error(err)
	}
	recorder := httptest.NewRecorder()
	testMux.ServeHTTP(recorder, req)
	res := map[string]int{}
	json.Unmarshal(recorder.Body.Bytes(), &res)
	if !reflect.DeepEqual(res, testMux.store) {
		t.Errorf("got %v want %v", res, testMux.store)
	}
}

package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// mock resp writer by implementing all funcs of resp writer
type mockRespWriter bytes.Buffer

func (*mockRespWriter) Header() http.Header        { return nil }
func (*mockRespWriter) WriteHeader(statusCode int) {}

func (w *mockRespWriter) Write(data []byte) (int, error) {
	// cast type to *bytes.Buffer to avoid call w.Write directly causing infinite loop
	buf := (*bytes.Buffer)(w)
	return buf.Write(data)
}

func (w *mockRespWriter) Bytes() []byte {
	// cast type to *bytes.Buffer to avoid call w.Write directly causing infinite loop
	buf := (*bytes.Buffer)(w)
	return buf.Bytes()
}

func sendPlayerReq(t *testing.T, name string, score string) {
	t.Helper()
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	if err != nil {
		t.Error(err)
	}
	/* use testhttp.ResponseRecorder */
	w := httptest.NewRecorder()
	defaultServer.ServeHTTP(w, req)
	got := w.Body.String()
	/* use custom mock response writer */
	// w := &mockRespWriter{}
	// defaultServer.ServeHTTP(w, req)
	// got := string(w.Bytes())
	if got != score {
		t.Errorf("got %s want %s", got, score)
	}
}

func TestGetHandler(t *testing.T) {
	tests := []struct {
		name  string
		score string
	}{
		{"TATA", "5"},
		{"RURU", "2"},
		{"OTHERS", ""},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			sendPlayerReq(t, test.name, test.score)
		})
	}
}

/*
func TestPlaystore(t *testing.T) {
	t.Run("test get req", func(t *testing.T) {
		resp, err := http.Get("http://localhost:8090/players/TATA")
		if err != nil {
			t.Error(err)
		}
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Error(err)
		}
		got := string(b)
		want := "5"
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
*/

package main

import (
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		label string
		want  string
		name  string
	}{
		{
			label: "default",
			want:  "Hello World!\n",
			name:  "",
		},
		{
			label: "override",
			want:  "Hello Override!\n",
			name:  "Override",
		},
	}

	for _, test := range tests {
		url := "/"
		if test.name != "" {
			url += "?name=" + test.name
		}

		req := httptest.NewRequest("GET", url, nil)
		rr := httptest.NewRecorder()
		handler(rr, req)

		if got := rr.Body.String(); got != test.want {
			t.Errorf("%s: got %q, want %q", test.label, got, test.want)
		}
	}
}

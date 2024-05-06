package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type test struct {
		input http.Header
		want  string
	}
	tests := []test{
		{input: http.Header{"Authorization": {"ApiKey TestAPIKEY"}}, want: "TestAPIKEY"},
		{input: http.Header{"Authorization": {"ApiKey "}}, want: ""},
		{input: http.Header{"Authorization": {"ApiKey 1234sdg&fjbng"}}, want: "1234sdg&fjbng"},
	}

	for _, tc := range tests {
		got, err := GetAPIKey(tc.input)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got %v", tc.want, got)
		}
	}
}

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestSuccessfulServerRequest(t *testing.T) {
	type input struct {
		url  string
		body string
	}
	go startServer()

	tests := []struct {
		name string
		args input
		want float64
	}{
		{
			"valid input returns expected response",
			input{
				fmt.Sprintf("http://localhost:%d/v1/data", 4000),
				"{\n    \"x\": \"123.12\",\n    \"y\": \"456.56\",\n    \"z\": \"789.89\",\n    \"vel\": \"20.0\"\n}"},
			1389.57,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := http.Post(test.args.url, "application/json", bytes.NewBuffer([]byte(test.args.body)))

			if err != nil {
				t.Errorf("test failed with error from endpoint for %s", test.name)
			}

			if result.StatusCode != http.StatusOK {
				t.Errorf("test failed to return successful http status code %d for %s", result.StatusCode, test.name)
			}

			data, err := io.ReadAll(result.Body)
			if err != nil {
				t.Errorf("test failed due to failure to read body for %s", test.name)
			}

			var responseObject response
			json.Unmarshal(data, &responseObject)

			if responseObject.Loc != test.want {
				t.Errorf("Test failed with expected want : %f not found for %s", test.want, test.name)
			}
		})
	}
}

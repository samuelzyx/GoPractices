package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
)

type MockRounTrip struct {
	RoundTripOutput *http.Response
}

func (m MockRounTrip) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.Header.Get("Authorization") != "Bearer abc" {
		return nil, fmt.Errorf("Wrong authorization header: %s", req.Header.Get("Authorization"))
	}
	return m.RoundTripOutput, nil
}

func TestRoundTrip(t *testing.T) {
	loginResponse := LoginResponse{
		Token: "abc",
	}

	loginResponseBytes, err := json.Marshal(loginResponse)
	if err != nil {
		t.Errorf("Marshal error: %s\n", err)
	}

	myJWTTransport := MyJWTTransport{
		Transport: MockRounTrip{
			RoundTripOutput: &http.Response{
				StatusCode: 200,
			},
		},
		HTTPClient: &MockClient{
			PostResponseOutput: &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(bytes.NewReader(loginResponseBytes)),
			},
		},
		Password: "xyz",
	}

	req := &http.Request{
		Header: make(http.Header), //Initizalice map header
	}

	res, err := myJWTTransport.RoundTrip(req)

	if err != nil {
		t.Fatalf("RounTrip error %s\n", err)
	}

	if res.StatusCode != 200 {
		t.Errorf("StatusCode is not 200, got %d\n", res.StatusCode)
	}
}

package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"
)

// create our own fake client
type MockClient struct {
	ResponseOutput *http.Response
}

func (m *MockClient) Get(url string) (resp *http.Response, err error) {
	return m.ResponseOutput, nil
}

func TestDoGetRequest(t *testing.T) {
	wordsPage := WordsPage{
		Page: Page{"words"},
		Words: Words{
			Input: "abc",
			Words: []string{"a", "b"},
		},
	}

	wordsBytes, err := json.Marshal(wordsPage)
	if err != nil {
		t.Errorf("Marshal error: %s\n", err)
	}

	apiInstance := API{
		Options: Options{},
		Client: &MockClient{
			ResponseOutput: &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(bytes.NewReader(wordsBytes)),
			},
		}, //http.client is with real approach need test server running
	}
	response, err := apiInstance.DoGetRequest("http:localhost:8080/words")
	if err != nil {
		t.Errorf("DoGetRequest error: %s\n", err)
	}

	if response == nil {
		t.Fatal("DoGetRequest is empty")
	}

	if response.GetResponse() != strings.Join([]string{"a", "b"}, ", ") {
		t.Errorf("Unexpected output: %s\n", response.GetResponse())
	}

}

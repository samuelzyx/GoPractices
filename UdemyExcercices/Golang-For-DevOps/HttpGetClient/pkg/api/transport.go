package api

import (
	"net/http"
)

type MyJWTTransport struct {
	Transport  http.RoundTripper
	Token      string
	Password   string
	LoginURL   string
	HTTPClient ClientIFace
}

func (m *MyJWTTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if m.Token == "" {
		if m.Password != "" {
			token, err := doLoginRequest(m.HTTPClient, m.LoginURL, m.Password)
			if err != nil {
				return nil, err
			}
			m.Token = token
		}
	}

	if m.Token != "" {
		req.Header.Add("Authorization", "Bearer "+m.Token)
	}
	return m.Transport.RoundTrip(req)
}

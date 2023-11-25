package api

import "net/http"

type Options struct {
	Password string
	LoginURL string
}

type APIIFace interface {
	DoGetRequest(requstURL string) (Response, error)
}

type API struct {
	Options Options
	Client  http.Client
}

func New(options Options) APIIFace {
	return API{
		Options: options,
		Client: http.Client{
			Transport: &MyJWTTransport{
				Transport: http.DefaultTransport,
				Password:  options.Password,
				LoginURL:  options.LoginURL,
			},
		},
	}
}

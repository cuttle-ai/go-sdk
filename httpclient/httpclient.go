// Copyright 2019 Cuttle.ai. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//Package httpclient contains the http clients that is a wrapper around other 3rd party libraries with
//necessary setup required for the cuttle platform
package httpclient

import (
	"io"
	"net/http"
	"time"

	"github.com/gojektech/heimdall"
	heimdallC "github.com/gojektech/heimdall/httpclient"
)

//Message is the message to be given for successfull response
type Message struct {
	//Message associated with
	Message string
	//Data is payload
	Data interface{}
}

type myHTTPClient struct {
	token    string
	tokenKey string
	domain   string
}

func (c *myHTTPClient) Do(request *http.Request) (*http.Response, error) {
	cookie := http.Cookie{Name: c.tokenKey, Value: c.token, Domain: c.domain, Path: "/"}
	request.AddCookie(&cookie)
	initalTimeout := 2 * time.Millisecond
	maxTimeout := 9 * time.Millisecond
	exponentFactor := float64(2)
	maximumJitterInterval := 2 * time.Millisecond
	backoff := heimdall.NewExponentialBackoff(initalTimeout, maxTimeout, exponentFactor, maximumJitterInterval)
	retrier := heimdall.NewRetrier(backoff)
	timeout := 1000 * time.Millisecond
	client := heimdallC.NewClient(
		heimdallC.WithHTTPTimeout(timeout),
		heimdallC.WithRetrier(retrier),
		heimdallC.WithRetryCount(4),
	)
	return client.Do(request)
}

//Get makes a get request to a api url with retry mechanisms
func Get(domain, url, token, tokenKey string) (*http.Response, error) {
	/*
	 * First we will initalize the client
	 * Then we will send the get request
	 * Then we will return the response
	 */
	//initalizing the client
	client := heimdallC.NewClient(
		heimdallC.WithHTTPClient(&myHTTPClient{
			token:    token,
			tokenKey: tokenKey,
			domain:   domain,
		}),
	)

	//then we will make the request
	res, err := client.Get(url, http.Header{})
	if err != nil {
		return nil, err
	}

	//return the response
	return res, nil
}

//Post makes a post request to a api url with retry mechanisms
func Post(domain, url, token, tokenKey string, body io.Reader) (*http.Response, error) {
	/*
	 * First we will initalize the client
	 * Then we will send the post request
	 * Then we will return the response
	 */
	//initalizing the client
	client := heimdallC.NewClient(
		heimdallC.WithHTTPClient(&myHTTPClient{
			token:    token,
			tokenKey: tokenKey,
			domain:   domain,
		}),
	)

	//then we will make the request
	res, err := client.Post(url, body, http.Header{})
	if err != nil {
		return nil, err
	}

	//return the response
	return res, nil
}

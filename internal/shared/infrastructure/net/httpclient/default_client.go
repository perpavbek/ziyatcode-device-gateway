package httpclient

import (
	"bytes"
	"io"
	"net/http"
)

type DefaultClient struct {
	client *http.Client
}

func New(client *http.Client) *DefaultClient {
	return &DefaultClient{client: client}
}

func (c *DefaultClient) Get(url string, headers map[string]string) (*Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	
	responseBody, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return &Response{
		Headers: resp.Header,
		StatusCode: resp.StatusCode,
		Body: responseBody,
	}, nil
}

func (c *DefaultClient) Post(url string, body []byte, headers map[string]string) (*Response, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	
	responseBody, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return &Response{
		Headers: resp.Header,
		StatusCode: resp.StatusCode,
		Body: responseBody,
	}, nil
}

func (c *DefaultClient) Put(url string, body []byte, headers map[string]string) (*Response, error) {
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	
	responseBody, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return &Response{
		Headers: resp.Header,
		StatusCode: resp.StatusCode,
		Body: responseBody,
	}, nil
}

func (c *DefaultClient) Delete(url string, headers map[string]string) (*Response, error){
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil{
		return nil, err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	
	responseBody, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return &Response{
		Headers: resp.Header,
		StatusCode: resp.StatusCode,
		Body: responseBody,
	}, nil
}
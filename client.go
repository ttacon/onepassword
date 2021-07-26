package onepassword

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type CursorOptions struct {
	ResetCursor *ResetCursor
	CurrCursor  string
}

type Client interface {
	Do(resource, method string, cursorOpts *CursorOptions, rcvr interface{}) error
	Service() EventsService
}

type httpClient struct {
	client *http.Client
	host   string
	token  string
}

func NewHTTPClient(client *http.Client, token, host string) (Client, error) {
	if len(token) == 0 {
		// Get out of here.
		return nil, errors.New("must provide 1password token")
	}

	if client == nil {
		client = http.DefaultClient
	}

	return &httpClient{
		client: client,
		host:   host,
		token:  token,
	}, nil
}

func (hc *httpClient) Do(resource, method string, cursorOpts *CursorOptions, rcvr interface{}) error {
	var body io.Reader
	if cursorOpts != nil {
		var (
			data []byte
			err  error
		)
		if cursorOpts.ResetCursor != nil {
			data, err = json.Marshal(cursorOpts.ResetCursor)
			if err != nil {
				return err
			}
		} else if len(cursorOpts.CurrCursor) == 0 {
			return errors.New("invalid cursor state")
		} else {
			data, err = json.Marshal(map[string]string{
				"cursor": cursorOpts.CurrCursor,
			})
		}
		body = bytes.NewReader(data)
	}

	req, err := http.NewRequest(
		method,
		fmt.Sprintf("https://%s/%s", hc.host, resource),
		body,
	)
	if err != nil {
		return nil
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", hc.token))

	resp, err := hc.client.Do(req)
	if err != nil {
		return err
	}

	if statusCode := resp.StatusCode; statusCode < 200 || statusCode > 299 {
		// DIDNTDO(ttacon): make this a custom error object so that we can
		// return better error information.
		return errors.New("request returned a non-200 response")
	}

	reader := json.NewDecoder(resp.Body)
	if err := reader.Decode(rcvr); err != nil {
		return err
	}
	return resp.Body.Close()
}

func (hc *httpClient) Service() EventsService {
	return &eventsService{
		client: hc,
	}
}

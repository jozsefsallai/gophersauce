package gophersauce

import (
	"errors"
	"io"
)

// Client is a Saucenao API client instance.
type Client struct {
	MaxResults int
	APIKey     string
	APIUrl     string
}

// Settings is a structure that contains the settings of the Gophersauce client
type Settings struct {
	MaxResults int
	APIKey     string
	APIUrl     string
}

// NewClient creates a new Gophersauce API client instance
func NewClient(settings *Settings) (*Client, error) {
	if settings == nil {
		settings = &Settings{
			MaxResults: 6,
		}
	}

	if settings.MaxResults == 0 {
		settings.MaxResults = 6
	}

	if settings.MaxResults < 0 {
		return nil, errors.New("number of max results needs to be greater than 0")
	}

	if len(settings.APIUrl) == 0 {
		settings.APIUrl = "https://saucenao.com/search.php"
	}

	client := &Client{
		APIKey:     settings.APIKey,
		APIUrl:     settings.APIUrl,
		MaxResults: settings.MaxResults,
	}

	return client, nil
}

// SetAPIKey updates the API key of the client
func (c *Client) SetAPIKey(key string) {
	c.APIKey = key
}

// SetAPIUrl updates the API URL of the client
func (c *Client) SetAPIUrl(url string) {
	c.APIUrl = url
}

// SetMaxResults updates the max results property of the client
func (c *Client) SetMaxResults(maxResults int) error {
	if maxResults == 0 {
		maxResults = 6
	}

	if maxResults < 0 {
		return errors.New("number of max results needs to be greater than 0")
	}

	c.MaxResults = maxResults
	return nil
}

// FromURL will look up an image from a given URL address
func (c *Client) FromURL(url string) (*SaucenaoResponse, error) {
	res, err := fetch("url", c, fetchOptions{
		URL: url,
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}

// FromFile will look up an image from a given file path
func (c *Client) FromFile(path string) (*SaucenaoResponse, error) {
	res, err := fetch("file", c, fetchOptions{
		FilePath: path,
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}

// FromReader will look up an image from a given IO reader
func (c *Client) FromReader(reader io.Reader) (*SaucenaoResponse, error) {
	res, err := fetch("reader", c, fetchOptions{
		Reader: reader,
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}

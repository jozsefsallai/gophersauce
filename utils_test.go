package gophersauce

import (
	"testing"
)

func assertionFailedStrings(t *testing.T, expected, actual string) {
	t.Errorf("expected %s, got %s", expected, actual)
}

func assertionFailedInts(t *testing.T, expected, actual int) {
	t.Errorf("expected %d, got %d", expected, actual)
}

func TestGetRequestURL(t *testing.T) {
	var url, expected string

	client, _ := NewClient(nil)
	url = getRequestURL(client)
	expected = "https://saucenao.com/search.php?output_type=2"
	if url != expected {
		assertionFailedStrings(t, expected, url)
	}

	client.SetAPIKey("MyAPIKey")
	url = getRequestURL(client)
	expected = "https://saucenao.com/search.php?api_key=MyAPIKey&output_type=2"
	if url != expected {
		assertionFailedStrings(t, expected, url)
	}

	client.SetAPIKey("")
	client.SetAPIUrl("https://my-saucenao.com/api")
	url = getRequestURL(client)
	expected = "https://my-saucenao.com/api?output_type=2"
	if url != expected {
		assertionFailedStrings(t, expected, url)
	}

	client.SetAPIUrl("https://saucenao.com/search.php")
	client.SetMaxResults(12)
	url = getRequestURL(client)
	expected = "https://saucenao.com/search.php?numres=12&output_type=2"
	if url != expected {
		assertionFailedStrings(t, expected, url)
	}
}

func TestParseIntInterface(t *testing.T) {
	var x interface{}
	var result, expected int

	x = "3"
	result, err := parseIntInterface(x)
	expected = 3

	if err != nil {
		t.Error(err)
	}

	if result != expected {
		assertionFailedInts(t, result, expected)
	}

	x = 12
	result, err = parseIntInterface(x)
	expected = 12

	if err != nil {
		t.Error(err)
	}

	if result != expected {
		assertionFailedInts(t, result, expected)
	}
}

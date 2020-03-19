package gophersauce

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/gabriel-vasile/mimetype"
)

type fetchOptions struct {
	URL      string
	Reader   io.Reader
	FilePath string
}

func escapeQuotes(s string) string {
	var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")
	return quoteEscaper.Replace(s)
}

func createFormFileWithContentType(w *multipart.Writer, fieldname, filename string, contentType string) (io.Writer, error) {
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition",
		fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
			escapeQuotes(fieldname),
			escapeQuotes(filename),
		))
	h.Set("Content-Type", contentType)
	return w.CreatePart(h)
}

func getRequestURL(c *Client) string {
	u, _ := url.Parse("https://saucenao.com/search.php")
	query, _ := url.ParseQuery(u.RawQuery)

	if len(c.APIKey) > 0 {
		query.Add("api_key", c.APIKey)
	}

	if c.MaxResults != 6 {
		query.Add("numres", strconv.Itoa(c.MaxResults))
	}

	query.Add("output_type", "2") // JSON

	u.RawQuery = query.Encode()

	return u.String()
}

func fetch(mode string, c *Client, options fetchOptions) (*SaucenaoResponse, error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	if mode == "url" {
		writer.WriteField("url", options.URL)
	}

	if mode == "file" {
		file, err := os.Open(options.FilePath)
		if err != nil {
			return nil, err
		}

		defer file.Close()

		info, err := file.Stat()
		if err != nil {
			return nil, err
		}

		contentType, err := mimetype.DetectReader(file)
		if err != nil {
			return nil, err
		}

		file.Seek(0, 0)

		part, err := createFormFileWithContentType(writer, "file", info.Name(), contentType.String())
		if err != nil {
			return nil, err
		}

		_, err = io.Copy(part, file)
		if err != nil {
			return nil, err
		}
	}

	if mode == "reader" {
		reader := options.Reader

		// In order to detect the MIME type of the read stream, the contents of the reader
		// has to be copied into a buffer. Otherwise, after figuring out the MIME type, the
		// first few bytes of the stream will be cut off and that will break the stream.
		// I wish I didn't have to do this but SauceNao does MIME checking and io.Reader
		// doesn't have a Seek method like os.File does.
		buffer := &bytes.Buffer{}
		_, err := io.Copy(buffer, reader)
		if err != nil {
			return nil, err
		}

		contentType := mimetype.Detect(buffer.Bytes())

		part, err := createFormFileWithContentType(writer, "file", "upload", contentType.String())
		if err != nil {
			return nil, err
		}

		_, err = io.Copy(part, buffer)
		if err != nil {
			return nil, err
		}
	}

	writer.Close()

	req, err := http.NewRequest("POST", getRequestURL(c), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	responseJSON := &SaucenaoResponse{}
	responseBody, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(responseBody, &responseJSON)
	if err != nil {
		return nil, err
	}

	if len(responseJSON.Header.Message) > 0 {
		return nil, fmt.Errorf("API error: %s", responseJSON.Header.Message)
	}

	return responseJSON, nil
}

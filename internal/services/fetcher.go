package services

import(
	"net/http"
	"io"
	"errors"
	"time"
)



func FetchHTML(url string) (string, error){
	if url == ""{
		return "", errors.New("empty URL provided")
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil{
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		return "", errors.New("failed to fetch URL: " + resp.Status)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil{
		return "", err
	}

	return string(bodyBytes), nil
}
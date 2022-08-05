package _example_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/caicaispace/gohelper/server/http/_example"
)

const serverAddr = "http://127.0.0.1:9601"

func TestTest(t *testing.T) {
	tt := []struct {
		name       string
		method     string
		url        string
		input      string
		want       string
		statusCode int
	}{
		{
			name:       "with test",
			method:     http.MethodGet,
			url:        serverAddr + "/v1/api/test",
			input:      "",
			want:       `{"code":0,"msg":"ok","data":null}`,
			statusCode: http.StatusOK,
		},
		{
			name:       "with pager",
			method:     http.MethodGet,
			url:        serverAddr + "/v1/api/test_pager",
			input:      "",
			want:       `{"code":0,"msg":"ok","data":{"limit":10,"page":1,"total":100}}`,
			statusCode: http.StatusOK,
		},
	}

	go _example.NewServer()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	tickerCount := 0
	for {
		<-ticker.C
		tc := tt[tickerCount]

		fmt.Println("")
		fmt.Println("---------------------------------------------")
		fmt.Println(tc.name, tc.url)
		fmt.Println("")

		body, status, err := request(tc.method, tc.url, tc.input)
		if err != nil {
			t.Fatal(err)
			return
		}

		if status != tc.statusCode {
			t.Errorf("Want status '%d', got '%d'", tc.statusCode, status)
		}

		if strings.TrimSpace(string(body)) != tc.want {
			t.Errorf("Want '%s', got '%s'", tc.want, body)
		}

		tickerCount++
		if tickerCount >= len(tt) {
			break
		}
	}
}

func request(methd, url, inBody string) (string, int, error) {
	req, err := http.NewRequest(methd, url, strings.NewReader(inBody))
	if err != nil {
		return "", 0, err
	}
	if methd != http.MethodGet {
		req.Header.Set("Content-Type", "application/json")
	}
	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()
	var body []byte
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", 0, err
	}
	return string(body), resp.StatusCode, nil
}

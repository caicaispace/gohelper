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
		input      string
		want       string
		statusCode int
	}{
		{
			name:       "with test",
			method:     http.MethodGet,
			input:      "",
			want:       `{"code":0,"msg":"ok","data":null}`,
			statusCode: http.StatusOK,
		},
	}

	go _example.NewServer()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	tickerCount := 0
	for {
		<-ticker.C
		fmt.Println("")
		fmt.Println("---------------------------------------------")
		fmt.Println("")

		tc := tt[tickerCount]
		req, err := http.NewRequest(tc.method, serverAddr+"/v1/api/test", strings.NewReader(tc.input))
		if err != nil {
			t.Fatal(err)
			return
		}
		req.Header.Set("Content-Type", "application/json")
		httpClient := &http.Client{}
		resp, err := httpClient.Do(req)
		if err != nil {
			t.Fatal(err)
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
			return
		}

		if resp.StatusCode != tc.statusCode {
			t.Errorf("Want status '%d', got '%d'", tc.statusCode, resp.StatusCode)
		}

		if strings.TrimSpace(string(body)) != tc.want {
			t.Errorf("Want '%s', got '%s'", tc.want, body)
		}

		tickerCount++
		if tickerCount >= len(tt) {
			break
		}
	}

	// for _, tc := range tt {
	// 	t.Run(tc.name, func(t *testing.T) {

	// 		req, err := http.NewRequest("GET", serverAddr+"/v1/api/test", strings.NewReader(tc.input))
	// 		if err != nil {
	// 			t.Fatal(err)
	// 			return
	// 		}
	// 		req.Header.Set("Content-Type", "application/json")
	// 		httpClient := &http.Client{}
	// 		resp, err := httpClient.Do(req)
	// 		if err != nil {
	// 			t.Fatal(err)
	// 			return
	// 		}
	// 		defer resp.Body.Close()
	// 		body, err := ioutil.ReadAll(resp.Body)
	// 		if err != nil {
	// 			t.Fatal(err)
	// 			return
	// 		}
	// 		fmt.Println(body)

	// 		// if responseRecorder.Code != tc.statusCode {
	// 		// 	t.Errorf("Want status '%d', got '%d'", tc.statusCode, responseRecorder.Code)
	// 		// }

	// 		// if strings.TrimSpace(responseRecorder.Body.String()) != tc.want {
	// 		// 	t.Errorf("Want '%s', got '%s'", tc.want, responseRecorder.Body)
	// 		// }
	// 	})
	// }
}

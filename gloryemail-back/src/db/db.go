package db

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func TestDb(writer http.ResponseWriter, request *http.Request) {
	requestURL := "http://zincsearch:4080/healthz"
	zincReq, _ := http.NewRequest(http.MethodGet, requestURL, nil)

	res, err := http.DefaultClient.Do(zincReq)
	if err != nil {
		fmt.Println("Error on HTTP Request Test DB: ")
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error on Test DB, could not read response body: ")
	}

	if _, errWrite := writer.Write([]byte(string(resBody))); errWrite != nil {
		fmt.Println("Error on Test DB, could not write response body: ", errWrite)
	}
}

func GetEmails(writer http.ResponseWriter, request *http.Request) {
	query := chi.URLParam(request, "query")
	fmt.Println(query)
	requestURL := "http://zincsearch:4080/api/email/_search"
	var jsonStr []byte
	if query != "" {
		jsonStr = []byte(fmt.Sprintf(`{
			"search_type": "match",
			"query": {
					"term": "%v",
					"field": "body"
			},
			"from": 0,
			"max_results": 20,
			"_source": []
		}`, query))
	} else {
		jsonStr = []byte(`{ "max_results": 20 }`)
	}
	zincReq, _ := http.NewRequest(http.MethodPost, requestURL, bytes.NewBuffer(jsonStr))
	zincReq.Header.Add("Content-Type", "application/json")
	zincReq.SetBasicAuth("admin", "Complexpass#123")

	res, err := http.DefaultClient.Do(zincReq)
	if err != nil {
		fmt.Println("Error on HTTP Request Test DB: ")
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error on Get Emails, could not read response body: ")
	}

	if _, errWrite := writer.Write([]byte(string(resBody))); errWrite != nil {
		fmt.Println("Error on B, could not write response body: ", errWrite)
	}
}

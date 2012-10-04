// Steve Phillips / elimisteve
// 2012.09.26

package tent

import (
	"encoding/json"
	"fmt"
	"github.com/elimisteve/fun"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	DEBUG = false
)

func PostStatus(currentAuthDetails, message string) ([]byte, error) {
	// TODO: JSON-escape `message`
	body := fmt.Sprintf(STATUS_BODY, time.Now().Unix(), message)
	req, err := genRequest(body)
	fun.MaybeFatalAt("http.NewRequest", err)
	defer req.Body.Close()

	authStr, err := genAuthStr(currentAuthDetails)
	fun.MaybeFatalAt("genAuthStr", err)

	addHeaders(req, authStr)

	if DEBUG {
		viewRequestBody(req)
	}

	// Perform request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error posting to %s: %v\n", ENV.URL.Host, err)
	}
	jsonResp, err := ioutil.ReadAll(resp.Body)
	fun.MaybeFatalAt("ioutil.ReadAll", err)
	return jsonResp, nil
}

func addHeaders(req *http.Request, authStr string) {
	for name, value := range STATUS_HEADERS {
		req.Header.Add(name, value)
	}
	req.Header.Add("Authorization", authStr)
}

func viewRequestBody(req *http.Request) {
	reqStr, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("Error reading from req.Body: %v\n", err)
	}
	fmt.Printf("Request body == '%s'\n\n", reqStr)
	log.Fatalf("`client.Do` won't work anyway; stopping early\n")
}

func genAuthStr(currentAuthDetails string) (string, error) {
	mac := Mac{}
	err := json.Unmarshal([]byte(currentAuthDetails), &mac)
	if err != nil {
		return "", err
	}
	authStr := signRequest(&mac)
	return authStr, nil
}

func genRequest(body string) (*http.Request, error) {
	reader := strings.NewReader(body)
	req, err := http.NewRequest(ENV.Method, ENV.URL.String(), reader)
	return req, err
}

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
	mac := Mac{}
	err := json.Unmarshal([]byte(currentAuthDetails), &mac)
	fun.MaybeFatalAt("json.Unmarshal", err)

	authStr := signRequest(&mac)
	// TODO: JSON-escape `message`
	body := fmt.Sprintf(STATUS_BODY, time.Now().Unix(), message)

	// Server communication setup
	client := &http.Client{}
	// jsonStr, err := json.Marshal(body)
	// fun.MaybeFatalAt("json.MaybeFatalAt", err)

	reader := strings.NewReader(body)
	req, err := http.NewRequest(ENV.Method, ENV.URL.String(), reader)
	fun.MaybeFatalAt("http.NewRequest", err)

	// Add headers
	for name, value := range STATUS_HEADERS {
		req.Header.Add(name, value)
	}
	req.Header.Add("Authorization", authStr)

	if DEBUG {
		reqStr, err := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Printf("Error reading from req.Body: %v\n", err)
		}
		fmt.Printf("Request body == '%s'\n\n", reqStr)
		log.Fatalf("`client.Do` won't work anyway; stopping early\n")
	}

	// Perform request
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error posting to %s: %v\n", ENV.URL.Host, err)
	}
	jsonResp, err := ioutil.ReadAll(resp.Body)
	fun.MaybeFatalAt("ioutil.ReadAll", err)
	return jsonResp, nil
}
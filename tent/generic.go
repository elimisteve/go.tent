// Steve Phillips / elimisteve
// 2012.09.26

package tent

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func genericRequest(req *http.Request, info *RequestInfo) ([]byte, error) {
	// Generate auth string used to sign request
	authStr := signRequest(req.Method, info)

	addHeaders(req, authStr)

	// fmt.Printf("\ngenericRequest: req == %+v\n\n", req)

	// Perform request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error posting to %s: %v\n", info.Host, err)
	}

	// Read and return response body
	jsonResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		str := "Error reading from resp.Body within genericRequest: %v"
		return nil, fmt.Errorf(str, err)
	}
	defer resp.Body.Close()

	return jsonResp, nil
}

func addHeaders(req *http.Request, authStr string) {
	for name, value := range STATUS_HEADERS {
		req.Header.Add(name, value)
	}
	req.Header.Add("Authorization", authStr)
}

func viewRequestBodyFatal(req *http.Request) {
	reqStr, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("Error reading from req.Body: %v\n", err)
	}
	fmt.Printf("Request body == '%s'\n\n", reqStr)
	log.Fatalf("`client.Do` won't work anyway; stopping early\n")
}

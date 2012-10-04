// Steve Phillips / elimisteve
// 2012.10.02

package tent

import (
	"fmt"
	"net/http"
	"strings"
)

func Post(info *RequestInfo, data string) ([]byte, error) {
	req, err := http.NewRequest("POST", info.FullURL(), strings.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("Error creating new request from Post: %v", err)
	}
	defer req.Body.Close()

	// Perform request
	resp, err := genericRequest(req, info)
	if err != nil {
		str := "Error calling genericRequest from Post: %v"
		return nil, fmt.Errorf(str, err)
	}
	defer resp.Body.Close()

	return resp, nil
}

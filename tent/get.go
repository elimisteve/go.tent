// Steve Phillips / elimisteve
// 2012.10.03

package tent

import (
	"fmt"
	"net/http"
)

func Get(info *RequestInfo) ([]byte, error) {
	// Build request
	req, err := http.NewRequest("GET", info.FullURL(), nil)
	if err != nil {
		str := "Error calling http.NewRequest from Get: %v"
		return nil, fmt.Errorf(str, err)
	}

	// Perform request
	jsonResp, err := genericRequest(req, info)
	if err != nil {
		return nil, fmt.Errorf("Error calling reqToResp from Get: %v", err)
	}

	return jsonResp, nil
}

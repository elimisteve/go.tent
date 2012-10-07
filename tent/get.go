// Steve Phillips / elimisteve
// 2012.10.03

package tent

import (
	"fmt"
	"net/http"
)

func get(info *RequestInfo) ([]byte, error) {
	// Build request
	req, err := http.NewRequest("GET", info.FullURL(), nil)
	if err != nil {
		return nil, fmt.Errorf("Error calling http.NewRequest from Get: %v",
			err)
	}

	// Perform request
	jsonResp, err := genericRequest(req, info)
	if err != nil {
		return nil, fmt.Errorf("Error calling reqToResp from Get: %v", err)
	}

	return jsonResp, nil
}

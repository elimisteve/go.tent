// Steve Phillips / elimisteve
// 2012.10.02

package tent

import (
	"fmt"
	"net/http"
	"strings"
)

func Post(info *RequestInfo, body string) ([]byte, error) {
	req, err := http.NewRequest("POST", info.FullURL(), strings.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("Error creating new request from Post: %v", err)
	}
	defer req.Body.Close()

	// Perform request
	jsonResp, err := genericRequest(req, info)
	if err != nil {
		str := "Error calling genericRequest from Post: %v"
		return nil, fmt.Errorf(str, err)
	}

	return jsonResp, nil
}

// NewSimpleStatus generates a new StatusPost using the given message
// intended for the given recipients
func NewSimpleStatus(message string, recipients ...Entity) (status StatusPost, err error) {
	licenses := []URL{LICENSE_CREATIVE_COMMONS}
	permissions := Permissions{Entities: recipients}

    status = StatusPost {
	Type: TYPE_STATUS,
	PublishedAt: time.Now().Unix(),
	Permissions: permissions,
	Licenses: licenses,
	}
}
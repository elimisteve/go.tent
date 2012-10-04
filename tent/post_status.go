// Steve Phillips / elimisteve
// 2012.09.26

package tent

import (
	"fmt"
	"time"
)

const (
	DEBUG = false
)

func PostStatus(info *RequestInfo, message string) ([]byte, error) {
	// TODO: JSON-escape `message` or use then JSONify a ~Status struct
	body := fmt.Sprintf(STATUS_BODY, time.Now().Unix(), message)
	return Post(info, body)
}

// Steve Phillips / elimisteve
// 2012.10.03

package tent

import (
	"fmt"
	"github.com/elimisteve/fun"
	"io/ioutil"
	"net/http"
)

func GetStatuses(currentAuthDetails string) ([]byte, error) {
	// Build request
	req, err := http.NewRequest(ENV.Method, ENV.URL.String(), nil)
	fun.MaybeFatalAt("http.Get", err)

	authStr, err := genAuthStr(currentAuthDetails)
	fun.MaybeFatalAt("genAuthStr", err)

	addHeaders(req, authStr)

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

// Steve Phillips / elimisteve
// 2012.09.26

package tent

import (
	"strings"
)

type RequestInfo struct {
	Proto       string
	Host        string
	Port        string
	Path        string
	Mac         *Mac
}

func newRequestInfo(host, path string, mac *Mac) *RequestInfo {
	req := RequestInfo{
		Proto:       "https",
		Host:        host,
		Port:        "443",
		Path:        path,
		Mac:         mac,
	}
	// If protocol included in Host string, strip it off
	if strings.HasPrefix(req.Host, "http://") {
		// Eveverything after the "http://"
		req.Host = req.Host[len("http://"):]
		req.Proto = "http"
	}
	if strings.HasPrefix(req.Host, "https://") {
		// Eveverything after the "https://"
		req.Host = req.Host[len("https://"):]
		req.Proto = "https"
	}
	return &req
}

func (req *RequestInfo) FullURL() string {
	maybeSlash := ""
	// If no '/' between domain and path, add one
	if !strings.HasSuffix(req.Host, "/") && !strings.HasPrefix(req.Path, "/") {
		maybeSlash = "/"
	}
	return strings.ToLower(req.Proto) + "://" + req.Host + ":" + req.Port +
		maybeSlash + req.Path
}

// Steve Phillips / elimisteve
// 2012.09.26

package tent

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/elimisteve/fun"
	"hash"
	"strconv"
	"strings"
	"time"
)

const HEX_CHARSET = `0123456789abcdef`

type Mac struct {
	KeyId     string `json:"mac_key_id"`
	Key       string `json:"mac_key"`
	Algorithm string `json:"mac_algorithm"`
}

func signRequest(method string, info *RequestInfo) (authHeader string) {
	mac := info.Mac
	now := strconv.Itoa(int(time.Now().Unix()))
	nonce := fun.RandStrOfLen(NONCE_LENGTH, HEX_CHARSET)
	reqStr := buildRequestString(info, method, now, nonce)
	signed, err := macSigner(mac.Algorithm, mac.Key, reqStr)
	fun.MaybeFatalAt("macSigner", err)
	// TODO: Make sure there should only be 1 replacement
	signature := strings.Replace(base64Encode(signed), "\n", "", 1)
	authHeader = buildAuthHeader(mac, now, nonce, signature)
	return
}

func buildRequestString(info *RequestInfo, method, now, nonce string) string {
	multiline := []string{now, nonce, strings.ToUpper(method),
		info.Path, info.Host, info.Port, "", ""}
	return strings.Join(multiline, "\n")
}

func buildAuthHeader(mac *Mac, now, nonce, signature string) string {
	formatStr := `MAC id="%s", ts="%s", nonce="%s", mac="%s"`
	return fmt.Sprintf(formatStr, mac.KeyId, now, nonce, signature)
}

func macSigner(algo, key, request string) (signed string, err error) {
	var signFunc func() hash.Hash
	switch algo {
	case "hmac-sha-1":
		signFunc = func() hash.Hash {
			return sha1.New()
		}
	case "hmac-sha-256":
		signFunc = func() hash.Hash {
			return sha256.New()
		}
	default:
		return "", fmt.Errorf("Signing algorithm '%s' not supported", algo)
	}
	h := hmac.New(signFunc, []byte(key))
	h.Write([]byte(request))
	signed = string(h.Sum(nil))
	return
}

func base64Encode(s string) (encoded string) {
	enc := make([]byte, base64.StdEncoding.EncodedLen(len(s)))
	base64.StdEncoding.Encode(enc, []byte(s))
	encoded = string(enc)
	return
}

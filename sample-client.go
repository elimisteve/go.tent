// Steve Phillips / elimisteve
// 2012.09.26

package main

import (
	"github.com/elimisteve/go.tent/tent"
	"github.com/elimisteve/fun"
)

func main() {
	// See https://github.com/tent/tent-client-ruby

	mac := tent.Mac{
        KeyId:     "be94a6bf",
        Key:       "974af035",
        Algorithm: "hmac-sha-256",
	}

	// Server communication
	client := tent.NewClient("https://tent.tent.is", &mac)

	// Tent profile discovery
	err := client.Discover("http://tent-user.example.org")
	fun.MaybeFatalAt("client.Discover", err)

	err = client.Following().Create("http://another-tent.example.com")
	fun.MaybeFatalAt("client.Following().Create", err)
}
// Steve Phillips / elimisteve
// 2012.09.26

package main

import (
	"fmt"
	"github.com/elimisteve/fun"
	"github.com/elimisteve/go.tent/tent"
	"log"
	"os"
)

const (
	USERNAME             = ""
	TENT_SERVER          = USERNAME + ".tent.is"
	CURRENT_AUTH_DETAILS = ``
)

func main() {
	if USERNAME == "" || CURRENT_AUTH_DETAILS == "" {
		log.Fatalf("Set USERNAME and CURRENT_AUTH_DETAILS in this file")
	}

	// TODO: Get rid of ENV. Don't copy the Ruby design; thread safety
	// is good.
	tent.ENV = tent.NewEnv(TENT_SERVER, "443", "/tent/posts", "post")

	if len(os.Args) < 2 {
		log.Fatalf("Say something!\n")
	}

	// Post new status update to TENT_SERVER
	responseBody, err := tent.PostStatus(CURRENT_AUTH_DETAILS, os.Args[1])
	fun.MaybeFatalAt("tent.PostStatus", err)
	fmt.Printf("%s\n", responseBody)

	//
	// TODO
	//

	// app := tent.NewApp(client)

	// // Tent profile discovery
	// err := client.Discover("http://tent-user.example.org")
	// fun.MaybeFatalAt("client.Discover", err)

	// err = client.Following().Create("http://another-tent.example.com")
	// fun.MaybeFatalAt("client.Following().Create", err)
}

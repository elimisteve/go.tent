// Steve Phillips / elimisteve
// 2012.09.26

package main

import (
	"fmt"
	"github.com/elimisteve/fun"
	"github.com/elimisteve/go.tent/tent"
	"log"
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
	//
	// Post new status update to TENT_SERVER
	//
	status := "This message posted with go.tent's sample client "
	status += "https://github.com/elimisteve/go.tent"
	responseBody, err := tent.PostStatus(CURRENT_AUTH_DETAILS, status)
	fun.MaybeFatalAt("tent.PostStatus", err)
	fmt.Printf("Post from %s: \n%s\n", TENT_SERVER, responseBody)


	tent.ENV = tent.NewEnv(TENT_SERVER, "443", "/tent/followings", "get")
	//
	// Get followers of TENT_SERVER
	//
	responseBody, err = tent.GetStatuses(CURRENT_AUTH_DETAILS)
	fun.MaybeFatalAt("tent.GetStatuses", err)
	fmt.Printf("\n\nWho %s is following: \n%s\n", TENT_SERVER, responseBody)


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

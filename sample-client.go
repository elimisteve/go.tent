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

	// Construct a client
	client, err := tent.NewClientWithAuthStr(TENT_SERVER, CURRENT_AUTH_DETAILS)

	//
	// Post new status update to TENT_SERVER
	//
	fun.MaybeFatalAt("tent.PostStatus", err)
	status := "This message posted with go.tent's sample client "
	status += "https://github.com/elimisteve/go.tent"
	responsePost, err := client.PostStatus(status)
	fun.MaybeFatalAt("tent.PostStatus", err)
	fmt.Printf("Post from %s: \n%+v\n\n", TENT_SERVER, responsePost)

	//
	// Get entities that the user follows
	//
	followings, err := client.GetFollowings()
	fun.MaybeFatalAt("client.GetFollowings", err)
	fmt.Printf("\n\n%d users %s is following:\n\n", len(followings), TENT_SERVER)
	// Loop over tent.Following vars
	for _, f := range followings[:] {
		fmt.Printf("%#v\n\n", f)
	}

	//
	// TODO
	//

	// // Tent profile discovery
	// err := client.Discover("http://tent-user.example.org")
	// fun.MaybeFatalAt("client.Discover", err)

	// err = client.Following().Create("http://another-tent.example.com")
	// fun.MaybeFatalAt("client.Following().Create", err)
}

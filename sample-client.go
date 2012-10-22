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
	fun.MaybeFatalAt("tent.NewClientWithAuthStr", err)

	//
	// Post new status update to TENT_SERVER
	//
	msg := "This message posted with go.tent's sample client "
	msg += "https://github.com/elimisteve/go.tent"
	statusPost, err := client.PostStatus(msg)
	fun.MaybeFatalAt("tent.PostStatus", err)
	fmt.Printf("You just posted to Tent.is! See it at https://%s/posts/%s",
		TENT_SERVER, statusPost.Id)

	//
	// Get entities that the user follows
	//
	followings, err := client.GetFollowings()
	fun.MaybeFatalAt("client.GetFollowings", err)
	fmt.Printf("\n\n%d users %s is following:\n\n", len(followings), TENT_SERVER)

	// Loop over tent.Following vars
	for ndx, f := range followings {
		if ndx != 0 {
			fmt.Printf(", ")
		}
		fmt.Printf("%s", f.Entity)
	}
	fmt.Printf("\n")
}

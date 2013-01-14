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
	// Get your most recent statuses
	//
	posts, err := client.GetStatuses()
	fun.MaybeFatalAt("client.GetStatuses", err)
	if len(posts) != 0 {
		post := posts[0]
		fmt.Printf("The latest post in %s's feed, from %s:\n%v",
			TENT_SERVER, post.Entity, post.Content.Text)
	}

	//
	// Get entities that the user follows
	//
	followings, err := client.GetFollowings()
	fun.MaybeFatalAt("client.GetFollowings", err)
	fmt.Printf("\n\n%s's %d most-recently-followed users:\n\n",
		TENT_SERVER, len(followings))

	// Loop over tent.Following vars
	for ndx, f := range followings {
		if ndx != 0 {
			fmt.Printf(", ")
		}
		fmt.Printf("%s", f.Entity)
	}
	fmt.Printf("\n\n")

	//
	// Post new private status that only you and ^https://tent.tent.is can see
	//
	msg := "Dear ^https://tent.tent.is -- don't tell anyone this, but..."
	statusPost, err := client.PostPrivateStatus(msg, "https://tent.tent.is",
		"https://"+TENT_SERVER)
	fun.MaybeFatalAt("client.PostPrivateStatus", err)
	fmt.Printf("See a status post that only you and ^https://tent.tent.is ")
	fmt.Printf("can see at https://%s/posts/%s\n\n", TENT_SERVER, statusPost.Id)

	//
	// Post new status update to TENT_SERVER
	//
	msg = "This message posted with go.tent's sample client "
	msg += "https://github.com/elimisteve/go.tent"
	publicPost, err := client.PostStatus(msg)
	fun.MaybeFatalAt("client.PostStatus", err)
	fmt.Printf("You just posted to Tent.is! See it at https://%s/posts/%s\n",
		TENT_SERVER, publicPost.Id)
}

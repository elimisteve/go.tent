// Steve Phillips / elimisteve
// 2012.09.26

package main

import (
	"encoding/json"
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

	//
	// Post new status update to TENT_SERVER
	//
	req := tent.NewRequestInfo(TENT_SERVER, "/tent/posts",
		CURRENT_AUTH_DETAILS)
	status := "This message posted with go.tent's sample client "
	status += "https://github.com/elimisteve/go.tent"
	responseBody, err := tent.PostStatus(req, status)
	fun.MaybeFatalAt("tent.PostStatus", err)
	fmt.Printf("Post from %s: \n%s\n\n", TENT_SERVER, responseBody)

	// Unmarshal responseBody into a StatusPost
	post := tent.StatusPost{}
	err = json.Unmarshal(responseBody, &post)
	fun.MaybeFatalAt("json.Unmarshal", err)
	fmt.Printf("Unmarshal'd response: \n%+v\n", post)

	//
	// Get users followed by TENT_SERVER
	//
	req = tent.NewRequestInfo(TENT_SERVER, "/tent/followings",
		CURRENT_AUTH_DETAILS)
	responseBody, err = tent.Get(req)
	fun.MaybeFatalAt("tent.GetStatuses", err)
	// fmt.Printf("%s\n\n\n", responseBody)

	followings := []tent.Following{}
	err = json.Unmarshal(responseBody, &followings)
	fun.MaybeFatalAt("json.Unmarshal", err)

	fmt.Printf("\n\n3 users %s is following:\n\n", TENT_SERVER)
	// Loop over tent.Following vars
	for _, f := range followings[:3] {
		fmt.Printf("%#v\n\n", f)
	}


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

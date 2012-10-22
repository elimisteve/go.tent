# [Go](http://golang.org) client for [Tent.io](http://tent.io/)

## What is Tent.io?

See <http://tent.io/> for details about the latest addition to
decentralized social networking!


## Install

    go get "github.com/elimisteve/go.tent/tent"


## Simple Config

Open `sample-client.go` in your favorite text editor, then set `TENT_SERVER` and
`CURRENT_AUTH_DETAILS`.

The correct value for `CURRENT_AUTH_DETAILS` is some JSON embedded in
the page source of your Tent.is page -- e.g.,
`https://elimisteve.tent.is` for me since my username `elimisteve`.
To view source, right-click anywhere on the page then select "View
Page Source" or similar and look for the line that begins with
`CURRENT_AUTH_DETAILS`.  Copy and paste the string within that line's
call to `JSON.parse`.

The resulting values should look something like this:

    TENT_SERVER = "elimisteve.tent.is"
    CURRENT_AUTH_DETAILS = `{"mac_key_id":"u:...","mac_key":"...","mac_algorithm":"hmac-sha-256"}`


## Usage

    go run simple-client.go

See the calls to `client.GetFollowings()`, `client.GetStatuses()`, and
`client.PostStatus()`, following `client, err :=
tent.NewClientWithAuthStr(...)`, in `sample-client.go`.  ...actually, I'll just make it easy:

```
	const (
		USERNAME             = "your_username_here"
		TENT_SERVER          = USERNAME + ".tent.is"
		CURRENT_AUTH_DETAILS = `{"mac_key_id":"u:...,"mac_key":"...,"mac_algorithm":"hmac-sha-256"}`
	)

	...

	// Construct a client
	client, err := tent.NewClientWithAuthStr(TENT_SERVER, CURRENT_AUTH_DETAILS)
	...


	//
	// Get user's recent statuses
	//
	posts, err := client.GetStatuses()
	...
	for _, post := range posts {
		fmt.Printf("%s: %s\n", post.Entity, post.Content.Text)
	}


	//
	// Get entities that the user follows
	//
	followings, err := client.GetFollowings()
	...
	for _, f := range followings {
		fmt.Printf("%s\n", f.Entity)
	}


	//
	// Post new status update to TENT_SERVER
	//
	statusPost, err := client.PostStatus("Hello, Tent!")
	...
	fmt.Printf("You just posted to Tent.is! See it at https://%s/posts/%s\n",
		TENT_SERVER, statusPost.Id)
}
```

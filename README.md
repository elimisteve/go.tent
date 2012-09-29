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

    go run simple-client.go "Hello to Tent.is from go.tent"

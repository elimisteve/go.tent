// Steve Phillips / elimisteve
// 2012.09.26

package tent

import (
	"encoding/json"
	"fmt"
	// "log"
	"time"
)

type Client struct {
	URI        string
	Mac        *Mac
	FollowList *FollowList
}

// Temporary constructor for AuthDetail string
func NewClientWithAuthStr(uri, authStr string) (*Client, error) {
	mac := Mac{}
	err := json.Unmarshal([]byte(authStr), &mac)
	if err != nil {
		return nil, err
	}
	return NewClient(uri, &mac), nil
}

func NewClient(uri string, mac *Mac) *Client {
	client := Client{
		URI: uri,
		Mac: mac,
	}
	return &client
}

func (c *Client) Discover(url string) error {
	// TODO: Fill in stub
	return nil
}

func (c *Client) Following() *FollowList {
	if c.FollowList == nil {
		c.FollowList = &FollowList{}
	}
	return c.FollowList
}

func (c *Client) PostStatus(message string) (status StatusPost, err error) {
	info := newRequestInfo(c.URI, "/tent/posts", c.Mac)
	// TODO: JSON-escape `message` or use then JSONify a ~Status struct
	body := fmt.Sprintf(STATUS_BODY, time.Now().Unix(), message)
	responseBody := []byte{}
	responseBody, err = Post(info, body)
	if err != nil {
		return
	}
	err = json.Unmarshal(responseBody, &status)
	return
}

func (c *Client) PostPrivateStatus(message string, recipients ...string) (*StatusPost, error) {
	info := newRequestInfo(c.URI, "/tent/posts", c.Mac)
	// Create new StatusPost and turn it into JSON to be POSTed
	jsonData, err := json.Marshal(NewStatus(message, recipients...))
	if err != nil {
		return nil, fmt.Errorf("Error turning status into JSON: %v", err)
	}
	// fmt.Printf("\nAbout to be POSTed:\n%s\n\n", jsonData)
	responseBody, err := Post(info, string(jsonData))
	if err != nil {
		return nil, fmt.Errorf("Error posting private status: %v", err)
	}
	// fmt.Printf("\nLiteral response: %s\n\n", responseBody)
	// Unmarshal response into a second StatusPost

	// status := StatusPost{
	// 	Permissions: Permissions{
	// 		Entities: make(map[Entity]bool, len(recipients)),
	// 	},
	// }

	status := StatusPost{}
	err = json.Unmarshal(responseBody, &status)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshaling server response: %v", err)
	}
	return &status, nil
}

func (c *Client) GetStatuses() ([]StatusPost, error) {
	info := newRequestInfo(c.URI, "/tent/posts", c.Mac)
	// fmt.Printf("mac == %+v\n", *c.Mac)
	// fmt.Printf("info == %+v\n", info)
	statuses := []StatusPost{}
	responseBody, err := Get(info)
	if err != nil {
		return statuses, fmt.Errorf("Error calling Get(): %v", err)
	}
	if err := json.Unmarshal(responseBody, &statuses); err != nil {
		errStr := "Error trying to Unmarshal '%s' into a []StatusPost: %v"
		return statuses, fmt.Errorf(errStr, responseBody, err)
	}
	return statuses, err
}

func (c *Client) GetFollowings() (followings []Following, err error) {
	req := newRequestInfo(c.URI, "/tent/followings", c.Mac)
	responseBody := []byte{}
	responseBody, err = Get(req)
	if err != nil {
		return
	}
	err = json.Unmarshal(responseBody, &followings)
	return
}

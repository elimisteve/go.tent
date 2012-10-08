// Steve Phillips / elimisteve
// 2012.09.26

package tent

import (
	"encoding/json"
	"fmt"
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

func (c *Client) PostStatus(message string) (rspPost StatusPost, err error) {
	info := newRequestInfo(c.URI, "/tent/posts", c.Mac)
	// TODO: JSON-escape `message` or use then JSONify a ~Status struct
	body := fmt.Sprintf(STATUS_BODY, time.Now().Unix(), message)
	responseBody := []byte{}
	responseBody, err = Post(info, body)
	if err != nil {
		return
	}
	err = json.Unmarshal(responseBody, &rspPost)
	return
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

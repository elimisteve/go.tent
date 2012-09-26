// Steve Phillips / elimisteve
// 2012.09.26

package tent

type Mac struct {
	KeyId     string
	Key       string
	Algorithm string
}


type FollowList []string

func (fl *FollowList) Create(url string) error {
	*fl = append(*fl, url)
	// TODO: Do remote call to grab info(???) from URL
	return nil
}


type Client struct {
	URL        string
	Mac        *Mac
	FollowList *FollowList
}

func NewClient(url string, mac *Mac) *Client {
	client := Client {
		URL: url,
		Mac: mac,
	}
	return &client
}

func (c *Client) Discover(url string) error {
	// TODO: Fill in stub
	return nil
}

func (c *Client) Following() *FollowList {
	c.FollowList = &FollowList{}
	return c.FollowList
}

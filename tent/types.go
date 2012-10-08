// Steve Phillips / elimisteve
// 2012.09.26

package tent

type FollowList []string

func (fl *FollowList) Create(url string) error {
	*fl = append(*fl, url)
	// TODO: Do remote call to grab info(???) from URL
	return nil
}

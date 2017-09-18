package pfs

import "fmt"

// FullID prints repoName/CommitID
func (c *Commit) FullID() string {
	return fmt.Sprintf("%s/%s", c.Repo.Name, c.ID)
}

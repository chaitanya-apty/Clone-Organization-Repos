package clone

import (
	gh "clone-repos/clone/github"
	typ "clone-repos/clone/types"
)

//newService -
func newService() typ.DownloadRepos {
	return &gh.Service{}
}

package endly

import (
	"github.com/viant/toolbox/url"
)

//DockerPushRequest represents a docker push request
type DockerPushRequest struct {
	SysPath []string
	Target  *url.Resource
	Tag     *DockerTag
}

//DockerPushResponse represents a docker push request
type DockerPushResponse struct {
}

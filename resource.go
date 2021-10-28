package bitbucket

import (
	"fmt"
	"net/url"

	"github.com/levigross/grequests"
)

type Resource struct {
	Base_url    *url.URL
	Base_path   string
	Api_path    string
	Session     *grequests.Session
	ChunkSize   int
	Url_builder URLBuilder
	Parameters  ParamsBuilder
}

func NewResource(base_url *url.URL, base_path, api_path string, session *grequests.Session, chunkSize int) (R Resource) {

	R.Base_url = base_url
	R.Base_path = base_path
	R.Api_path = api_path
	R.Session = session
	R.ChunkSize = chunkSize
	R.Url_builder = URLBuilderNew(base_url, base_path, api_path)
	R.Parameters = NewParamsBuilder()

	return R
}

func (R Resource) String() string {
	return fmt.Sprintf("<[%s]>", R.path())
}

func (R Resource) path() string {
	return fmt.Sprintf("%s", R.Base_path+R.Api_path)
}

func (R Resource) _request() BitbucketRequest {

	return BitbucketRequestNew(R.Parameters, R.Session, R.Url_builder, 0, R)
}

func (R Resource) Get() {

}

func (R Resource) Update() {

}

func (R Resource) Create() {

}

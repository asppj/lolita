package article

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type reproduceParams struct {
	URL   string `form:"url"` // 转载url
	Biz   string `form:"biz"`
	Mid   string `form:"mid"`
	Idx   string `form:"idx"`
	Sn    string `form:"sn"`
	Chksm string `form:"chksm"`
}

// Reproduced 转载
func Reproduced(ctx *gin.Context) {
	param := &reproduceParams{}
	if err := ctx.BindQuery(param); err != nil {
		return
	}

	return
}

// ClientHTTP client
type ClientHTTP struct {
	client  *http.Client
	headers map[string]string
}

func newClientHTTP() *ClientHTTP {
	return &ClientHTTP{
		headers: make(map[string]string),
		client: &http.Client{
			Transport:     nil,
			CheckRedirect: nil,
			Jar:           nil,
			Timeout:       0,
		}}
}

// AddHeaders 添加header
func (c *ClientHTTP) AddHeaders(k, v string) {
	c.headers[k] = v
}

// Get get
func (c *ClientHTTP) Get(url string) (resp *http.Response, err error) {
	hresp, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	for i, v := range c.headers {
		hresp.Header.Add(i, v)
	}
	resp, err = c.client.Do(hresp)
	return
}

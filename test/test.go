package test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

// TestSinaURLShort 测试
func TestSinaURLShort(t *testing.T) {
	originurl := "https://github.com"
	dao, _ := New()
	response, _ := dao.SinaURLShort(originurl)
	if response.URLS[0].ShortURL == "http://t.cn/RxnlTYR" {
		t.Log("SinaURLShort passed")
	} else {
		t.Error("SinaURLShort failed", originurl, " to ", response.URLS[0].ShortURL)
	}
}

// ShortDao s
type ShortDao struct{}

func New() (*ShortDao, error) {
	return &ShortDao{}, nil
}

func (dao *ShortDao) SinaURLShort(origin string) (*Response, error) {
	// TODO origin可能为非法字符串，需要考虑下校验
	resp, err := http.Get("https://api.weibo.com/2/short_url/shorten.json?source=2257828842&url_long=" + origin)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil
	}
	var response Response
	json.Unmarshal(body, &response)
	return &response, nil
}

type Request struct {
	OriginURL string
}

/*
{
    "urls": [
        {
            "result": true,
            "url_short": "http://t.cn/RxnlTYR",
            "url_long": "https://github.com",
            "object_type": "",
            "type": 0,
            "object_id": ""
        }
    ]
}
*/
type ResponseEntry struct {
	Success   bool   `json:"result"`
	ShortURL  string `json:"url_short"`
	OriginURL string `json:"url_long"`
}

type Response struct {
	URLS []ResponseEntry
}

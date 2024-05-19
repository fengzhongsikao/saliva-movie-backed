package hot

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"io"
	"log"
	"net/http"

	"saliva-movie/api/hot/v1"
)

func (c *ControllerV1) GetHotList(ctx context.Context, req *v1.GetHotListReq) (res *v1.GetHotListRes, err error) {

	hotList := getHotList()
	g.RequestFromCtx(ctx).Response.WriteJson(hotList)
	return
	//nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func getHotList() []map[string]string {

	url := "https://movie.douban.com/j/search_subjects?&type=movie&sort=recommend&tag=%E7%83%AD%E9%97%A8&page_limit=21&page_start=1"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36 Edg/124.0.0.0")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Host", "movie.douban.com")
	req.Header.Add("Connection", "keep-alive")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	jsonStr := string(body)

	err2 := json.Unmarshal([]byte(jsonStr), &data)
	if err2 != nil {
		log.Fatal("JSON解析错误:", err)
	}

	hotList := make([]map[string]string, len(data.Subjects))

	for _, v := range data.Subjects {

		m := map[string]string{
			"cover": "",
			"title": "",
			"rate":  "",
		}
		m["cover"] = v.Cover
		m["title"] = v.Title
		m["rate"] = v.Rate
		hotList = append(hotList, m)

	}

	var nonEmptyMapSlice []map[string]string
	for _, m := range hotList {
		if m != nil && len(m) > 0 {
			// 如果map不是nil且长度大于0，则添加到新切片中
			nonEmptyMapSlice = append(nonEmptyMapSlice, m)
		}
	}

	return nonEmptyMapSlice
}

var data struct {
	Subjects []struct {
		EpisodesInfo string `json:"episodes_info"`
		Rate         string `json:"rate"`
		CoverX       int    `json:"cover_x"`
		Title        string `json:"title"`
		Url          string `json:"url"`
		Playable     bool   `json:"playable"`
		Cover        string `json:"cover"`
		Id           string `json:"id"`
		CoverY       int    `json:"cover_y"`
		IsNew        bool   `json:"is_new"`
	} `json:"subjects"`
}

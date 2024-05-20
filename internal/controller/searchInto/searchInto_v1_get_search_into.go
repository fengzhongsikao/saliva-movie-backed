package searchInto

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"saliva-movie/api/searchInto/v1"
)

func (c *ControllerV1) GetSearchInto(ctx context.Context, req *v1.GetSearchIntoReq) (res *v1.GetSearchIntoRes, err error) {
	searchInto := GetSearchInfo()
	g.RequestFromCtx(ctx).Response.WriteJson(searchInto)
	return
}

func GetSearchInfo() map[string]interface{} {

	file, err := os.Open("./resource/search.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	var data data
	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil
	}

	text := "老狐狸"

	currentSearchUrl := data[0].Url + "?ac=detail&wd=" + url.QueryEscape(text)

	println(currentSearchUrl)

	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, currentSearchUrl, nil)

	if err != nil {
		fmt.Println(err)
		return nil
	}

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

	var search Search

	err2 := json.Unmarshal([]byte(jsonStr), &search)
	if err2 != nil {
		log.Fatal("JSON解析错误:", err)
	}

	vodName := ""
	vodPic := ""
	vodPlayUrl := ""
	vodContent := ""
	//vodClass := ""
	subTitle := ""
	vodId := 0

	for _, v := range search.List {

		vodName = v.VodName
		vodPic = v.VodPic

		lastDollarIndex := strings.LastIndex(v.VodPlayURL, "$")
		substring := v.VodPlayURL[lastDollarIndex+1:]

		vodPlayUrl = substring
		vodContent = v.VodContent
		//vodClass = v.VodClass
		subTitle = v.VodYear + "/" + v.VodArea + "/" + v.VodClass + "/" + v.VodTime
		vodId = v.VodID
	}

	m := map[string]interface{}{
		"title":        vodName,
		"subTitle":     subTitle,
		"pic":          vodPic,
		"brief":        vodContent,
		"playUrl":      vodPlayUrl,
		"movieId":      vodId,
		"resourceLink": data[0].Url,
	}
	//vodClass := search.MyClass

	return m

}

type data []struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Search struct {
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	Page      int    `json:"page"`
	PageCount int    `json:"pagecount"`
	Limit     string `json:"limit"`
	Total     int    `json:"total"`
	List      []struct {
		VodID            int    `json:"vod_id"`
		TypeID           int    `json:"type_id"`
		TypeID1          int    `json:"type_id_1"`
		GroupID          int    `json:"group_id"`
		VodName          string `json:"vod_name"`
		VodSub           string `json:"vod_sub"`
		VodEn            string `json:"vod_en"`
		VodStatus        int    `json:"vod_status"`
		VodLetter        string `json:"vod_letter"`
		VodColor         string `json:"vod_color"`
		VodTag           string `json:"vod_tag"`
		VodClass         string `json:"vod_class"`
		VodPic           string `json:"vod_pic"`
		VodPicThumb      string `json:"vod_pic_thumb"`
		VodPicSlide      string `json:"vod_pic_slide"`
		VodPicScreenshot string `json:"vod_pic_screenshot"`
		VodActor         string `json:"vod_actor"`
		VodDirector      string `json:"vod_director"`
		VodWriter        string `json:"vod_writer"`
		VodBehind        string `json:"vod_behind"`
		VodBlurb         string `json:"vod_blurb"`
		VodRemarks       string `json:"vod_remarks"`
		VodPubDate       string `json:"vod_pubdate"`
		VodTotal         int    `json:"vod_total"`
		VodSerial        string `json:"vod_serial"`
		VodTv            string `json:"vod_tv"`
		VodWeekday       string `json:"vod_weekday"`
		VodArea          string `json:"vod_area"`
		VodLang          string `json:"vod_lang"`
		VodYear          string `json:"vod_year"`
		VodVersion       string `json:"vod_version"`
		VodState         string `json:"vod_state"`
		VodAuthor        string `json:"vod_author"`
		VodJumpUrl       string `json:"vod_jumpurl"`
		VodTpl           string `json:"vod_tpl"`
		VodTplPlay       string `json:"vod_tpl_play"`
		VodTplDown       string `json:"vod_tpl_down"`
		VodISend         int    `json:"vod_isend"`
		VodLock          int    `json:"vod_lock"`
		VodLevel         int    `json:"vod_level"`
		VodCopyright     int    `json:"vod_copyright"`
		VodPoints        int    `json:"vod_points"`
		VodPointsPlay    int    `json:"vod_points_play"`
		VodPointsDown    int    `json:"vod_points_down"`
		VodHits          int    `json:"vod_hits"`
		VodHitsDay       int    `json:"vod_hits_day"`
		VodHitsWeek      int    `json:"vod_hits_week"`
		VodHitsMonth     int    `json:"vod_hits_month"`
		VodDuration      string `json:"vod_duration"`
		VodUp            int    `json:"vod_up"`
		VodDown          int    `json:"vod_down"`
		VodScore         string `json:"vod_score"`
		VodScoreAll      int    `json:"vod_score_all"`
		VodScoreNum      int    `json:"vod_score_num"`
		VodTime          string `json:"vod_time"`
		VodTimeAdd       int    `json:"vod_time_add"`
		VodTimeHits      int    `json:"vod_time_hits"`
		VodTimeMake      int    `json:"vod_time_make"`
		VodTrySee        int    `json:"vod_trysee"`
		VodDouBanID      int    `json:"vod_douban_id"`
		VodDouBanScore   string `json:"vod_douban_score"`
		VodReUrl         string `json:"vod_reurl"`
		VodRelVod        string `json:"vod_rel_vod"`
		VodRelArt        string `json:"vod_rel_art"`
		VodPwd           string `json:"vod_pwd"`
		VodPwdURL        string `json:"vod_pwd_url"`
		VodPwdPlay       string `json:"vod_pwd_play"`
		VodPwdPlayURL    string `json:"vod_pwd_play_url"`
		VodPwdDown       string `json:"vod_pwd_down"`
		VodPwdDownURL    string `json:"vod_pwd_down_url"`
		VodContent       string `json:"vod_content"`
		VodPlayFrom      string `json:"vod_play_from"`
		VodPlayServer    string `json:"vod_play_server"`
		VodPlayNote      string `json:"vod_play_note"`
		VodPlayURL       string `json:"vod_play_url"`
		VodDownFrom      string `json:"vod_down_from"`
		VodDownServer    string `json:"vod_down_server"`
		VodDownNote      string `json:"vod_down_note"`
		VodDownURL       string `json:"vod_down_url"`
		VodPlot          int    `json:"vod_plot"`
		VodPlotName      string `json:"vod_plot_name"`
		VodPlotDetail    string `json:"vod_plot_detail"`
		TypeName         string `json:"type_name"`
	} `json:"list"`
}

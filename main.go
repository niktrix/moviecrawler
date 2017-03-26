package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type VootResponse struct {
	Assets []struct {
		ID          string `json:"id"`
		Type        int    `json:"type"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Images      []struct {
			Ratio  string `json:"ratio"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
			URL    string `json:"url"`
		} `json:"images"`
		Metas struct {
			ContentSynopsis string `json:"ContentSynopsis"`
			ContentType     string `json:"ContentType"`
			ContentFileName string `json:"ContentFileName"`
			MovieMainTitle  string `json:"MovieMainTitle"`
			SBU             string `json:"SBU"`
			IsDownable      string `json:"IsDownable"`
			ContentDuration string `json:"ContentDuration"`
			ReleaseYear     string `json:"ReleaseYear"`
		} `json:"metas"`
		Tags struct {
			Keywords        []string `json:"Keywords"`
			CharacterList   []string `json:"CharacterList"`
			ContributorList []string `json:"ContributorList"`
			Scene1          []string `json:"Scene1"`
			Scene2          []string `json:"Scene2"`
			Scene3          []string `json:"Scene3"`
			Scene4          []string `json:"Scene4"`
			Scene5          []string `json:"Scene5"`
			Scene6          []string `json:"Scene6"`
			Genre           []string `json:"Genre"`
			Language        []string `json:"Language"`
			AdCueTime1      []string `json:"AdCueTime1"`
			AdCueTime2      []string `json:"AdCueTime2"`
			AdCueTime3      []string `json:"AdCueTime3"`
			AdCueTime4      []string `json:"AdCueTime4"`
			AdCueTime5      []string `json:"AdCueTime5"`
			AdCueTime6      []string `json:"AdCueTime6"`
			AdCueTime7      []string `json:"AdCueTime7"`
			AdCueTime8      []string `json:"AdCueTime8"`
			MediaExternalID []string `json:"MediaExternalId"`
			WatermarkURL    []string `json:"WatermarkURL"`
			MovieDirector   []string `json:"MovieDirector"`
		} `json:"tags"`
		StartDate   int   `json:"start_date"`
		EndDate     int64 `json:"end_date"`
		ExtraParams struct {
			SysStartDate string      `json:"sys_start_date"`
			SysFinalDate string      `json:"sys_final_date"`
			ExternalIds  interface{} `json:"external_ids"`
			EntryID      string      `json:"entry_id"`
		} `json:"extra_params"`
		RURL interface{} `json:"rURL"`
	} `json:"assets"`
	TotalItems int `json:"total_items"`
	Status     struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"status"`
}

type MovieRequester struct {
	url       string
	pageIndex string
	request   *http.Request
}

func main() {
	mr := MovieRequester{}
	mr.url = "https://wapi.voot.com/ws/ott/searchAssets.json?platform=Web&pId=2"

	index := 0
	totalMovies := 0
	for {
		mr.pageIndex = strconv.Itoa(index)
		b, _ := mr.get(mr.pageIndex)
		count := mr.unmarshalMovies(b)
		totalMovies = totalMovies + count
		if count == 0 {
			break
		}
		index++

	}
	fmt.Println(totalMovies)

}

func (mr *MovieRequester) get(pageIndex string) ([]byte, error) {
	//url := "https://wapi.voot.com/ws/ott/searchAssets.json?platform=Web&pId=2"
	mr.request, _ = mr.requesrUrl()
	mr.request.Header.Add("content-type", "application/x-www-form-urlencoded")
	res, _ := http.DefaultClient.Do(mr.request)
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)

}

func (mr *MovieRequester) unmarshalMovies(b []byte) int {
	r := VootResponse{}
	json.Unmarshal(b, &r)
	for _, movies := range r.Assets {
		fmt.Println(movies.Name)
	}
	return len(r.Assets)
}

func (mr *MovieRequester) getPostVars() io.Reader {
	form := url.Values{}
	form.Add("filterTypes", "390")
	//form.Add("filter", "(and (and  contentType='Movie' ))")
	form.Add("pageIndex", mr.pageIndex)
	form.Add("pageSize", "10")
	fmt.Println(form.Encode())
	return strings.NewReader(form.Encode())
}

func (mr *MovieRequester) requesrUrl() (*http.Request, error) {
	return http.NewRequest("POST", mr.url, mr.getPostVars())

}

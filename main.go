package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type webconfig struct {
	name string
	url  string
}

var availableWebsites = []webconfig{
	//{"voot", "https://wapi.voot.com/ws/ott/searchAssets.json?platform=Web&pId=2"},
	//{"hotstar", "http://search.hotstar.com/AVS/besc"},
	{"erosnow", "http://erosnow.com/v2/catalog/movies"},
}

type MovieRequester struct {
	url       string
	pageIndex string
	request   *http.Request
	website   string
}

func main() {

	for _, webs := range availableWebsites {
		fmt.Println("Getting movies for ", webs.name)
		fmt.Println("Getting movies for ", webs.url)

		mr := MovieRequester{}
		mr.url = webs.url
		mr.website = webs.name
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

}

func (mr *MovieRequester) get(pageIndex string) ([]byte, error) {
	//url := "https://wapi.voot.com/ws/ott/searchAssets.json?platform=Web&pId=2"
	mr.request, _ = mr.requesrUrl()
	//if mr.website == "voot" {
	mr.request.Header.Add("content-type", "application/x-www-form-urlencoded")
	//	}
	//	dump, _ := httputil.DumpRequestOut(mr.request, true)
	//	fmt.Println(string(dump))
	res, _ := http.DefaultClient.Do(mr.request)
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)

}

func (mr *MovieRequester) unmarshalMovies(b []byte) int {

	switch mr.website {
	case "hotstar":
		r := HotstarResponse{}
		json.Unmarshal(b, &r)
		for _, movies := range r.ResultObj.Response.Docs {
			fmt.Println(movies.ContentTitle + "    " + mr.website)
		}
		return len(r.ResultObj.Response.Docs)
	case "voot":
		r := VootResponse{}
		err := json.Unmarshal(b, &r)
		if err != nil {
			log.Println("Error while unmarshalling voot", err)
		}
		for _, movies := range r.Assets {
			fmt.Println(movies.Name + "    " + mr.website)
		}
		return len(r.Assets)
	case "erosnow":
		r := ErosNowResponse{}
		err := json.Unmarshal(b, &r)
		if err != nil {
			log.Println("Error while unmarshalling erosnow", err)
		}
		for _, movies := range r.Rows {
			fmt.Println(movies.Title + "    " + mr.website)
		}
		return len(r.Rows)

	}
	return 0
}

func (mr *MovieRequester) getPostVars() io.Reader {
	form := url.Values{}
	switch mr.website {
	case "hotstar":
		form.Add("action", "SearchContents")
		form.Add("appVersion", "5.0.40")
		form.Add("channel", "PCTV")
		form.Add("maxResult", "12")
		form.Add("query", "*")
		form.Add("searchOrder", "counter_day desc")
		form.Add("startIndex", mr.pageIndex)
		form.Add("type", "type:MOVIE")
	case "voot":
		form := url.Values{}
		form.Add("filterTypes", "390")
		form.Add("filter", "(and (and  contentType='Movie' ))")
		form.Add("pageIndex", mr.pageIndex)
		form.Add("pageSize", "10")
		//	fmt.Println(form.Encode())
	}

	return strings.NewReader(form.Encode())
}

func (mr *MovieRequester) requesrUrl() (*http.Request, error) {
	//TODO temp workaround

	switch mr.website {
	case "hotstar":
		v := "?action=SearchContents&appVersion=5.0.40&channel=PCTV&maxResult=12&moreFilters=language:hindi%3B&query=*&searchOrder=counter_day+desc&startIndex=" + mr.pageIndex + "&type=MOVIE"
		return http.NewRequest("GET", mr.url+v, nil)
	case "erosnow":
		v := "?content_type_id=1&start_index=" + mr.pageIndex + "&max_result=20&cc=IN"
		return http.NewRequest("GET", mr.url+v, nil)
	case "voot":
		mr.getPostVars()
		payload := strings.NewReader("filterTypes=390&filter=(and%20(and%20%20contentType%3D'Movie'%20))&pageIndex=" + mr.pageIndex)
		return http.NewRequest("POST", mr.url, payload)
	default:
		return http.NewRequest("GET", mr.url, nil)
	}
	return nil, nil
}

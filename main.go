package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
)

type webconfig struct {
	name string
	url  string
}

var availableWebsites = []webconfig{
	{"voot", "https://wapi.voot.com/ws/ott/searchAssets.json?platform=Web&pId=2"},
	{"hotstar", "http://search.hotstar.com/AVS/besc"},
	{"erosnow", "http://erosnow.com/v2/catalog/movies"},
}

type MovieRequester struct {
	url       string
	pageIndex int
	request   *http.Request
	website   string
	db        *mgo.Collection
}

var session *mgo.Session

func main() {

	var err error

	session, err = mgo.Dial("52.168.20.79")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	fetchMovies()
}

func fetchMovies() {
	for _, webs := range availableWebsites {
		fmt.Println("Getting movies for ", webs.name)
		fmt.Println("Getting movies for ", webs.url)

		mr := MovieRequester{}
		mr.url = webs.url
		mr.db = session.DB("movies").C("list")
		mr.website = webs.name
		totalMovies := 0
		mr.pageIndex = 0
		for {
			b, _ := mr.get()
			count := mr.unmarshalMovies(b)
			totalMovies = totalMovies + count
			if count == 0 {
				break
			}
			//TODO workaround for erosnow and voot to as they take index instead of page number
			if mr.website == "voot" {
				mr.pageIndex = mr.pageIndex + 1
			} else {
				mr.pageIndex = mr.pageIndex + 20

			}

		}
		fmt.Println(totalMovies)
	}

}

func (mr *MovieRequester) get() ([]byte, error) {
	mr.request, _ = mr.requesrUrl()
	mr.request.Header.Add("content-type", "application/x-www-form-urlencoded")
	res, _ := http.DefaultClient.Do(mr.request)
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)

}

func (mr *MovieRequester) unmarshalMovies(b []byte) int {

	switch mr.website {
	case "hotstar":
		r := HotstarResponse{}
		json.Unmarshal(b, &r)
		for _, movie := range r.ResultObj.Response.Docs {
			movie.Website = mr.website
			mr.db.Insert(movie)
			fmt.Println(movie.ContentTitle + "    " + mr.website)
		}
		return len(r.ResultObj.Response.Docs)
	case "voot":
		r := VootResponse{}
		err := json.Unmarshal(b, &r)
		if err != nil {
			log.Println("Error while unmarshalling voot", err)
		}
		for _, movie := range r.Assets {
			movie.Website = mr.website
			mr.db.Insert(movie)
			fmt.Println(movie.Name + "    " + mr.website)
		}
		return len(r.Assets)
	case "erosnow":
		r := ErosNowResponse{}
		err := json.Unmarshal(b, &r)
		if err != nil {
			log.Println("Error while unmarshalling erosnow", err)
		}
		for _, movie := range r.Rows {
			movie.Website = mr.website
			mr.db.Insert(movie)
			fmt.Println(movie.Title + "    " + mr.website)
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
		form.Add("startIndex", string(mr.pageIndex))
		form.Add("type", "type:MOVIE")
	case "voot":
		form := url.Values{}
		form.Add("filterTypes", "390")
		form.Add("filter", "(and (and  contentType='Movie' ))")
		form.Add("pageIndex", string(mr.pageIndex))
		form.Add("pageSize", "10")
		//	fmt.Println(form.Encode())
	}

	return strings.NewReader(form.Encode())
}

func (mr *MovieRequester) requesrUrl() (*http.Request, error) {
	//TODO temp workaround

	switch mr.website {
	case "hotstar":
		v := "?action=SearchContents&appVersion=5.0.40&channel=PCTV&maxResult=20&moreFilters=language:hindi%3B&query=*&searchOrder=counter_day+desc&startIndex=" + fmt.Sprintf("%v", mr.pageIndex) + "&type=MOVIE"
		return http.NewRequest("GET", mr.url+v, nil)
	case "erosnow":
		v := "?content_type_id=1&start_index=" + fmt.Sprintf("%v", mr.pageIndex) + "&max_result=20&cc=IN"
		return http.NewRequest("GET", mr.url+v, nil)
	case "voot":
		mr.getPostVars()
		payload := strings.NewReader("filterTypes=390&filter=(and%20(and%20%20contentType%3D'Movie'%20))&pageIndex=" + fmt.Sprintf("%v", mr.pageIndex))
		return http.NewRequest("POST", mr.url, payload)
	default:
		return http.NewRequest("GET", mr.url, nil)
	}
	return nil, nil
}

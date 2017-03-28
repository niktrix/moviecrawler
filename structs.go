package main

import (
	"time"
)

type HotstarResponse struct {
	ErrorDescription string `json:"errorDescription"`
	Message          string `json:"message"`
	ResultCode       string `json:"resultCode"`
	ResultObj        struct {
		Response struct {
			Docs []struct {
				CommonResponse
				Actors            string        `json:"actors"`
				Anchors           string        `json:"anchors"`
				Authors           string        `json:"authors"`
				BroadcastDate     int           `json:"broadcastDate"`
				CategoryName      string        `json:"categoryName"`
				ChannelName       string        `json:"channelName"`
				ContentID         int           `json:"contentId"`
				ContentSubtitle   string        `json:"contentSubtitle"`
				ContentTitle      string        `json:"contentTitle"`
				ContentType       string        `json:"contentType"`
				ContractEnd       time.Time     `json:"contractEnd"`
				ContractStart     time.Time     `json:"contractStart"`
				Counter           string        `json:"counter"`
				CounterDay        string        `json:"counter_day"`
				CounterWeek       string        `json:"counter_week"`
				Country           string        `json:"country"`
				Directors         string        `json:"directors"`
				Duration          int           `json:"duration"`
				EpisodeNumber     int           `json:"episodeNumber"`
				EpisodeTitle      string        `json:"episodeTitle"`
				Genre             string        `json:"genre"`
				IsAdult           string        `json:"isAdult"`
				IsLastDays        string        `json:"isLastDays"`
				IsNew             string        `json:"isNew"`
				Language          string        `json:"language"`
				Lastupdatedate    int           `json:"lastupdatedate"`
				Latest            string        `json:"latest"`
				LongDescription   string        `json:"longDescription"`
				ObjectSubtype     string        `json:"objectSubtype"`
				ObjectType        string        `json:"objectType"`
				OnAir             string        `json:"onAir"`
				PackageID         string        `json:"packageId"`
				PackageList       []interface{} `json:"packageList"`
				PcExtendedRatings string        `json:"pcExtendedRatings"`
				PcLevelVod        string        `json:"pcLevelVod"`
				PopularEpisode    string        `json:"popularEpisode"`
				SearchKeywords    string        `json:"searchKeywords"`
				Season            string        `json:"season"`
				Series            string        `json:"series"`
				ShortDescription  string        `json:"shortDescription"`
				TitleBrief        string        `json:"titleBrief"`
				URLPictures       string        `json:"urlPictures"`
				Year              string        `json:"year"`
			} `json:"docs"`
			Facets   []interface{} `json:"facets"`
			NumFound int           `json:"numFound"`
			Start    int           `json:"start"`
			Type     string        `json:"type"`
		} `json:"response"`
		ResponseHeader struct {
			QTime  int `json:"QTime"`
			Status int `json:"status"`
		} `json:"responseHeader"`
	} `json:"resultObj"`
	SystemTime int `json:"systemTime"`
}

type VootResponse struct {
	Assets []struct {
		CommonResponse
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

type CommonResponse struct {
	FilmName        string
	FilmWebsite     string
	FilmDescription string
	FilmDirector    string
	FilmCast        string
}

type ErosNowResponse struct {
	Count string `json:"count"`
	Total string `json:"total"`
	Rows  []struct {
		CommonResponse
		AssetID     string   `json:"asset_id"`
		Title       string   `json:"title"`
		Language    string   `json:"language"`
		Rating      string   `json:"rating"`
		Description string   `json:"description"`
		Subtitles   []string `json:"subtitles"`
		AccessLevel string   `json:"access_level"`
		Duration    string   `json:"duration"`
		People      struct {
			Producer      []string `json:"Producer"`
			MusicDirector []string `json:"Music director"`
			Actor         []string `json:"Actor"`
			Director      []string `json:"Director"`
		} `json:"people"`
		ShortDescription string `json:"short_description"`
		Free             string `json:"free"`
		AssetType        string `json:"asset_type"`
		ReleaseYear      string `json:"release_year"`
		Images           struct {
			Num8  string `json:"8"`
			Num9  string `json:"9"`
			Num12 string `json:"12"`
			Num13 string `json:"13"`
			Num17 string `json:"17"`
			Num22 string `json:"22"`
		} `json:"images"`
		ErosRating string `json:"eros_rating,omitempty"`
	} `json:"rows"`
}

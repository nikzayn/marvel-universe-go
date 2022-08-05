package models

type Response struct {
	Code int  `json:"code"`
	Data Data `json:"data"`
}

type Data struct {
	Offset     int          `json:"offset"`
	Limit      int          `json:"limit"`
	Total      int          `json:"total"`
	Count      int          `json:"count"`
	MarvelData []MarvelData `json:"results"`
}

type MarvelData struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Thumbnail   Thumbnail `json:"thumbnail"`
	Comics      Comics    `json:"comics"`
	Series      Series    `json:"series"`
	Stories     Stories   `json:"stories"`
	Events      Events    `json:"events"`
	Urls        []Urls    `json:"urls"`
}

type Comics struct {
	Available int `json:"available"`
}

type Series struct {
	Available int `json:"available"`
}

type Stories struct {
	Available int `json:"available"`
}

type Events struct {
	Available int `json:"available"`
}

type Urls struct {
	Type    string `json:"type"`
	UrlData string `json:"url"`
}

type Thumbnail struct {
	Path string `json:"path"`
}

package models

// Response Data
type Response struct {
	Code int  `json:"code"`
	Data Data `json:"data"`
}

// Data
type Data struct {
	Offset     int          `json:"offset"`
	Limit      int          `json:"limit"`
	Total      int          `json:"total"`
	Count      int          `json:"count"`
	MarvelData []MarvelData `json:"results"`
}

// Marvel data response
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

// Available comics
type Comics struct {
	Available int `json:"available"`
}

// Available series
type Series struct {
	Available int `json:"available"`
}

// Available stories
type Stories struct {
	Available int `json:"available"`
}

// Available events
type Events struct {
	Available int `json:"available"`
}

// URL data and type response
type Urls struct {
	Type    string `json:"type"`
	UrlData string `json:"url"`
}

// Thumbnail data
type Thumbnail struct {
	Path string `json:"path"`
}

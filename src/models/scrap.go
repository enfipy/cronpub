package models

type Listing struct {
	Kind     string `json:"kind"`
	MetaData struct {
		Modhash string         `json:"modhash"`
		Dist    int            `json:"dist"`
		Posts   []PostMetaData `json:"children"`
	} `json:"data"`
}

type PostMetaData struct {
	Kind string `json:"kind"`
	Post struct {
		Title string `json:"title"`
		Link  string `json:"url"`
	} `json:"data"`
}

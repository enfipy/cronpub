package controller

import (
	"encoding/json"
	"net/http"

	"github.com/enfipy/cronpub/src/helpers"
	"github.com/enfipy/cronpub/src/models"
	"github.com/enfipy/cronpub/src/services/scraper"
)

type scraperController struct {
	httpClient *http.Client
}

func NewController(httpClient *http.Client) scraper.Controller {
	return &scraperController{
		httpClient: httpClient,
	}
}

func (cnr *scraperController) FetchFromReddit(url string) []string {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	helpers.PanicOnError(err)
	req.Header.Set("User-Agent", "reddit-post")

	res, err := cnr.httpClient.Do(req)
	helpers.PanicOnError(err)

	data := new(models.Listing)
	err = json.NewDecoder(res.Body).Decode(data)
	helpers.PanicOnError(err)

	var links []string
	for _, post := range data.MetaData.Posts {
		links = append(links, post.Post.Link)
	}

	return links
}

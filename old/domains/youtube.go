package domains

import (
	"encoding/json"
	"net/http"
	"net/url"
)

const (
	ERROR_TYPE    = -1
	VIDEO_TYPE    = 0
	PLAYLIST_TYPE = 1
)

type (
	videoResponse struct {
		Formats []struct {
			Url string `json:"url"`
		} `json:"formats"`
		Title string `json:"title"`
	}

	VideoResult struct {
		Media string
		Title string
	}

	PlaylistVideo struct {
		Id string `json:"id"`
	}

	YTSearchContent struct {
		Id           string `json:"id"`
		Title        string `json:"title"`
		Description  string `json:"description"`
		ChannelTitle string `json:"channel_title"`
		Duration     string `json:"duration"`
	}

	ytApiResponse struct {
		Error   bool              `json:"error"`
		Content []YTSearchContent `json:"content"`
	}

	Youtube struct {
		Conf *Config
	}
)

func (youtube Youtube) buildUrl(query string) (*string, error) {
	base := "/v1/youtube/search"
	address, err := url.Parse(base)
	if err != nil {
		return nil, err
	}
	params := url.Values{}
	params.Add("search", query)
	address.RawQuery = params.Encode()
	str := address.String()
	return &str, nil
}

func (youtube Youtube) Search(query string) ([]YTSearchContent, error) {
	addr, err := youtube.buildUrl(query)
	if err != nil {
		return nil, err
	}
	resp, err := http.Get(*addr)
	if err != nil {
		return nil, err
	}
	var apiResp ytApiResponse
	json.NewDecoder(resp.Body).Decode(&apiResp)
	return apiResp.Content, nil
}

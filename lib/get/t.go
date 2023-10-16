package get

type ghasset struct {
	Name string `json:"name"`
	BUrl string `json:"browser_download_url"`
	Url  string `json:"url"`
}

type ghrel struct {
	TagName string    `json:"tag_name"`
	Assets  []ghasset `json:"assets"`
}

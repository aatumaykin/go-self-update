package github

type Asset struct {
	Name        string `json:"name"`
	Size        int    `json:"size"`
	DownloadURL string `json:"browser_download_url"`
}

func (a *Asset) GetName() string {
	return a.Name
}

func (a *Asset) GetSize() int {
	return a.Size
}

func (a *Asset) GetDownloadURL() string {
	return a.DownloadURL
}

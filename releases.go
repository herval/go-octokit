package octokit

import (
	"github.com/lostisland/go-sawyer/hypermedia"
	"net/url"
	"time"
)

var (
	ReleasesURL = Hyperlink("/repos/{owner}/{repo}/releases{/id}")
)

type Release struct {
	*hypermedia.HALResource

	ID              int        `json:"id,omitempty"`
	URL             string     `json:"url,omitempty"`
	HTMLURL         string     `json:"html_url,omitempty"`
	AssetsURL       string     `json:"assets_url,omitempty"`
	UploadURL       string     `json:"upload_url,omitempty"`
	TagName         string     `json:"tag_name,omitempty"`
	TargetCommitish string     `json:"target_commitish,omitempty"`
	Name            string     `json:"name,omitempty"`
	Body            string     `json:"body,omitempty"`
	Draft           bool       `json:"draft,omitempty"`
	Prerelease      bool       `json:"prerelease,omitempty"`
	CreatedAt       *time.Time `json:"created_at,omitempty"`
	PublishedAt     *time.Time `json:"published_at,omitempty"`
	Assets          []Asset    `json:"assets,omitempty"`
}

type Asset struct {
	ID            int        `json:"id,omitempty"`
	Name          string     `json:"name,omitempty"`
	Label         string     `json:"label,omitempty"`
	ContentType   string     `json:"content_type,omitempty"`
	State         string     `json:"state,omitempty"`
	Size          int        `json:"size,omitempty"`
	DownloadCount int        `json:"download_count,omitempty"`
	URL           string     `json:"url,omitempty"`
	CreatedAt     *time.Time `json:"created_at,omitempty"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty"`
}

// Create a ReleasesService with the base Hyperlink and the params M to expand the Hyperlink
// If no Hyperlink is passed in, it will use ReleasesURL.
func (c *Client) Releases(link *Hyperlink, m M) (releases *ReleasesService, err error) {
	if link == nil {
		link = &ReleasesURL
	}

	url, err := link.Expand(m)
	if err != nil {
		return
	}

	releases = &ReleasesService{client: c, URL: url}
	return
}

type ReleasesService struct {
	client *Client
	URL    *url.URL
}

func (r *ReleasesService) GetAll() (releases []Release, result *Result) {
	result = r.client.Get(r.URL, &releases)
	return
}

func (r *ReleasesService) Create(params interface{}) (release *Release, result *Result) {
	result = r.client.Post(r.URL, params, &release)
	return
}

type ReleaseParams struct {
	TagName         string `json:"tag_name,omitempty"`
	TargetCommitish string `json:"target_commitish,omitempty"`
	Name            string `json:"name,omitempty"`
	Body            string `json:"body,omitempty"`
	Draft           bool   `json:"draft,omitempty"`
	Prerelease      bool   `json:"prerelease,omitempty"`
}

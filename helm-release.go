package cloud66

import (
	"time"
)

type HelmRelease struct {
	Uid           string    `json:"uid"`
	Name          string    `json:"name"`
	Version       string    `json:"version"`
	RepositoryURL string    `json:"repository"`
	Values        string    `json:"values"`
	Body          string    `json:"body"`
	CreatedAt     time.Time `json:"created_at_iso"`
	UpdatedAt     time.Time `json:"updated_at_iso"`
}

func (p HelmRelease) String() string {
	return p.Name
}

func (c *Client) AddHelmReleases(stackUid string, formationUid string, releases []*HelmRelease, message string) ([]HelmRelease, error) {
	var releasesRes []HelmRelease = make([]HelmRelease, 0)
	var singleRes *HelmRelease
	for _, helmRelease := range releases {
		params := struct {
			Message     string       `json:"message"`
			HelmRelease *HelmRelease `json:"helm_release"`
		}{
			Message:     message,
			HelmRelease: helmRelease,
		}
		singleRes = nil

		req, err := c.NewRequest("POST", "/stacks/"+stackUid+"/formations/"+formationUid+"/helm_releases.json", params, nil)
		if err != nil {
			return nil, err
		}

		err = c.DoReq(req, &singleRes, nil)
		if err != nil {
			return nil, err
		}
		releasesRes = append(releasesRes, *singleRes)
	}

	return releasesRes, nil
}

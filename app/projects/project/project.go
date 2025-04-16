package project

import (
	"imageresizerservice/app/projects/project/projectID"
	"imageresizerservice/app/projects/project/projectName"
	"imageresizerservice/app/projects/projectRoutes"
	"imageresizerservice/app/users/userID"
	"net/url"
	"strings"
	"time"
)

type Project struct {
	ID              projectID.ProjectID
	CreatedByUserID userID.UserID
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Name            projectName.ProjectName
	AllowedDomains  []url.URL
	PageURL         string
}

func UrlLinesToUrlList(urls string) []url.URL {
	var validURLs []url.URL
	urlsList := strings.Split(urls, "\n")
	for _, urlStr := range urlsList {
		urlStr = strings.TrimSpace(urlStr)
		if urlStr == "" {
			continue
		}

		parsedURL, err := url.Parse(urlStr)
		if err == nil && parsedURL.Scheme != "" && parsedURL.Host != "" {
			validURLs = append(validURLs, *parsedURL)
		}
	}
	return validURLs
}

func (p *Project) EnsureComputed() *Project {
	p.PageURL = projectRoutes.ToProjectPage(p.ID)
	return p
}

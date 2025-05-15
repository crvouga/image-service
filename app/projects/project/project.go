package project

import (
	"imageService/app/projects/project/projectID"
	"imageService/app/projects/project/projectName"
	"imageService/app/projects/projectRoutes"
	"imageService/app/users/userID"
	"net/url"
	"strings"
	"time"
)

type Project struct {
	ID                   projectID.ProjectID
	CreatedByUserID      userID.UserID
	CreatedAt            time.Time
	UpdatedAt            time.Time
	Name                 projectName.ProjectName
	AllowedDomains       []url.URL
	URL                  string
	EditURL              string
	DeleteURL            string
	AllowedDomainsString string
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
	p.URL = projectRoutes.ToGetProject(p.ID)
	p.EditURL = projectRoutes.ToEditProject(p.ID)
	p.DeleteURL = projectRoutes.ToDeleteProject(p.ID)
	var domains []string
	for _, domain := range p.AllowedDomains {
		domains = append(domains, domain.String())
	}
	p.AllowedDomainsString = strings.Join(domains, "\n")
	return p
}

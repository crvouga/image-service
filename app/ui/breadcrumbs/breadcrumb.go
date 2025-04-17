package breadcrumbs

import (
	"fmt"
	"net/url"
)

type Breadcrumbs = []Breadcrumb

func New(items ...Breadcrumb) Breadcrumbs {
	return items
}

type Breadcrumb struct {
	Label string
	Href  string // Optional
}

// FromQueryParams parses breadcrumbs from URL query parameters.
// Breadcrumbs are stored with keys like "bc_0_label", "bc_0_href", "bc_1_label", etc.
func FromQueryParams(query url.Values) []Breadcrumb {
	var breadcrumbs []Breadcrumb
	i := 0
	for {
		label := query.Get(fmt.Sprintf("bc_%d_label", i))
		if label == "" {
			break
		}
		breadcrumbs = append(breadcrumbs, Breadcrumb{
			Label: label,
			Href:  query.Get(fmt.Sprintf("bc_%d_href", i)),
		})
		i++
	}
	return breadcrumbs
}

// ToQueryParams adds breadcrumbs to URL query parameters.
// Breadcrumbs are stored with keys like "bc_0_label", "bc_0_href", "bc_1_label", etc.
func ToQueryParams(breadcrumbs []Breadcrumb) url.Values {
	query := url.Values{}
	for i, bc := range breadcrumbs {
		query.Set(fmt.Sprintf("bc_%d_label", i), bc.Label)
		if bc.Href != "" {
			query.Set(fmt.Sprintf("bc_%d_href", i), bc.Href)
		}
	}
	return query
}

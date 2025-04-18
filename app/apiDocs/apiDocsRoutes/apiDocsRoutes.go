package apiDocsRoutes

const (
	ApiDocsPage = "/api-docs"
)

func ToApiDocsPage(endpoint string) string {
	return ApiDocsPage + "?endpoint=" + endpoint
}

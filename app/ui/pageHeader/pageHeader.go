package pageHeader

type PageHeader struct {
	Title   string
	Actions []Action
}

type Action struct {
	Label string
	URL   string
}

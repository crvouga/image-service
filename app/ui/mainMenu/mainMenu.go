package mainMenu

type MainMenu struct {
	Items []MainMenuItem
}

type MainMenuItem struct {
	Label       string
	URL         string
	Description string
}

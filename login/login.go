package login

import (
	"image-resizer-service/page"
	"net/http"
	"time"
)

type LoginPageData struct {
	Action string
}

func HandlerSendLoginLink(w http.ResponseWriter, r *http.Request) {

	time.Sleep(time.Second)
	time.Sleep(time.Second)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", page.Handler("./login/login_page.html", nil))
	mux.HandleFunc("/send-login-link", HandlerSendLoginLink)

	return mux
}

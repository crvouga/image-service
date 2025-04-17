package loginRoutes

const (
	Prefix       = "/login/"
	SendLinkPage = Prefix + "send-link"
	UseLinkPage  = Prefix + "use-link"
)

func ToSendLink(email string) string {
	return SendLinkPage + "?Email=" + email
}

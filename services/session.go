package services

var webSession string

func AddSession(email string) {
	webSession = email
}
func ReturnSession() string {
	return webSession
}

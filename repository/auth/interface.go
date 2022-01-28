package auth

type Auth interface {
	Login(email string, password string) (int, string, string, error)
}

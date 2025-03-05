package auth

type SignUpDTO struct {
	Email    string
	Username string
	Password string
	DeviceID string
}

type LogInDTO struct {
	Email    string
	Password string
	DeviceID string
}

type TokenDTO struct {
	AccessToken  string
	RefreshToken string
}

package service

type VerificationOptions struct {
	Token string `url:"token"`
}

type VerificationModel struct {
	Link string
}

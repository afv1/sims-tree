package service

type Provider struct {
    Auth *AuthService
}

func NewProvider() *Provider {
    return &Provider{
        Auth: newAuthService(),
    }
}
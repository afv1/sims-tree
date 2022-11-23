package dto

// AuthUser DTO with cred
type AuthUser struct {
    Login    string `json:"login"`
    Password string `json:"password"`
}

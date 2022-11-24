package service

import (
    "crypto/hmac"
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "github.com/afv1/sims-tree/cfg"
    "time"
)

const secret = "R6IFofjN4LzZ6YEt6IllVQrQ"

func generatePasswordHash(pass string) string {
    h := hmac.New(sha256.New, []byte(secret))
    h.Write([]byte(pass))

    return hex.EncodeToString(h.Sum(nil))
}

func generateCookie(input string) string {
    h := hmac.New(sha256.New, []byte(cfg.App.AuthSessKey))
    h.Write([]byte(input + fmt.Sprintf("%d", time.Now().Unix())))

    return hex.EncodeToString(h.Sum(nil))
}
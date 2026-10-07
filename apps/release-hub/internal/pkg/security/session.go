package security

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

const UserSessionCookie = "release_hub_user"

func SignUserSession(userID, secret string) string {
	return userID + "." + signSessionValue(userID, secret)
}

func ParseUserSession(value, secret string) (string, error) {
	parts := strings.Split(value, ".")
	if len(parts) != 2 || !hmac.Equal([]byte(parts[1]), []byte(signSessionValue(parts[0], secret))) {
		return "", fmt.Errorf("invalid user session")
	}
	if len(parts[0]) != 32 {
		return "", fmt.Errorf("invalid user session")
	}
	return parts[0], nil
}

func signSessionValue(value, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	_, _ = mac.Write([]byte(value))
	return hex.EncodeToString(mac.Sum(nil))
}

package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/DGclasher/keylang/config"
)

func ParseJSON(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("request body is empty")
	}
	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteJSON(w http.ResponseWriter, status int, payload any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(payload)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error": err.Error()})
}

func GenLicenseKey(email string, productId string) string {
	h := hmac.New(sha256.New, []byte(config.Envs.Secret))
	h.Write([]byte(email + productId))
	hmacBytes := h.Sum(nil)

	licenseKey := base64.StdEncoding.EncodeToString(hmacBytes)
	licenseKey = strings.ReplaceAll(licenseKey, "=", "")
	licenseKey = strings.ReplaceAll(licenseKey, "+", "")
	licenseKey = strings.ReplaceAll(licenseKey, "/", "")
	licenseKey = strings.ToUpper(licenseKey)
	if len(licenseKey) > 16 {
		licenseKey = licenseKey[:16]
	} else if len(licenseKey) < 16 {
		licenseKey = licenseKey + strings.Repeat("A", 16-len(licenseKey))
	}
	return licenseKey
}

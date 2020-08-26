package shared

import (
	"encoding/base64"
	"os"

	"github.com/pkg/errors"
)

// GetCredentials gets GOOGLE_APPLICATION_CREDENTIALS from env
func GetCredentials() []byte {
	p := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")

	b, err := base64.StdEncoding.DecodeString(p)
	if err != nil {
		panic(errors.Wrap(err, "Failed to decode base64 credential string"))
	}

	return b
}

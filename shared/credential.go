package shared

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

const (
	credentialPathVar = "GOOGLE_APPLICATION_CREDENTIALS"
)

// GetCredentials gets GOOGLE_APPLICATION_CREDENTIALS from env
func GetCredentials() []byte {
	p := os.Getenv(credentialPathVar)

	ud, _ := os.UserHomeDir()

	b, err := ioutil.ReadFile(filepath.Join(ud, p))
	if err != nil {
		panic(errors.Wrap(err, "Couldn't read Credential File"))
	}

	return b
}

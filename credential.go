package credential

import (
	"github.com/danieljoos/wincred"
)

type credential struct {
	username string `json:"username"`
	password string `json:"password"`
}

const credentialName = "itsscb:xtam"

func Save(username, password string) error {
	cred := wincred.NewGenericCredential(credentialName)

	cred.UserName = username
	cred.CredentialBlob = []byte(password)

	return cred.Write()
}

func read() (*wincred.GenericCredential, error) {
	cred, err := wincred.GetGenericCredential(credentialName)
	if err != nil {
		return nil, err
	}

	return cred, nil
}

func User() string {
	cred, err := read()
	if err != nil {
		return ""
	}

	return cred.UserName
}
func Password() string {
	cred, err := read()
	if err != nil {
		return ""
	}

	return string(cred.CredentialBlob)
}

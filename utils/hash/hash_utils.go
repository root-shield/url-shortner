package hash

import (
	"crypto/sha256"
	"repository.com/my_username/repo_name/utils/errors"
)

func UrlToHash(url string) ([]byte, *errors.RestErr) {
	h := sha256.New()
	_, err := h.Write([]byte(url))
	if err != nil {
		return nil, errors.NewInternalServerError("Error when trying to hash")
	}
	return h.Sum(nil), nil
}

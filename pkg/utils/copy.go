package utils

import (
	"github.com/quangdangfit/go-admin/pkg/errors"
)

// Copy by marshal json data
func Copy(dest interface{}, src interface{}) error {
	data, err := json.Marshal(src)
	if err != nil {
		return errors.ErrorMarshal.Newm(err.Error())
	}

	err = json.Unmarshal(data, dest)
	if err != nil {
		return errors.ErrorUnmarshal.Newm(err.Error())
	}

	return nil
}

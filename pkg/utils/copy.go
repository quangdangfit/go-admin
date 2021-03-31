package utils

import (
	"github.com/quangdangfit/gosdk/utils/logger"
)

// Copy by marshal json data
func Copy(dest interface{}, src interface{}) error {
	data, err := json.Marshal(src)
	if err != nil {
		logger.Error("Failed to marshal data")
		return err
	}

	json.Unmarshal(data, dest)

	return nil
}

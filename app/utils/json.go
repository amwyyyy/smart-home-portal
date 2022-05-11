package utils

import (
	"com/denis/smarthome/app/logger"
	"encoding/json"
	"strings"
)

func ToString(req interface{}) string {
	buf, err := json.Marshal(req)
	if err != nil {
		logger.Errorf("json obj toString error: %s", err)
	}

	str := string(buf)
	if strings.HasPrefix(str, "\"") {
		if strings.HasSuffix(str, "\"") {
			str = strings.TrimPrefix(str, "\"")
			str = strings.TrimSuffix(str, "\"")
		}
	}
	return str
}

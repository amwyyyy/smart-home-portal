package utils

import (
	"testing"
)

func TestToString(t *testing.T) {
	m := make(map[string]interface{})
	m["name"] = "denis"
	m["age"] = 13

	s := ToString(m)
	if s != "{\"age\":13,\"name\":\"denis\"}" {
		t.Error("json parse error!")
	}
}

package system

import (
	"encoding/json"
	"testing"
)

func TestSystemInfo(t *testing.T) {
	data := GetNowSystemInfo()
	SJson, err := json.Marshal(data)
	t.Log("SystemInfo: ", data)
	t.Log("json: ", string(SJson), err)
}

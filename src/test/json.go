package test

import (
	"encoding/json"
	"fmt"
)

func TestJson() {
	jsonStr := `["nihao", "hello", "afafa", "\"www\".com"]`
	var decodeStr []string
	json.Unmarshal([]byte(jsonStr), &decodeStr)
	fmt.Println(decodeStr)
}

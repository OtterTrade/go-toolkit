package feishumsg

import (
	"fmt"
	"net/http"
	"strings"
)

func SendMsg(api_url string, msg string) {
	// json
	contentType := "application/json"

	// data
	sendData := `{
    "msg_type": "text",
    "content": {"text": "` + "消息通知: " + msg + `"}
   }`

	// request
	result, err := http.Post(api_url, contentType, strings.NewReader(sendData))
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return
	}
	defer result.Body.Close()
}

package feishumsg

import (
	"fmt"
	"net/http"
	"strings"
)

func SendTxtMsg(api_url string, msg string) {
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

func SendMsg(api_url string, lvl int, msg string) {
	// json
	contentType := "application/json"
   
   // color and content
	var color = "blue"
	var content = "提示消息"
	if lvl == 1 {
		color = "yellow"
		content = "告警消息"
	} else if lvl == 2 {
		color = "red"
		content = "错误消息"
	}
   
	sendData := `{
      "msg_type": "interactive",
      "card": {
         "config": {
            "wide_screen_mode": true
         },
         "header": {
            "title": {
               "tag": "plain_text",
               "content": "` + content + `"
            },
            "template": "` + color + `"
         },
         "elements": [{
            "tag": "div",
            "text": {
                "tag": "lark_md",
                "content": "` + msg + `"
            }
         }]

      }
   }`

	// request
	result, err := http.Post(api_url, contentType, strings.NewReader(sendData))
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return
	}
	defer result.Body.Close()

}

func Info(api_url string, msg string) {
	SendMsg(api_url, 0, msg)
}

func Warn(api_url string, msg string) {
	SendMsg(api_url, 1, msg)
}

func Error(api_url string, msg string) {
	SendMsg(api_url, 2, msg)
}

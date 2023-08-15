package feishu

/*
example:
API_URL string = "https://open.feishu.cn/open-apis/bot/v2/hook/74fe8f5b-e77e-4547-b994-xxxxxxxxxxxx"
fmsg := feishu.NewFeishuMsg(API_URL)
fmsg.Info("welcome to feishu group")
fmsg.Warn("请留意告警")
fmsg.Error("发生严重错误")
*/

import (
	"fmt"
	"net/http"
	"strings"
)

type FeishuMsg struct {
	api_url string
}

func NewFeishuMsg(api_url string) *FeishuMsg {
	p := new(FeishuMsg)
	p.api_url = api_url
	return p
}

func (fsm FeishuMsg) SendTxtMsg(msg string) {
	// json
	contentType := "application/json"

	// data
	sendData := `{
		"msg_type": "text",
		"content": {"text": "` + "消息通知: " + msg + `"}
	}`

	// request
	result, err := http.Post(fsm.api_url, contentType, strings.NewReader(sendData))
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return
	}
	defer result.Body.Close()
}

func (fsm FeishuMsg) SendMsg(lvl int, msg string) {
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
	result, err := http.Post(fsm.api_url, contentType, strings.NewReader(sendData))
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return
	}
	defer result.Body.Close()

}

func (fsm FeishuMsg) Info(msg string) {
	fsm.SendMsg(0, msg)
}

func (fsm FeishuMsg) Warn(msg string) {
	fsm.SendMsg(1, msg)
}

func (fsm FeishuMsg) Error(msg string) {
	fsm.SendMsg(2, msg)
}

/*
@Time : 2019/9/9 15:18
@Author : Tux
@File : message
@Description :
*/

package dingrobot

const (
	msgTypeText       = "text"
	msgTypeLink       = "link"
	msgTypeMarkdown   = "markdown"
	msgTypeActionCard = "actionCard"
)

type textParams struct {
	Content string `json:"content"`
}

type atParams struct {
	AtMobiles []string `json:"atMobiles,omitempty"`
	IsAtAll   bool     `json:"isAtAll,omitempty"`
}

// textMessage 文本消息
type textMessage struct {
	MsgType string     `json:"msgtype"`
	Text    textParams `json:"text"`
	At      atParams   `json:"at"`
}

type linkParams struct {
	Title      string `json:"title"`
	Text       string `json:"text"`
	MessageURL string `json:"messageUrl"`
	PicURL     string `json:"picUrl,omitempty"`
}

// linkMessage link类型
type linkMessage struct {
	MsgType string     `json:"msgtype"`
	Link    linkParams `json:"link"`
}

type markdownParams struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

// markdownMessage markdown消息
type markdownMessage struct {
	MsgType  string         `json:"msgtype"`
	MarkDown markdownParams `json:"markdown"`
	At       atParams       `json:"at"`
}

type actionCardParams struct {
	Title          string `json:"title"`
	Text           string `json:"text"`
	SingleTitle    string `json:"singleTitle"`
	SingleURL      string `json:"singleUrl"`
	BtnOrientation string `json:"btnOrientation,omitempty"`
	HideAvatar     string `json:"hideAvatar,omitempty"`
}

// actionCardMessage 整体跳转ActionCard类型
type actionCardMessage struct {
	MsgType    string           `json:"msgtype"`
	ActionCard actionCardParams `json:"actionCard"`
}

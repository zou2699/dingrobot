/*
@Time : 2019/9/9 16:03
@Author : Tux
@File : dingrobot_test.go
@Description :
*/

package dingrobot

import (
	"reflect"
	"testing"
)

const webhook = "https://oapi.dingtalk.com/robot/send?access_token=xxxx"

func TestNewRobot(t *testing.T) {
	type args struct {
		webhook string
	}
	var tests = []struct {
		name string
		args args
		want Roboter
	}{
		{
			name: "",
			args: args{
				webhook: webhook,
			},
			want: &Robot{Webhook: webhook},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRobot(tt.args.webhook); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRobot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRobot_SendActionCard(t *testing.T) {
	type fields struct {
		Webhook string
	}
	type args struct {
		title          string
		text           string
		singleTitle    string
		singleURL      string
		btnOrientation string
		hideAvatar     string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "sendActionCard",
			fields: fields{Webhook: webhook},
			args: args{
				title:          "乔布斯 20 年前想打造一间苹果咖啡厅，而它正是 Apple Store 的前身",
				text:           "![screenshot](@lADOpwk3K80C0M0FoA) \n ### 乔布斯 20 年前想打造的苹果咖啡厅 \n Apple Store 的设计正从原来满满的科技感走向生活化，而其生活化的走向其实可以追溯到 20 年前苹果一个建立咖啡馆的计划",
				singleURL:      "https://www.dingtalk.com/",
				singleTitle:    "阅读全文",
				btnOrientation: "0",
				hideAvatar:     "0",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Robot{
				Webhook: tt.fields.Webhook,
			}
			if err := r.SendActionCard(tt.args.title, tt.args.text, tt.args.singleTitle, tt.args.singleURL, tt.args.btnOrientation, tt.args.hideAvatar); (err != nil) != tt.wantErr {
				t.Errorf("SendActionCard() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRobot_SendLink(t *testing.T) {
	type fields struct {
		Webhook string
	}
	type args struct {
		title      string
		text       string
		messageURL string
		picURL     string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "SendLink",
			fields: fields{
				Webhook: webhook,
			},
			args: args{
				title:      "时代的火车向前开",
				text:       "这个即将发布的新版本，创始人陈航（花名“无招”）称它为“红树林”。 \n 而在此之前，每当面临重大升级，产品经理们都会取一个应景的代号，这一次，为什么是“红树林”？",
				messageURL: "https://www.dingtalk.com/s?__biz=MzA4NjMwMTA2Ng==&mid=2650316842&idx=1&sn=60da3ea2b29f1dcc43a7c8e4a7c97a16&scene=2&srcid=09189AnRJEdIiWVaKltFzNTw&from=timeline&isappinstalled=0&key=&ascene=2&uin=&devicetype=android-23&version=26031933&nettype=WIFI",
				picURL:     "",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Robot{
				Webhook: tt.fields.Webhook,
			}
			if err := r.SendLink(tt.args.title, tt.args.text, tt.args.messageURL, tt.args.picURL); (err != nil) != tt.wantErr {
				t.Errorf("SendLink() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRobot_SendMarkdown(t *testing.T) {
	type fields struct {
		Webhook string
	}
	type args struct {
		title     string
		text      string
		atMobiles []string
		isAtAll   bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "SendMarkdown",
			fields:  fields{
				Webhook: webhook,
			},
			args:    args{
				title:     "杭州天气",
				text:      "#### 杭州天气 @156xxxx8827\n" +
					"> 9度，西北风1级，空气良89，相对温度73%\n\n" +
					"> ![screenshot](https://gw.alicdn.com/tfs/TB1ut3xxbsrBKNjSZFpXXcXhFXa-846-786.png)\n"  +
					"> ###### 10点20分发布 [天气](http://www.thinkpage.cn/) \n",
				atMobiles: []string{"1568827", "1898325"},
				isAtAll:   false,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Robot{
				Webhook: tt.fields.Webhook,
			}
			if err := r.SendMarkdown(tt.args.title, tt.args.text, tt.args.atMobiles, tt.args.isAtAll); (err != nil) != tt.wantErr {
				t.Errorf("SendMarkdown() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRobot_SendText(t *testing.T) {
	type fields struct {
		Webhook string
	}
	type args struct {
		content   string
		atMobiles []string
		isAtAll   bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "SendText",
			fields:  fields{
				Webhook:webhook,
			},
			args:    args{
				content:   "我就是我, 是不一样的烟火@156xxxx8827",
				atMobiles: []string{"156xxxx8827","156xxxx8821"},
				isAtAll:   false,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Robot{
				Webhook: tt.fields.Webhook,
			}
			if err := r.SendText(tt.args.content, tt.args.atMobiles, tt.args.isAtAll); (err != nil) != tt.wantErr {
				t.Errorf("SendText() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}


## 钉钉机器人消息发送接口

### 支持发送类型
- text类型 SendText
- link类型 SendLink
- markdown类型 SendMarkdown
- 整体跳转ActionCard类型 SendActionCard

### Usage
```go
func main() {
	robot := dingrobot.NewRobot("https://oapi.dingtalk.com/robot/send?access_token=xxxx")
	err := robot.SendText("测试", nil, false) 
	if err != nil {
		fmt.Println(err)
	}

}
```
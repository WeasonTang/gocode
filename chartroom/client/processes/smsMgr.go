package processes

import (
	"encoding/json"
	"fmt"
	"gocode/chartroom/common/message"
)

func outPutGroupMes(mes *message.Message) {

	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(mes.Data), &smsMes) err=", err)
		return
	}
	//显示信息
	info := fmt.Sprintf("用户id:\t%d 发送消息:\t%s",
		smsMes.User.UserId, smsMes.Content)
	fmt.Println(info)
	fmt.Println()
}

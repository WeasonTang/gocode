package processes

import (
	"fmt"
	"gocode/chartroom/client/entity"
	"gocode/chartroom/common/message"
	"gocode/chartroom/server/model"
)

//客户端要维护的map
var onlineUsers = make(map[int]*model.User, 10)
var curUser entity.CurUser

//在客户端显示当前在线的用户
func outPutOnlineUser() {
	fmt.Println("当前用户列表")
	for id, user := range onlineUsers {
		fmt.Printf("用户id:\t%d 状态:%d\n", id, user.UserStatus)
	}
}

//更新用户状态
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {

	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if ok {
		user.UserStatus = notifyUserStatusMes.Status
	} else {
		user = &model.User{
			UserId:     notifyUserStatusMes.UserId,
			UserStatus: notifyUserStatusMes.Status,
		}
		onlineUsers[user.UserId] = user
	}

}

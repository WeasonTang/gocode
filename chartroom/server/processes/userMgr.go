package processes

import "fmt"

//UserMgr实例在服务器端有且只有一个
//很多地方会使用到 因此定义为全局变量
var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

//初始化UserMgr
func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

//完成对onlineUsers添加
func (userMgr *UserMgr) AddOnlineUser(up *UserProcess) {
	userMgr.onlineUsers[up.UserId] = up
}

//删除
func (userMgr *UserMgr) DelOnlineUser(userId int) {
	delete(userMgr.onlineUsers, userId)
}

//返回当前所有在线用户
func (userMgr *UserMgr) GetAllOnlineUser() map[int]*UserProcess {
	return userMgr.onlineUsers
}

//根据id返回对应的UserProcess
func (userMgr *UserMgr) GetOnlineUserById(userId int) (up *UserProcess, err error) {
	up, ok := userMgr.onlineUsers[userId]
	if !ok {
		err = fmt.Errorf("用户%d 不存在", userId)
	}
	return
}

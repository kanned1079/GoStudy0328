package users

const (
	Unknow_Error       = -1   // 未知错误
	User_Exist         = iota // 用户存在
	User_Not_Exist            // 用户不存在
	Password_Incorrect        // 密码错误
	Auth_Pass                 // 凭据验证通过
	Auth_Failure              // 凭据验证失败
	Query_Pass                // 请求通过
	Query_Failure             // 请求失败
	Insert_Success            // 插入数据成功
	Insert_Failure            // 插入数据失败
	Update_Success            // 更新成功
	Update_Failure            // 更新失败
	Delete_Success            // 删除成功
	Delete_Failure            // 删除失败
)

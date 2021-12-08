package errcode

var (
	ErrorCreateUserFail    = NewError(4018001, "创建用户失败")
	ErrorUserExist         = NewError(4018002, "用户已存在")
	ErrorUserNotExist      = NewError(4018003, "用户未存在")
	ErrorPasswdWrong       = NewError(4018004, "密码错误")
	ErrorUpdateFail        = NewError(4018005, "更新用户失败")
	ErrorCreateProgramFail = NewError(5011001, "创建题目失败")
	ErrorReturnProgramListFail = NewError(5011002, "获取题目列表失败")
	ErrorReturnProgramDetail = NewError(5011003, "获取题目详情失败")
	ErrorSubmitProgramFail = NewError(5011004,"提交题目失败")
	ErrorReturnRecord = NewError(5011005,"获取提交记录失败")
	ErrorDeleteProgram = NewError(5011006,"删除题目失败")
	ErrorCreateEmailFail = NewError(5011007,"添加邮箱失败")
)

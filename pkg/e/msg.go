package e

//MsgFlags is the map of error conent to error code.
var MsgFlags = map[int]string{
	SUCCESS:                        "ok",
	ERROR:                          "fail",
	INVALID_PARAMS:                 "请求参数错误",
	JSON_PARAMS_ERROR:              "JSON参数错误",
	WEBSOCKET_ERROR:                "WebSocket错误",
	CONTENT_TYPE_IS_NOT_JSON:       "请求格式错误",
	CONVERT_DATA_TO_JSON_FAILED:    "将结果转换到JSON失败",
	NO_ADMIN_RECORD_FOUND:          "没有找到指定的Admin",
	ADMIN_LOGIN_FAILED:             "Admin登录失败",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "验证Token失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "验证Token超时",
	ERROR_AUTH_TOKEN:               "处理Token错误",
	ERROR_AUTH:                     "处理验证错误",
	CREATE_FAILED:                  "创建失败",
	UPDATE_FAILED:                  "更新失败",
	REMOVE_FAILED:                  "删除失败",
	GET_TOTAL_FAILED: "获取总数失败",
	NO_FUNC_RECORD_FOUND:           "没有找到指定的SystemFunction",
	NO_ROLE_RECORD_FOUND:           "没有找到指定的Role",
	NO_CATEGORY_RECORD_FOUND:       "没有找到指定的Category",
	NO_SUB_CATEGORY_RECORD_FOUND:   "没有找到指定的Sub-Category",
	NO_CONTENT_RECORD_FOUND:        "没有找到指定的Content",
	NO_ITEM_RECORD_FOUND:           "没有找到指定的Item",
	NO_HOME_RECORD_FOUND:           "没有找到指定的Home",
}

//GetMsg get message conent by error code.
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}

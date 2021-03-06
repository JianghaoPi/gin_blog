package error

var CodeToMessage = map[int]string {
	SUCCESS: "success",
	INTERNAL_ERROR: "internal error",
	INVALID_PARAMS: "请求参数错误",
	ERROR_EXIST_TAG: "标签已存在",
	ERROR_NOT_EXIST_TAG: "标签不存在",
	ERROR_NOT_EXIST_ARTICLE: "文章不存在",
	ERROR_AUTH_CHECK_TOKEN_FAIL: "token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "token已超时",
	ERROR_AUTH_TOKEN: "token生成失败",
	ERROR_AUTH: "token错误",
}

func GetMsg(code int) string {
	msg, ok := CodeToMessage[code]
	if ok {
		return msg
	}
	return CodeToMessage[INTERNAL_ERROR]
}

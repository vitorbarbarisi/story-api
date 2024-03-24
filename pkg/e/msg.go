package e

var MsgFlags = map[int]string{
	SUCCESS:                 "ok",
	ERROR:                   "fail",
	INVALID_PARAMS:          "请求参数错误",
	ERROR_EXIST_STORY:       "已存在该标签名称",
	ERROR_EXIST_STORY_FAIL:  "获取已存在标签失败",
	ERROR_NOT_EXIST_STORY:   "该标签不存在",
	ERROR_GET_STORIES_FAIL:  "获取所有标签失败",
	ERROR_COUNT_STORY_FAIL:  "统计标签失败",
	ERROR_ADD_STORY_FAIL:    "新增标签失败",
	ERROR_UPDATE_STORY_FAIL: "修改标签失败",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}

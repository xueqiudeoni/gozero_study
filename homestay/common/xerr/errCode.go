package xerr

const OK uint32 = 200
const (
	SEVER_COMMON_ERROR            uint32 = 100001
	REQUEST_PARAM_ERROR           uint32 = 100002
	TOKEN_EXPIRE_ERROR            uint32 = 100003
	TOKEN_GENERATE_ERROR          uint32 = 100004
	DB_ERROR                      uint32 = 100005
	DB_UPDATE_AFFECTED_ZERO_ERROR uint32 = 100006
)

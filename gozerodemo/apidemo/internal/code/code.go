package code

const (
	SUCCESS                    = 200
	ERROR                      = 500
	TOKEN_ERROR                = 401
	NOT_FOUND                  = 404
	PARAMS_ERROR               = 10101
	ADMIN_NOT_FOUND            = 10102
	ADMIN_NOT_EXIST            = 10103
	USER_EXIST                 = 10104
	USER_NOT_EXIST             = 10105
	USER_CREATE_ERROR          = 10106
	TASK_CREATE_FAIL           = 10201
	CREATE_SHOP_ITEM_ERROR     = 10301
	UPDATE_SHOP_ITEM_ERROR     = 10302
	USERNAME_OR_PASSWORD_ERROR = 10401
	USERNAME_ALREADY_EXISTS    = 10402
	CREATE_TOKEN_ERROR         = 10501
)

func Text(code int) string {
	return zhCNText[code]
}

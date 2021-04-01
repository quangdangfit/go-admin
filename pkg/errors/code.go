package errors

// ErrorCodeMap define map error and error message
var ErrorCodeMap = map[ErrorType]string{
	Success:                    "SUCCESS",
	Error:                      "ERROR",
	InvalidParams:              "INVALID_PARAMS",
	ErrorAuthCheckTokenFail:    "ERROR_AUTH_CHECK_TOKEN_FAIL",
	ErrorAuthCheckTokenTimeout: "ERROR_AUTH_CHECK_TOKEN_TIMEOUT",
	ErrorAuthToken:             "ERROR_AUTH_TOKEN",
	ErrorAuth:                  "ERROR_AUTH",
	ErrorInternalServer:        "ERROR_INTERNAL_SERVER",
	ErrorExistEmail:            "ERROR_EXIST_EMAIL",
	ErrorBadRequest:            "ERROR_BAD_REQUEST",
	ErrorInvalidParent:         "ERROR_INVALID_PARENT",
	ErrorAllowDeleteWithChild:  "ERROR_ALLOW_DELETE_WITH_CHILD",
	ErrorNotAllowDelete:        "ERROR_NOT_ALLOW_DELETE",
	ErrorInvalidOldPass:        "ERROR_INVALID_OLD_PASS",
	ErrorNotFound:              "ERROR_NOT_FOUND",
	ErrorPasswordRequired:      "ERROR_PASSWORD_REQUIRED",
	ErrorExistMenuName:         "ERROR_EXIST_MENU_NAME",
	ErrorUserDisabled:          "ERROR_USER_DISABLED",
	ErrorNoPermission:          "ERROR_NO_PERMISSION",
	ErrorMethodNotAllow:        "ERROR_METHOD_NOT_ALLOW",
	ErrorTooManyRequest:        "ERROR_TOO_MANY_REQUEST",
	ErrorLoginFailed:           "ERROR_LOGIN_FAILED",
	ErrorExistRole:             "ERROR_EXIST_ROLE",
	ErrorNotExistUser:          "ERROR_NOT_EXIST_USER",
	ErrorExistRoleUser:         "ERROR_EXIST_ROLE_USER",
	ErrorNotExistRole:          "ERROR_NOT_EXIST_ROLE",
	ErrorTokenExpired:          "ERROR_TOKEN_EXPIRED",
	ErrorTokenInvalid:          "ERROR_TOKEN_INVALID",
	ErrorTokenMalformed:        "ERROR_TOKEN_MALFORMED",
}

// GetCode get error code
func GetCode(status int) string {
	msg, ok := ErrorCodeMap[ErrorType(status)]
	if ok {
		return msg
	}
	return ErrorCodeMap[Error]
}

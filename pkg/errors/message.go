package errors

// MsgMap message map
var MsgMap = map[ErrorType]string{
	Success:                    "OK",
	InvalidParams:              "Request parameter error - %s",
	ErrorAuthCheckTokenFail:    "Token authentication failed",
	ErrorAuthCheckTokenTimeout: "Token time out",
	ErrorAuthToken:             "Token build failed",
	ErrorAuth:                  "Token error",
	Error:                      "Error occurred",
	ErrorInternalServer:        "Server error",
	ErrorExistEmail:            "The Email Address entered already exists in the system",
	ErrorBadRequest:            "Request error",
	ErrorInvalidParent:         "Invalid parent node",
	ErrorAllowDeleteWithChild:  "Contains children, cannot be deleted",
	ErrorNotAllowDelete:        "Resources are not allowed to be deleted",
	ErrorInvalidOldPass:        "Old password is incorrect",
	ErrorNotFound:              "Resource does not exist",
	ErrorPasswordRequired:      "Password is required",
	ErrorExistMenuName:         "Menu name already exists",
	ErrorUserDisabled:          "User is disabled, please contact administrator",
	ErrorNoPermission:          "No access",
	ErrorMethodNotAllow:        "Method is not allowed",
	ErrorTooManyRequest:        "Requests are too frequent",
	ErrorLoginFailed:           "Email or password is invalid",
	ErrorExistRole:             "Role name already exists",
	ErrorNotExistUser:          "Account is invalid",
	ErrorExistRoleUser:         "The role has been given to the user and is not allowed to be deleted",
	ErrorNotExistRole:          "Role user is disabled, please contact administrator",
	ErrorTokenExpired:          "Token is expired",
	ErrorTokenInvalid:          "Token is invalid",
	ErrorTokenMalformed:        "That's not even a token",
}

// GetMsg from status
func GetMsg(status int) string {
	msg, ok := MsgMap[ErrorType(status)]
	if ok {
		return msg
	}
	return MsgMap[Error]
}

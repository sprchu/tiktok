package errno

const (
	SuccessCode    = "0"
	ServiceErrCode = "10001"
	ParamErrCode   = "10002"
	AuthErrCode    = "10003"

	LoginErrCode            = "20001"
	UserNotExistErrCode     = "20002"
	UserAlreadyExistErrCode = "20003"
	RegisterErrCode         = "20004"
	GetUserInfoErrCode      = "20005"

	UploadErrCode        = "30001"
	FeedErrCode          = "30002"
	PublishActionErrCode = "30003"
	PublishListErrCode   = "30004"

	FollowActionErrCode   = "40001"
	FollowListErrCode     = "40002"
	FollowerListErrCode   = "40003"
	FavoriteActionErrCode = "40004"
	FavoriteListErrCode   = "40005"
	CommentActionErrCode  = "40006"
	CommentListErrCode    = "40007"
)

type ErrNo struct {
	ErrCode string
	ErrMsg  string
}

func (e ErrNo) Error() string {
	return e.ErrMsg
}

func NewErrNo(code string, msg string) error {
	return ErrNo{code, msg}
}

var (
	Success    = NewErrNo(SuccessCode, "Success")
	ServiceErr = NewErrNo(ServiceErrCode, "Service is unable to start successfully")
	ParamErr   = NewErrNo(ParamErrCode, "Wrong Parameter has been given")
	AuthErr    = NewErrNo(AuthErrCode, "Invalid token")

	LoginErr            = NewErrNo(LoginErrCode, "Wrong username or password")
	UserNotExistErr     = NewErrNo(UserNotExistErrCode, "User does not exists")
	UserAlreadyExistErr = NewErrNo(UserAlreadyExistErrCode, "User already exists")
)

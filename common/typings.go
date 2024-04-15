package common

const (
	SMS  = "SMS"
	User = "USER"
)

type RequestType struct {
	Data string `json:"data"`
	Type string `json:"type"`
}

type ResponseType struct {
	Msg  string `json:"msg"`
	Code uint32 `json:"code"`
	Data string `json:"data"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRes struct {
	Token string `json:"token"`
}

type RegistryReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegistryRes struct {
}

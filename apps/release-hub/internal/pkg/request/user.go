package request

type User struct {
	PublicId        string `json:"publicId"`
	Token           string `json:"token,omitempty"`
	AccessToken     string `json:"accessToken,omitempty"`
	Email           string `json:"email"`
	GitProviderType int    `json:"gitProviderType"` // 当前只有github
}

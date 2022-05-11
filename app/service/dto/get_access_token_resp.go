package dto

type GetAccessTokenResp struct {
	Value AccessToken `json:"value"`
}

type AccessToken struct {
	AccessToken string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	ExpireTime int64 `json:"expireTime"`
}

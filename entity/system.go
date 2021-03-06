/*
Copyright 2020 RS4
@Author: Weny Xu
@Date: 2021/01/10 18:04
*/

package entity

type AlipaySystemOauthTokenResponse struct {
	AlipaySystemOauthToken `json:"alipay_system_oauth_token_response"`
	*ErrorResponse         `json:"error_response"`
	AlipayCertSn           string `json:"alipay_cert_sn"`
	Sign                   string `json:"sign"`
}

type AlipaySystemOauthToken struct {
	UserID       string `json:"user_id"`
	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	ReExpiresIn  int64  `json:"re_expires_in"`
}

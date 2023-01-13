package user

import (
	"crypto/sha1"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reservation/constant"
	"reservation/github.com/reservation/api"
)

type WXLoginResp struct {
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

// WXLogin 这个函数以 code 作为输入, 返回调用微信接口得到的对象指针和异常情况
func WXLogin(req *api.WXLoginReq) (*api.WXLoginResp, error) {
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	// 合成url, 这里的appId和secret是在微信公众平台上获取的
	url = fmt.Sprintf(url, "wx06797c8809af4258", "b0720ee70c3c1e1a152b754af8657ae8", req.Code)

	// 创建http get请求
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 解析http请求中body 数据到我们定义的结构体中
	wxResp := WXLoginResp{}
	decoder := json.NewDecoder(resp.Body)
	fmt.Println("resp.Body: ", resp.Body)
	if err := decoder.Decode(&wxResp); err != nil {
		return nil, err
	}
	// 判断微信接口返回的是否是一个异常情况
	if wxResp.ErrCode != 0 {
		return nil, errors.New(fmt.Sprintf("ErrCode:%s  ErrMsg:%s", wxResp.ErrCode, wxResp.ErrMsg))
	}
	return &api.WXLoginResp{
		Data: &api.WXLoginResp_Data{
			OpenId:     wxResp.OpenId,
			SessionKey: wxResp.SessionKey,
			UnionId:    wxResp.UnionId,
		},
		Success:   true,
		ErrorCode: constant.SUCCESS_ERROR_CODE,
	}, nil
}

//// AppletWeChatLogin /wechat/applet_login?code=xxx [get]  路由
//// 微信小程序登录
//func AppletWeChatLogin(c *gin.Context) {
//	code := c.Query("code") //  获取code
//	// 根据code获取 openID 和 session_key
//	wxLoginResp, err := WXLogin(code)
//	if err != nil {
//		c.JSON(400, util.Fail(err.Error()))
//		return
//	}
//	// 保存登录态
//	session := sessions.Default(c)
//	session.Set("openid", wxLoginResp.OpenId)
//	session.Set("sessionKey", wxLoginResp.SessionKey)
//
//	// 这里用openid和sessionkey的串接 进行MD5之后作为该用户的自定义登录态
//	mySession := GetMD5Encode(wxLoginResp.OpenId + wxLoginResp.SessionKey)
//	// 接下来可以将openid 和 sessionkey, mySession 存储到数据库中,
//	// 但这里要保证mySession 唯一, 以便于用mySession去索引openid 和sessionkey
//	c.String(200, mySession)
//}

//// GetMD5Encode 将一个字符串进行MD5加密后返回加密后的字符串
//func GetMD5Encode(data string) string {
//	h := md5.New()
//	h.Write([]byte(data))
//	return hex.EncodeToString(h.Sum(nil))
//}
//
//// ValidateUserInfo 校验微信返回的用户数据
//func ValidateUserInfo(rawData, sessionKey, signature string) bool {
//	signature2 := GetSha1(rawData + sessionKey)
//	return signature == signature2
//}

// GetSha1 SHA-1 加密
func GetSha1(str string) string {
	data := []byte(str)
	has := sha1.Sum(data)
	res := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return res
}

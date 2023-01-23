package user

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"
	"net/http"
	"reservation/constant"
	"reservation/dal"
	"reservation/github.com/reservation/api"
	"reservation/model"
	"strconv"
)

func InitLogger() {
	// 通过 zap.NewProduction() 创建一个 logger
	logger, _ = zap.NewProduction()
}

var logger *zap.Logger

// CreateToken 生成 JWT
func CreateToken(id string) string {
	fmt.Println("CreateToken id", id)
	mySigningKey := []byte(id)
	// Create the Claims
	claims := &jwt.StandardClaims{
		ExpiresAt: 15000,
		Issuer:    "test",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	fmt.Printf("%v %v", ss, err)
	return ss
}

type WXLoginResp struct {
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

// WXLogin 这个函数以 code 作为输入, 返回调用微信接口得到的对象指针和异常情况
func WXLogin(req *api.WXLoginReq) (*api.WXLoginResp, error) {
	db := dal.Init()
	var user model.UserDO

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
	if err := decoder.Decode(&wxResp); err != nil {
		return nil, err
	}

	// 判断微信接口返回的是否是一个异常情况
	if wxResp.ErrCode != 0 {
		return nil, errors.New(fmt.Sprintf("ErrCode:%s  ErrMsg:%s", wxResp.ErrCode, wxResp.ErrMsg))
	}
	encrypted, _ := base64.StdEncoding.DecodeString(req.EncryptedData)
	keyB, _ := base64.StdEncoding.DecodeString(wxResp.SessionKey)
	ivB, _ := base64.StdEncoding.DecodeString(req.Iv)
	decrypted := AesDecryptCBC(encrypted, keyB, ivB)
	returnMap := make(map[string]interface{})
	json.Unmarshal(decrypted, &returnMap)
	fmt.Println("returnMap:", returnMap, " wxResp.OpenId: ", wxResp.OpenId)
	e := db.Model(&model.UserDO{}).Where("open_id = ? ", wxResp.OpenId).Find(&user).Limit(1)
	var id int64
	if e.Error != nil {
		fmt.Println("failed to get record by openId: ", e.Error.Error())
		//用户不存在，MySQL中创建一个用户记录

		user := model.UserDO{
			NickName:    "sdsdfds",
			PhoneNumber: "13166661111",
			OpenID:      wxResp.OpenId,
		}
		db.Create(&user)
		id = user.ID
		fmt.Println("创建User成功, id: ", id)
	} else {
		//用户存在，直接返回id
		id = user.ID
	}
	return &api.WXLoginResp{
		Data: &api.WXLoginResp_Data{
			Token: CreateToken(strconv.Itoa(int(id))),
		},
		Success:   true,
		ErrorCode: constant.SUCCESS_ERROR_CODE,
	}, nil
}
func AesDecryptCBC(encrypted []byte, key []byte, iv []byte) (decrypted []byte) {
	block, _ := aes.NewCipher(key) // 分组秘钥
	//blockSize := block.BlockSize()  // 获取秘钥块的长度
	blockMode := cipher.NewCBCDecrypter(block, iv) // 加密模式
	decrypted = make([]byte, len(encrypted))       // 创建数组
	fmt.Println("len(encrypted): ", len(encrypted), " encrypted: ", string(encrypted), " key: ", string(key), " iv: ", iv)
	blockMode.CryptBlocks(decrypted, encrypted) // 解密
	decrypted = pkcs7UnPadding(decrypted)       // 去除补全码
	return decrypted
}

func pkcs7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
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

func Register(req *api.RegisterReq) (*api.RegisterResp, error) {
	return nil, nil
}

// EditUser 编辑用户
func EditUser(req *api.EditUserReq) (*api.EditUserResp, error) {
	db := dal.Init()
	err := db.Model(model.UserDO{}).Where("id = ?", req.Id).Updates(model.UserDO{
		HeadImage:       req.HeadImage,
		NickName:        req.NickName,
		Gender:          req.Gender.String(),
		Height:          req.Height,
		Weight:          req.Weight,
		Hometown:        req.Hometown,
		Location:        req.Location,
		EmotionalStatus: req.EmotionalStatus.String(),
		Education:       req.Education.String(),
		University:      req.University,
		Occupation:      req.Occupation,
		Company:         req.Company,
		WechatNumber:    req.WechatNumber,
		PhoneNumber:     req.PhoneNumber,
	}).Error
	if err == nil {
		return &api.EditUserResp{
			Success:   true,
			ErrorCode: constant.SUCCESS_ERROR_CODE,
		}, nil
	} else {
		return &api.EditUserResp{
			Success:   false,
			ErrorCode: constant.EDIT_RECORD_ERROR,
			ErrorMsg:  err.Error(),
		}, nil
	}
}

// QueryUserInfo 查询用户信息
func QueryUserInfo(req *api.QueryUserInfoReq) (*api.QueryUserInfoResp, error) {
	db := dal.Init()
	var user model.UserDO
	db.Model(&model.UserDO{}).Where("id = ? ", req.Id).Find(&user).Limit(1)
	fmt.Println("user: ", user)
	u := api.User{
		Id:              user.ID,
		HeadImage:       user.HeadImage,
		NickName:        user.NickName,
		Gender:          api.Gender(api.Gender_value[user.Gender]),
		Height:          user.Height,
		Weight:          user.Weight,
		Hometown:        user.Hometown,
		Location:        user.Location,
		EmotionalStatus: api.EmotionalStatus(api.EmotionalStatus_value[user.EmotionalStatus]),
		Education:       api.Education(api.Education_value[user.Education]),
		University:      user.University,
		WechatNumber:    user.WechatNumber,
		PhoneNumber:     user.PhoneNumber,
	}
	return &api.QueryUserInfoResp{
		Data:      &u,
		Success:   true,
		ErrorCode: constant.SUCCESS_ERROR_CODE,
	}, nil
}

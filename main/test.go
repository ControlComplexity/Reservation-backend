package main

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段，如下
//iss (issuer)：签发人
//exp (expiration time)：过期时间
//sub (subject)：主题
//aud (audience)：受众
//nbf (Not Before)：生效时间
//iat (Issued At)：签发时间
//jti (JWT ID)：编号
// 如果想要保存更多信息，都可以添加到这个结构体中

type myClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

//定义jwt过期时间

const TokenExpireTime = time.Hour * 2

//定义Secret,这个密钥只有服务器才知道，不能泄露给用户

var MySecret interface{} = []byte("啦啦啦")

func main() {
	t, e := GetToken("2")
	if e != nil {
		log.Fatal("failed to get token: ", e.Error())
	}
	c, e := ParseToken(t)
	fmt.Println("t:", t, "c.Username: ", c)
}

//生成jwt

func GetToken(username string) (string, error) {
	//实例化结构体并初始化
	c := myClaims{
		username, //自定义字段
		jwt.StandardClaims{
			//设置过期时间
			ExpiresAt: time.Now().Add(TokenExpireTime).Unix(), //设置过期时间
			Issuer:    "my_project",                           //签发人
		},
	}
	//使用指定的签名方式创建签名对象
	/*jwt.NewWithClaims 方法根据 Claims 结构体创建 Token 示例。
	参数 1 是 jwt.SigningMethod，
	其中包含
	jwt.SigningMethodHS256，
	jwt.SigningMethodHS384，
	jwt.SigningMethodHS512
	三种 crypto.Hash 加密算法的方案。
	参数 2 是 Claims，包含自定义类型和 StandardClaim，StandardClaim 嵌入在自定义类型中，
	以方便对标准声明进行编码，解析和验证。
	*/
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	//使用指定的Secret签名并获得完整的编码后的字符串token
	//SignedString 方法根据传入的空接口类型参数 key，返回完整的签名令牌。
	return token.SignedString(MySecret.([]byte))
}

//解析jwt

func ParseToken(tokenString string) (*myClaims, error) {
	//解析token
	//jwt.ParseWithClaims 方法用于解析鉴权的声明，返回 *jwt.Token。
	token, err := jwt.ParseWithClaims(tokenString, &myClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		log.Fatal("failed to parse token: ", err)
		return nil, err
	}
	//Valid 方法用于校验鉴权的声明
	if claims, ok := token.Claims.(*myClaims); ok && token.Valid { //校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

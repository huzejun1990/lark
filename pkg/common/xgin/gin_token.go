// @Author huzejun 2024/2/28 21:06:00
package xgin

import (
	"github.com/golang-jwt/jwt/v4"
	uuid "github.com/satori/go.uuid"
	"time"
)

const jwt_ket = "lark_2023"

func CreateToken(uid int64, platfrom int8, access bool, ex int) (tk string, err error) {
	//token := jwt.New(jwt.SigningMethodES256)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["iss"] = "lark.con"
	now := time.Now()
	claims["exp"] = now.Add(time.Duration(ex) * time.Second)
	claims["iat"] = now.Unix()
	claims["session_id"] = uuid.NewV4().String()
	claims["uid"] = uid
	claims["platfrom"] = platfrom
	key := []byte(jwt_ket)
	tk, err = token.SignedString(key)
	return
}

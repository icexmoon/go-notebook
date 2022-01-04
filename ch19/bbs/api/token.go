package api

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/icexmoon/go-notebook/ch17/rest-xml/model"
)

const SECURITY_TOKEN = "stn"

type Token struct {
	Id      int       `json:"id"`
	Expire  time.Time `json:"expire"`
	SecCode string    `json:"scode"`
}

//生成安全码
func (t *Token) GetSecurityCode() (sc string, err error) {
	user := model.User{Id: t.Id}
	user.Get()
	sc = strconv.Itoa(t.Id) + user.Password + t.Expire.Format("2006-01-02 15:04:05") + SECURITY_TOKEN
	hash := md5.New()
	hash.Write([]byte(sc))
	sc = hex.EncodeToString(hash.Sum(nil))
	return
}

//生成访问令牌 string token
func (t *Token) String() (st string, err error) {
	t.SecCode, err = t.GetSecurityCode()
	if err != nil {
		return
	}
	jsonBytes, err := json.Marshal(t)
	if err != nil {
		return
	}
	st = base64.StdEncoding.EncodeToString(jsonBytes)
	return
}

//从string token解析生成Token
func (t *Token) Parse(token string) (err error) {
	bBytes, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return
	}
	err = json.Unmarshal(bBytes, t)
	return
}

//判断是否为有效token
func (t *Token) Validate() error {
	//过期检查
	if t.Expire.Before(time.Now()) {
		return errors.New("token is out of date, please login again")
	}
	//安全码检查
	sc, err := t.GetSecurityCode()
	if err != nil {
		fmt.Println(err)
		return err
	}
	if t.SecCode == sc {
		return nil
	}
	return errors.New("token is invalid, please login again")
}

//获取一个新的Token
func NewToken(id int) *Token {
	token := Token{Id: id}
	//有效期设置为1天
	token.Expire = time.Now().Add(24 * time.Hour)
	return &token
}

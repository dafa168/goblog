package util

import (
	"crypto/sha256"
	"encoding/hex"
	"goblog/models"
)

func init() {
	//log.Info("------------util init-----------")
}

/**
加密密码
 */
func EncodePass(pass string) string {
	sum := sha256.Sum256([]byte(pass));
	return hex.EncodeToString(sum[:]);
}

/**
验证密码
 */
func CheckPass(pass, m_pass string) bool {
	return EncodePass(pass) == m_pass
}

/**
检查登录
 */
func CheckLogin(user models.User) int64 {

	return 1;
}

func LoginSession(user models.User) bool {



	return false;
}

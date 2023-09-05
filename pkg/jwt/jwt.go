package jwt

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const jwtsecretKey = "CRpM4de2f0G6Vkvr"
const expiredHour = 24

// func generateJWT(secretKey string) (string, error) {
func EncodeJwt(tokenInfo map[string]any) (string, error) {
	return generateJWT(tokenInfo, jwtsecretKey)
}

func DecodeJwt(token string) (jwt.MapClaims, error) {
	return parseJWT(token, jwtsecretKey)
}

func Test() {
	// 密钥，用于签名和验证 JWT
	// secretKey := "your_secret_key"
	secretKey := jwtsecretKey
	// 创建一个 JWT
	token, err := generateJWTtest(secretKey)
	if err != nil {
		fmt.Println("生成 JWT 时发生错误：", err)
		return
	}
	fmt.Println("生成的 JWT：", token)

	// 验证和解析 JWT
	claims, err := parseJWT(token, secretKey)
	if err != nil {
		fmt.Println("解析 JWT 时发生错误：", err)
		return
	}
	fmt.Println("解析的 JWT 结果：", claims)
}

// 生成 JWT
func generateJWT(info map[string]any, secretKey string) (string, error) {
	claims := jwt.MapClaims{
		"info": info,
		"exp":  time.Now().Add(time.Hour * expiredHour).Unix(), // 过期时间为当前时间加24小时
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用密钥签名 JWT
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func generateJWTtest(secretKey string) (string, error) {
	claims := jwt.MapClaims{
		"username": "your_username",
		"exp":      time.Now().Add(time.Hour * expiredHour).Unix(), // 过期时间为当前时间加24小时
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用密钥签名 JWT
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// 解析 JWT
func parseJWT(tokenString, secretKey string) (jwt.MapClaims, error) {
	// 解析 JWT
	// fmt.Println(" parseJWT 111111", tokenString, secretKey)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 验证签名方法
		// fmt.Println(" parseJWT 111111", 222)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("无效的签名方法：%v", token.Header["alg"])
		}
		// fmt.Println(" parseJWT 111111", 3333, secretKey)
		// 返回密钥
		return []byte(secretKey), nil
	})
	if err != nil {
		// fmt.Println(" parseJWT 111111", 3344, err)
		return nil, err
	}
	// fmt.Println(" parseJWT 111111", 444)
	// 验证 JWT
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	// fmt.Println(" parseJWT 111111", 555)
	return nil, fmt.Errorf("无效的 JWT")
}

package auth

import (
	"log"
	pb "new/structs"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	signingkey = "visca barsa"
)

func GeneratedAccessJWTToken(req *pb.UserInfoRes) (string, error) {

	token := *jwt.New(jwt.SigningMethodHS256)

	//payload
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = req.UserID
	claims["role"] = req.Role
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(30 * time.Minute).Unix()

	newToken, err := token.SignedString([]byte(signingkey))
	if err != nil {
		log.Println(err)
		return "", err
	}

	return newToken, nil
}

func ValidateAccessToken(tokenStr string) (bool, error) {
	_, err := ExtractAccessClaim(tokenStr)
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExtractAccessClaim(tokenStr string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(signingkey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		return nil, err
	}

	return &claims, nil
}

func GetUserIdFromAccessToken(accessTokenString string) (*pb.UserInfoRes, error) {
	refreshToken, err := jwt.Parse(accessTokenString, func(token *jwt.Token) (interface{}, error) { return []byte(signingkey), nil })
	if err != nil || !refreshToken.Valid {
		return nil, err
	}
	claims, ok := refreshToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}
	user := pb.UserInfoRes{}
	user.UserID = claims["user_id"].(string)
	user.Role = claims["role"].(string)

	return &user, nil
}

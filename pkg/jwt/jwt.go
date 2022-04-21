package jwt

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

type AccessDetails struct {
	AccessUuid string
	UserId     string
}

type TokenDetails struct {
	AccessToken string
	AccessUuid  string
	Expires     int64
}

// CreateToken generates jwt token by userid string, returns *TokenDetails, error
func CreateToken(userid string) (*TokenDetails, error) {
	var err error
	td := &TokenDetails{}
	td.Expires = time.Now().Add(time.Minute * 15).Unix()
	td.AccessUuid = uuid.NewV4().String()

	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUuid
	atClaims["user_id"] = userid
	atClaims["exp"] = td.Expires
	atClaims["exp"] = time.Now().Add(time.Minute * 10).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("JWT_ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}
	return td, nil
}

// VerifyToken verifies token validity, returns *jwt.Token, error
func VerifyToken(t string) (*jwt.Token, error) {
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// ExtractTokenMetadata extracts access details from token, returns *AccessDetails, error
func ExtractTokenMetadata(t string) (*AccessDetails, error) {
	vt, err := VerifyToken(t)
	if err != nil {
		return nil, err
	}
	claims, ok := vt.Claims.(jwt.MapClaims)
	if ok && vt.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		userId := claims["user_id"].(string)
		if err != nil {
			return nil, err
		}
		return &AccessDetails{
			AccessUuid: accessUuid,
			UserId:     userId,
		}, nil
	}
	return nil, err
}

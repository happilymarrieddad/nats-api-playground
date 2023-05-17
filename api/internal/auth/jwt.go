package auth

import (
	"crypto/rsa"
	"errors"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	privKeyPath  = "/../../keys/app.rsa"     // openssl genrsa -out app.rsa keysize
	pubKeyPath   = "/../../keys/app.rsa.pub" // openssl rsa -in app.rsa -pubout > app.rsa.pub
	HOURS_IN_DAY = 24
	DAYS_IN_WEEK = 7
)

var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

func init() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	signBytes, err := ioutil.ReadFile(basepath + privKeyPath)
	if err != nil {
		panic(err)
	}
	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		panic(err)
	}
	verifyBytes, err := ioutil.ReadFile(basepath + pubKeyPath)
	if err != nil {
		panic(err)
	}
	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		panic(err)
	}
}

func CreateToken(keys map[string]interface{}) (string, error) {
	token := jwt.New(jwt.SigningMethodRS512)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * HOURS_IN_DAY * DAYS_IN_WEEK).Unix()
	claims["iat"] = time.Now().Unix()

	// Assign vals to claims
	for key, val := range keys {
		claims[key] = val
	}

	token.Claims = claims

	return token.SignedString(signKey)
}

func IsTokenValid(val string) (int64, error) {
	token, err := jwt.Parse(val, func(token *jwt.Token) (interface{}, error) {
		return verifyKey, nil
	})

	switch e := err.(type) {
	case nil:
		if !token.Valid {
			return 0, errors.New("token is invalid")
		}

		var user_id int64

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return 0, errors.New("token is invalid")
		}

		userID, exists := claims["id"]
		if exists {
			user_id = int64(userID.(float64))
		}

		return user_id, nil
	case *jwt.ValidationError:
		switch e.Errors {
		case jwt.ValidationErrorExpired:
			return 0, errors.New("token Expired, get a new one")
		default:
			return 0, errors.New("error while Parsing Token")
		}
	default:
		return 0, errors.New("unable to parse token")
	}
}

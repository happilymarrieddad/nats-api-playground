package auth

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/happilymarrieddad/nats-api-playground/api/types"
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

func GetUserFromToken(val string) (usr *types.User, err error) {
	token, err := jwt.Parse(val, func(token *jwt.Token) (interface{}, error) {
		return verifyKey, nil
	})

	switch err.(type) {
	case nil:
		if !token.Valid {
			return nil, errors.New("token is invalid")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return nil, errors.New("token is invalid")
		}

		rawUsr, exists := claims["user"]
		if !exists {
			return nil, errors.New("user does not exist on token 1")
		}

		usrMp, ok := rawUsr.(map[string]interface{})
		if !ok {
			return nil, errors.New("user does not exist on token 2")
		}

		return types.GetUserFromMap(usrMp), nil
	default:
		return nil, errors.New("invalid token")
	}
}

func IsTokenValid(val string) error {
	token, err := jwt.Parse(val, func(token *jwt.Token) (interface{}, error) {
		return verifyKey, nil
	})

	switch e := err.(type) {
	case nil:
		if !token.Valid {
			return errors.New("token is invalid")
		}

		_, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return errors.New("token is invalid")
		}

		return nil
	case *jwt.ValidationError:
		switch e.Errors {
		case jwt.ValidationErrorExpired:
			return errors.New("token Expired, get a new one")
		default:
			return fmt.Errorf("error while Parsing Token err: %s", e.Error())
		}
	default:
		return errors.New("unable to parse token")
	}
}

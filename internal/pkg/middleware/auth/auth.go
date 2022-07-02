package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/golang-jwt/jwt/v4"
	"strings"
	"time"
)

const (
	RealWorldToken = "Token"
)

func GenerateToken(secret, username string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"nbf":      time.Date(2022, 6, 26, 22, 42, 10, 0, time.UTC).Unix(),
	})
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		//panic(err)
		fmt.Println("err: ", err)
	}
	return tokenString
}

func JWTAuth(secret string) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				tokenString := tr.RequestHeader().Get("Authorization")
				fmt.Println("tokenString: ", tokenString)
				auths := strings.SplitN(tokenString, " ", 2)
				fmt.Println("auths: ", auths)
				if len(auths) != 2 || !strings.EqualFold(auths[0], RealWorldToken) {
					return nil, errors.New("JWT Token missing")
				}
				spew.Dump("token: ", tokenString)
				// Parse takes the token string and a function for looking up the key. The latter is especially
				// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
				// head of the token to identify which key to use, but the parsed token (head and claims) is provided
				// to the callback, providing flexibility.
				token, err := jwt.Parse(auths[1], func(token *jwt.Token) (interface{}, error) {
					// Don't forget to validate the alg is what you expect:
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
					}

					// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
					return []byte(secret), nil
				})
				if err != nil {
					return nil, err
				}

				if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
					spew.Dump(claims["username"])
				} else {
					return nil, errors.New("token Invaild")
				}
			}
			return handler(ctx, req)
		}
	}
}

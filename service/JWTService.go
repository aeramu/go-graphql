package service

import(
  "github.com/dgrijalva/jwt-go"
)

var JWTSecretKey = []byte("INIB4Ru53kretKeY")

type JWTClaims struct{
  jwt.StandardClaims
  Payload string
}

func CreateJWT(payload string)(string){
  jwtClaims := &JWTClaims{
    Payload: payload,
  }
  token,_ := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims).SignedString(JWTSecretKey)

  return "token="+token
}

func DecodeJWT(token string)(string){
  token = token[6:]
  claims := new(JWTClaims)
  jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return token, nil
	})
  return claims.Payload
}

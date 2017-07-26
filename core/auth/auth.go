package auth

import (
	"github.com/dgrijalva/jwt-go"
)

func CreateTokenJWT(_data map[string]interface{}, _secretKey interface{}) ([]byte) {

	/* Create the token */
	token := jwt.New(jwt.SigningMethodHS256)

	/* Create a map to store our claims */
	claims := token.Claims.(jwt.MapClaims)

	/* Set token claims */
	for _key, _value := range _data {
		claims[_key] = _value
	}
	//claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	/* Sign the token with our secret */
	tokenString, _ := token.SignedString(_secretKey)

	/* Finally, write the token to the browser window */
	return []byte(tokenString)
}

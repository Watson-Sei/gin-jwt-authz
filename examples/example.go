package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"

	jwtAuthz "github.com/Watson-Sei/gin-jwt-authz"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gin-gonic/gin"
)

type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}

type JSONWebKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

func main() {
	router := gin.Default()

	// Only requesting users who have all the specified privileges can pass through.
	router.GET("/sample1", CheckJWT(), jwtAuthz.CheckPermissions([]string{"create:books", "update:books", "delete:books"}, jwtAuthz.DefaultOptions()), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "It has create:books, update:books and delete:books permissions.",
		})
	})

	// The conditions are the same as above.
	router.GET("/sample2", CheckJWT(), jwtAuthz.CheckPermissions([]string{"create:books"}, jwtAuthz.DefaultOptions()), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "It has create:books permission.",
		})
	})

	router.Run()
}

var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		aud := os.Getenv("AUTH0_IDENTIFIER")
		convAud, ok := token.Claims.(jwt.MapClaims)["aud"].([]interface{})
		if !ok {
			strAud, ok := token.Claims.(jwt.MapClaims)["aud"].(string)
			if !ok {
				return token, errors.New("Invalid audience.")
			}
			if strAud != aud {
				return token, errors.New("Invalid audience.")
			}
		} else {
			for _, v := range convAud {
				if v == aud {
					break
				} else {
					return token, errors.New("Invalid audience.")
				}
			}
		}
		iss := "https://" + os.Getenv("AUTH0_DOMAIN") + "/"
		checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
		if !checkIss {
			return token, errors.New("Invalid issuer.")
		}

		cert, err := getPemCert(token)
		if err != nil {
			panic(err.Error())
		}

		result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
		return result, nil
	},
})

func CheckJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtMid := *jwtMiddleware
		if err := jwtMid.CheckJWT(c.Writer, c.Request); err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

func getPemCert(token *jwt.Token) (string, error) {
	cert := ""
	resp, err := http.Get("https://" + os.Getenv("AUTH0_DOMAIN") + "/.well-known/jwks.json")

	if err != nil {
		return cert, err
	}
	defer resp.Body.Close()

	var jwks = Jwks{}
	err = json.NewDecoder(resp.Body).Decode(&jwks)

	if err != nil {
		return cert, err
	}

	for k, _ := range jwks.Keys {
		if token.Header["kid"] == jwks.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + jwks.Keys[k].X5c[0] + "\n-----END CERTIFICATE-----"
		}
	}

	if cert == "" {
		err := errors.New("Unable to find appropriate key.")
		return cert, err
	}

	return cert, nil
}

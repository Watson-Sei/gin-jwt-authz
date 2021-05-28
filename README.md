# gin-jwt-authz
Validate a JWTs `scope` to authorize access to an endpoint. gin-jwt-authz is a Gin middleware/handler inspired by [express-jwt-authz](https://github.com/auth0/express-jwt-authz)

## Install 
```
$ go get -u github.com/Watson-Sei/gin-jwt-authz
```

## Usage
Use it with [go-jwt-middleware](https://github.com/auth0/go-jwt-middleware) to validate the JWT and make sure you have the proper permissions to call the endpoints.

[Sample code is here.](https://github.com/Watson-Sei/gin-jwt-authz/tree/main/examples)

## Options
- `failWithError`:  When set to true, will forward errors to next instead of ending the response directly. Defaults to true.
- `checkAllScopes`: When set to true, all the expected scopes will be checked against the user's scopes. Defaults to true.
# gin-jwt-authz Example

## Configuration
### Environment Variables

Copy ".env.example" and set your Auth0 environment variable.

```
AUTH0_DOMAIN=YOUR_AUTH0_DOMAIN
AUTH0_IDENTIFIER=YOUR_AUTH0_AUDIENCE
```

### Export Enviroment variable
```
$ export $(cat .env | grep -v ^# | xargs)
```
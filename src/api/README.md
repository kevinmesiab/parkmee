# Parmkmee.com API
## Official API documentation

## Healtheck

`curl http://localhost:8000/healtheck`

The response from the healtheck should be a `200` response code.  The body of the healtheck body should return "👍🏻"

### Base API URL
`http://localhost:8000/api/v1`

## Users

`curl -X GET http://localhost:8000/api/v1/users/{id}`

```
{
  "id" : 0,
  "firstname": "name",
  "lastname" : "lastname",
  "email" : "user@email.tld"
}
```
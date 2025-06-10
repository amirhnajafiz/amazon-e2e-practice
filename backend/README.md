# Backend

The backend service of this project is an authentication API which provides two endpoints for signin and signup.

## Endpoints

Here is a table of API endpoints.

| Method | Endpoint          | Description                      | Auth Required |
|--------|-------------------|----------------------------------|---------------|
| PUT    | `/api/signup`     | Register a new user              | No            |
| POST   | `/api/signin`     | Authenticate user and get token  | No            |
| GET    | `/api`            | Validate the token               | Yes           |

### Examples

```sh
❯ curl localhost:8080/health
OK

❯ curl -X PUT -H "Content-Type: application/json" -d '{"username":"admin", "password":"12345678"}'  "localhost:8080/api"
{"message":"User created successfully"}

❯ curl -X POST -H "Content-Type: application/json" -d '{"username":"admin", "password":"12345678"}'  "localhost:8080/api"
{"message":"Signin successful","token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3NDk1ODg1MDIsInVzZXJuYW1lIjoiYWRtaW4ifQ.GjkWvSHzzKcmWhHUSBHT1nRlYC0jQ_bWDLY-T926UlI"}

❯ curl -H "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3NDk1ODg1MDIsInVzZXJuYW1lIjoiYWRtaW4ifQ.GjkWvSHzzKcmWhHUSBHT1nRlYC0jQ_bWDLY-T926UlI" "localhost:8080/api"
{"message":"Token is valid","username":"admin"}
```

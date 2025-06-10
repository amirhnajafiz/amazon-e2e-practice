# Backend

The backend service of this project is an authentication API which provides two endpoints for signin and signup.

## Endpoints

Here is a table of API endpoints.

| Method | Endpoint          | Description                      | Auth Required |
|--------|-------------------|----------------------------------|---------------|
| PUT    | `/api/signup`     | Register a new user              | No            |
| POST   | `/api/signin`     | Authenticate user and get token  | No            |
| GET    | `/api`            | Validate the token               | Yes           |

# API

This system is the backend service of our practice. It is responsible for managing panel URLs, users sessions, and history stats. The API provides JSON endpoints that interact with the frontend application.

## Endpoints

| Path                | Method                         |
|:-------------------:|:------------------------------:|
| `/api/urls`         | `GET`                          |
| `/api/urls/:id`     | `PUT`                          |
| `/api/stats`        | `GET`                          |

## cURLs

```
curl -X POST -H "Content-Type: application/json" -H "X-Admin-Key: admin" -d '{"shortened_id": "Go", "address": "https://go.com", "description": "Golang website"}'  "localhost:8080/admin/urls"

curl -X DELETE -H "Content-Type: application/json" -H "X-Admin-Key: admin"  "localhost:8080/admin/urls/26"
```
# API

This system is the backend service of our practice. It is responsible for managing panel URLs, users sessions, and history stats. The API provides JSON endpoints that interact with the frontend application.

## Endpoints

| Path                | Method                         |
|:-------------------:|:------------------------------:|
| `/api/urls`         | `GET`                          |
| `/api/urls/:id`     | `PUT`                          |
| `/api/stats`        | `GET`                          |

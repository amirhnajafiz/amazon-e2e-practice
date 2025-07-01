# API Overview

This API serves as the backend for the Amazon End-to-End practice project, managing the URLs used in our unified panel application. It offers JSON endpoints to retrieve and update URL data, view statistics, and interact with a PostgreSQL database. The API is built in Golang and includes both public and private (admin-only) endpoints for full URL management.

## Available Endpoints

The following table lists the available API endpoints:

| Path                 | Method   | Description                     | Public |
| -------------------- | -------- | ------------------------------- | ------ |
| `/api/urls`          | `GET`    | Retrieve all stored URLs        | Yes    |
| `/api/urls/:id`      | `PUT`    | Update visit stats for a URL    | Yes    |
| `/api/stats?limit=N` | `GET`    | Fetch top N URLs by visit count | Yes    |
| `/admin/urls`        | `POST`   | Create a new URL                | No     |
| `/admin/urls/:id`    | `DELETE` | Delete a URL by ID              | No     |

> **Note:** Admin-only endpoints require an `X-Admin-Key` header. This key is configured during API setup.

## Example Requests (Using `curl`)

Here are two sample `curl` commands for interacting with the admin endpoints:

**Create a new URL:**

```sh
curl -X POST -H "Content-Type: application/json" \
     -H "X-Admin-Key: admin" \
     -d '{"shortened_id": "Go", "address": "https://go.com", "description": "Golang website"}' \
     http://localhost:8080/admin/urls
```

**Delete a URL by ID:**

```sh
curl -X DELETE -H "Content-Type: application/json" \
     -H "X-Admin-Key: admin" \
     http://localhost:8080/admin/urls/26
```

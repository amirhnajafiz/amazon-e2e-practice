# Amazon End-to-End Practice

A key skill for any system tech professional is the ability to develop, deploy, and manage applications. In this practice, I built a full-stack application consisting of:

* A **unified panel frontend** (Vue.js + Vite) that allows users to access various URLs.
* A **backend API** written in Go, which interacts with a PostgreSQL database to manage and track these URLs.

To complete the end-to-end workflow, I containerized the system and deployed it using three different methods:

1. **Standard Docker commands**: In `scripts` directory, the `local_build.sh` and `local_deployment.sh`.
2. **Docker Compose with Nginx as a reverse proxy**: Using `docker compose up -d`.
3. **Infrastructure automation using Ansible and Terraform**

This project demonstrates the full lifecycle of application development and deployment—from writing code to automating infrastructure.

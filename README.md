# Amazon E2E Practice

Welcome to the **Amazon E2E Practice** repository! This project showcases a complete workflow for building, deploying, and automating modern applications.

## 🚀 Overview

| Layer            | Stack                                                                 | Purpose                                              |
|------------------|----------------------------------------------------------------------|------------------------------------------------------|
| Frontend         | ![Vue.js](https://img.shields.io/badge/Vue.js-35495E?logo=vue.js&logoColor=4FC08D) ![Vite](https://img.shields.io/badge/Vite-646CFF?logo=vite&logoColor=FFD62E) | User interface for managing URLs                     |
| Backend          | ![Go](https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=white) | REST API and business logic                          |
| Database         | ![PostgreSQL](https://img.shields.io/badge/PostgreSQL-4169E1?logo=postgresql&logoColor=white) | Persistent data storage                              |
| Containerization | ![Docker](https://img.shields.io/badge/Docker-2496ED?logo=docker&logoColor=white) | Consistent packaging and deployment                  |
| Orchestration    | ![Docker Compose](https://img.shields.io/badge/Docker%20Compose-000000?logo=docker&logoColor=2496ED) | Multi-container management with Nginx reverse proxy  |
| IaC              | ![Terraform](https://img.shields.io/badge/Terraform-7B42BC?logo=terraform&logoColor=white) | Automated infrastructure provisioning                |

## ✨ Features

- **Unified Frontend**: Modern, responsive panel for URL management.
- **High-Performance Backend**: Go-based API with PostgreSQL storage.
- **Containerized Workflows**: Docker ensures environment consistency.
- **Flexible Deployments**:
    - Local Docker scripts:  
        - `scripts/local_build.sh` – Build images  
        - `scripts/local_deployment.sh` – Deploy locally
    - Docker Compose:  
        - `docker compose up -d` – Orchestrate services with Nginx
    - Terraform:  
        - `scripts/terraform_deployment.sh` – Provision infrastructure  
        - `scripts/terraform_teardown.sh` – Teardown resources  
        - Configurations in `terraform/`

## ⚙️ Workflow

1. **Develop**:  
     Build frontend (Vue.js + Vite) and backend (Go) components.

2. **Build & Test**:  
     Use Docker for local builds and testing.

3. **Deploy**:  
     - Local scripts for quick testing
     - Docker Compose for orchestrated environments
     - Terraform for cloud automation

4. **Automate**:  
     Streamline deployment, scaling, and teardown with scripts and IaC.

## 📂 Structure

| Path                  | Description                                 |
|-----------------------|---------------------------------------------|
| `frontend/`           | Frontend source (Vue.js + Vite)             |
| `api/`                | Backend API (Go)                            |
| `build/`              | Dockerfiles                                 |
| `configs/`            | Configuration files (API, NginX)            |
| `scripts/`            | Automation scripts                          |
| `terraform/`          | Terraform configs                           |
| `docker-compose.yml`  | Service orchestration                       |
| `README.md`           | Project documentation                       |

## 💡 Key Concepts

- End-to-end full-stack development
- Best practices in containerization
- Multiple deployment strategies
- Infrastructure automation with Terraform

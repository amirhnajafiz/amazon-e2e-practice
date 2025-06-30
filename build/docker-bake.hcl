group "default" {
  targets = ["api", "frontend"]
}

target "api" {
  context    = "api"
  dockerfile = "../build/Dockerfile.api"
  tags       = ["amazon-e2e-practice-api:latest"]
}

target "frontend" {
  context    = "frontend"
  dockerfile = "../build/Dockerfile.frontend"
  tags       = ["amazon-e2e-practice-frontend:latest"]
}

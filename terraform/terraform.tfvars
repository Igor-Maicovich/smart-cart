app_image   = "ghcr.io/igor-maicovich/smart-cart:latest"
db_image    = "postgres:15-alpine"
redis_image = "redis:7-alpine"

app_port   = 8080
db_port    = 5432
redis_port = 6379

replicas = 1
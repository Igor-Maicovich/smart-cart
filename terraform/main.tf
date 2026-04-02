resource "docker_network" "backend" {
  name   = "backend"
  driver = "bridge"
}

resource "docker_volume" "postgres_data" {
  name = "postgres_data"
}

resource "docker_container" "db" {
  name    = "smart-cart-db"
  image   = var.db_image
  restart = "always"
  env = [
    "POSTGRES_DB=smartcart",
    "POSTGRES_USER=user",
    "POSTGRES_PASSWORD=password"
  ]
  volumes {
    volume_name    = docker_volume.postgres_data.name
    container_path = "/var/lib/postgresql/data"
  }
  networks_advanced {
    name = docker_network.backend.name
  }
}

resource "docker_container" "cache" {
  name    = "smart-cart-redis"
  image   = var.redis_image
  restart = "always"
  networks_advanced {
    name = docker_network.backend.name
  }
}

resource "docker_container" "app" {
  count   = var.replicas
  name    = "smart-cart-app-${count.index}"
  image   = var.app_image
  restart = "always"
  env = [
    "DB_HOST=${docker_container.db.name}",
    "DB_PORT=${var.db_port}",
    "DB_NAME=smartcart",
    "DB_USER=user",
    "DB_PASSWORD=password",
    "REDIS_HOST=${docker_container.cache.name}",
    "REDIS_PORT=${var.redis_port}"
  ]
  ports {
    internal = 8080
    external = var.app_port + count.index
  }
  networks_advanced {
    name = docker_network.backend.name
  }
  depends_on = [
    docker_container.db,
    docker_container.cache
  ]
}
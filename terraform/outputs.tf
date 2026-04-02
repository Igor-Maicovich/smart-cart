output "app_container_names" {
  description = "Names of app containers"
  value = docker_container.app[*].name
}

output "app_container_ids" {
  description = "IDs of app containers"
  value = docker_container.app[*].id
}

output "app_external_ports" {
  description = "External ports for app containers"
  value = docker_container.app[*].ports[0].external
}

output "app_urls" {
  description = "URLs to access application"
  value = [
    for port in docker_container.app[*].ports[0].external :
    "http://localhost:${port}"
  ]
}

output "db_container_name" {
  description = "Database container name"
  value = docker_container.db.name
}

output "db_container_id" {
  description = "Database container ID"
  value = docker_container.db.id
}

output "db_volume_name" {
  description = "Volume used by database"
  value = docker_volume.postgres_data.name
}

output "db_connection_info" {
  description = "Internal DB connection info"
  value = {
    host = docker_container.db.name
    port = var.db_port
  }
}

output "cache_container_name" {
  description = "Redis container name"
  value = docker_container.cache.name
}

output "cache_container_id" {
  description = "Redis container ID"
  value = docker_container.cache.id
}

output "cache_connection_info" {
  description = "Internal Redis connection info"
  value = {
    host = docker_container.cache.name
    port = var.redis_port
  }
}

output "network_name" {
  description = "Docker network name"
  value = docker_network.backend.name
}

output "infrastructure_summary" {
  description = "Full infrastructure overview"
  value = {
    app = [
      for c in docker_container.app : {
        name = c.name
        port = c.ports[0].external
      }
    ]
    db = {
      name = docker_container.db.name
      volume = docker_volume.postgres_data.name
    }
    cache = {
      name = docker_container.cache.name
    }
    network = docker_network.backend.name
  }
}
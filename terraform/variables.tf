variable "app_image" {
  description = "App image"
  type        = string
}

variable "db_image" {
  type = string
}

variable "redis_image" {
  type = string
}

variable "app_port" {
  type = number
}

variable "db_port" {
  type = number
}

variable "redis_port" {
  type = number
}

variable "replicas" {
  type = number
}
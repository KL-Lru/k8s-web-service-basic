variable "project_id" {
  type = string
}

variable "sql_user_password" {
  type = string
}

variable "region" {
  type    = string
  default = "asia-northeast1"
}

variable "location" {
  type    = string
  default = "asia-northeast1-a"
}

variable "machine_type" {
  type    = string
  default = "e2-medium"
}

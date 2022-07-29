resource "google_sql_database_instance" "sample" {
  name                = "primary0"
  database_version    = "POSTGRES_14"
  region              = var.region
  deletion_protection = false
  settings {
    tier              = "db-f1-micro"
    activation_policy = "ALWAYS"
    backup_configuration {
      enabled = true
    }
  }
}

resource "google_sql_database" "database" {
  name      = "sample-db"
  instance  = google_sql_database_instance.sample.name
  charset   = "UTF8"
  collation = "en_US.UTF8"
}

### WARNING ### sensitive data "password" in tfstate as plain text.
###              Please use encryptable backend or not manage by terraform.
resource "google_sql_user" "users" {
  instance = google_sql_database_instance.sample.name
  name     = "proxy-user"
  password = var.sql_user_password
  type     = "BUILT_IN"
}

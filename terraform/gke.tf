resource "google_container_cluster" "primary" {
  project                  = var.project_id
  name                     = "sample"
  location                 = var.location
  initial_node_count       = 1
  remove_default_node_pool = true
  network                  = google_compute_network.gke_sample.name
  subnetwork               = google_compute_subnetwork.sample_subnet.name
  networking_mode          = "VPC_NATIVE"
  ip_allocation_policy {
    cluster_secondary_range_name  = "pod-range"
    services_secondary_range_name = "service-range"
  }
  release_channel {
    channel = "UNSPECIFIED"
  }
  min_master_version    = "1.24.1-gke.1800"
  enable_shielded_nodes = true
  addons_config {
    http_load_balancing {
      disabled = false
    }
  }
}

resource "google_container_node_pool" "primary_node_pool" {
  project    = var.project_id
  name       = "sample-nodes"
  node_count = 1
  location   = var.location
  cluster    = google_container_cluster.primary.name

  management {
    auto_repair  = true
    auto_upgrade = false
  }

  node_config {
    disk_size_gb = 15
    machine_type = var.machine_type
    oauth_scopes = [
      "https://www.googleapis.com/auth/cloud-platform",
      "https://www.googleapis.com/auth/logging.write",
      "https://www.googleapis.com/auth/monitoring",
    ]
    shielded_instance_config {
      enable_integrity_monitoring = true
      enable_secure_boot          = true
    }
  }
  depends_on = [
    google_compute_network.gke_sample,
    google_compute_subnetwork.sample_subnet
  ]
}

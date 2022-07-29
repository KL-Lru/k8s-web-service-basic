resource "google_compute_network" "gke_sample" {
  project                 = var.project_id
  name                    = "gke-sample"
  auto_create_subnetworks = false
}


resource "google_compute_subnetwork" "sample_subnet" {
  name          = "sample-subnet"
  ip_cidr_range = "10.0.0.0/16"
  secondary_ip_range = [
    {
      range_name    = "core-range",
      ip_cidr_range = "10.1.0.0/16"
    },
    {
      range_name    = "pod-range",
      ip_cidr_range = "10.2.0.0/16"
    },
    {
      range_name    = "service-range",
      ip_cidr_range = "10.3.0.0/16"
    },
  ]
  region  = var.region
  network = google_compute_network.gke_sample.id
}

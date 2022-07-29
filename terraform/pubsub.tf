resource "google_pubsub_topic" "sample" {
  name = "SAMPLE"
}

resource "google_pubsub_subscription" "sample" {
  name                       = "sample"
  topic                      = google_pubsub_topic.sample.name
  ack_deadline_seconds       = 10
  message_retention_duration = "604800s"
  retain_acked_messages      = false
}

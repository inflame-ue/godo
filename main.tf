terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "6.8.0"
    }
  }
}

provider "google" {
  credentials = file(var.credentials_file)
  project     = var.project_id
  region      = var.region
  zone        = var.zone
}

resource "google_compute_instance" "vm_instance" {
  name         = "todo-vm"
  machine_type = "e2-micro"
  zone         = "us-central1-c"

  boot_disk {
    initialize_params {
      image = "ubuntu-os-cloud/ubuntu-2404-lts-amd64"
    }
  }

  network_interface {
    network = "default"
    access_config {}
  }

  metadata = {
    ssh-keys = "ubuntu:${file(var.ssh_public_key_file)}"
  }

  tags = ["web"]
}

resource "google_compute_firewall" "default" {
  name    = "todo-firewall"
  network = "default"

  allow {
    protocol = "tcp"
    ports    = ["80", "8080"]
  }

  source_ranges = ["0.0.0.0/0"]
}

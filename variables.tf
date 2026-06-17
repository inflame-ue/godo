variable "project_id" {
  description = "GCP Project ID"
  default     = "todo-499405"
  type        = string
}

variable "region" {
  description = "GCP Region"
  type        = string
  default     = "us-central1"
}

variable "credentials_file" {
  description = "Path to service account JSON Key File"
  default     = "todo-499405-240a4fa698ae.json"
  type        = string
}

variable "ssh_public_key_file" {
  description = "Path to ssh .pub file"
  default     = "/home/inflame-uwu/.ssh/godo.pub"
  type        = string
}

variable "zone" {
  description = "GCP Zone Region"
  default     = "us-central1-c"
  type        = string
}

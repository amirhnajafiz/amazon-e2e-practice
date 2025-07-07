# Terraform variables definitions

variable "vm_ip" {
  description = "IP address of the VM"
  type        = string
}

variable "vm_user" {
  description = "SSH username"
  type        = string
  default     = "ubuntu"
}

variable "ssh_private_key_path" {
  description = "Path to your private SSH key"
  type        = string
  default     = "~/.ssh/id_rsa"
}

variable "git_repo_url" {
    description = "URL of the Git repository to clone"
    type        = string
    default     = ""
}

variable "nginx_conf_path" {
  description = "Path to the nginx configuration file"
  type        = string
  default     = "${path.module}/nginx.conf"
}

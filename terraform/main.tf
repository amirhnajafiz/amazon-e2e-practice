provider "null" {}

resource "null_resource" "app_setup" {
  connection {
    type        = "ssh"
    host        = var.vm_ip
    user        = var.vm_user
    private_key = file(var.ssh_private_key_path)
  }

  provisioner "remote-exec" {
    inline = [
      "sudo apt update",
      "sudo apt install -y git docker.io docker-compose",
      "sudo usermod -aG docker ${var.vm_user}",
      "git clone ${var.git_repo_url} /home/${var.vm_user}/app || true",
      "cd /home/${var.vm_user}/app && sudo docker-compose up -d"
    ]
  }

  # Optional cleanup provisioner
  provisioner "remote-exec" {
    when    = destroy
    inline  = [
      "cd /home/${var.vm_user}/app && sudo docker-compose down",
      "sudo rm -rf /home/${var.vm_user}/app"
    ]
  }
}

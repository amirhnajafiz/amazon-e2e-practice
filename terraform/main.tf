# Main Terraform configuration file
provider "null" {}

resource "null_resource" "app_setup" {
  # Open SSH connection to the VM
  connection {
    type        = "ssh"
    host        = var.vm_ip
    user        = var.vm_user
    private_key = file(var.ssh_private_key_path)
  }

  # Upload nginx.conf to the VM
  provisioner "file" {
    source      = "${var.nginx_conf_path}"
    destination = "/home/${var.vm_user}/nginx.conf"
  }

  # Install packages
  provisioner "apt" {
    update = true
    packages = [
      "git",
      "docker.io",
      "docker-compose",
      "nginx",
      "openssl"
    ]
  }

  # Set up Docker and Nginx
  provisioner "remote-exec" {
    inline = [
      "sudo usermod -aG docker ${var.vm_user}",
      "git clone ${var.git_repo_url} /home/${var.vm_user}/app || true",
      "cd /home/${var.vm_user}/app && sudo docker-compose up -d",
      # Move nginx.conf to the correct location
      "sudo mv /home/${var.vm_user}/nginx.conf /etc/nginx/sites-available/app.conf",
      "sudo ln -sf /etc/nginx/sites-available/app.conf /etc/nginx/sites-enabled/app.conf",
      # Generate SSL keys
      "sudo openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout /etc/ssl/private/example.com.key -out /etc/ssl/certs/example.com.crt -subj '/CN=example.com'",
      # Restart Nginx
      "sudo systemctl restart nginx"
    ]
  }

  # Clean up on destroy
  provisioner "remote-exec" {
    when    = destroy
    inline  = [
      "cd /home/${var.vm_user}/app && sudo docker-compose down",
      "sudo rm -rf /home/${var.vm_user}/app",
      "sudo rm -f /etc/nginx/sites-available/app.conf /etc/nginx/sites-enabled/app.conf",
      "sudo rm -f /etc/ssl/certs/example.com.crt /etc/ssl/private/example.com.key"
    ]
  }
}

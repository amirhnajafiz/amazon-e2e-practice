# Terraform variables

vm_ip              = "192.168.1.10"
vm_user            = "ubuntu"
ssh_private_key_path = "~/.ssh/id_rsa"
git_repo_url       = "https://github.com/amirhnajafiz/amazon-e2e-practice.git"
nginx_conf_path    = "${path.module}/../configs/nginx.conf"

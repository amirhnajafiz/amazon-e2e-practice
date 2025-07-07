#!/bin/bash

cd terraform

terraform init
terraform apply -auto-approve
if [ $? -ne 0 ]; then
    echo "Terraform deployment failed."
    exit 1
fi

echo "Terraform deployment completed successfully."

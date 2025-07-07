#!/bin/bash

cd terraform

terraform destroy -auto-approve
if [ $? -ne 0 ]; then
    echo "Terraform teardown failed."
    exit 1
fi

echo "Terraform teardown completed successfully."

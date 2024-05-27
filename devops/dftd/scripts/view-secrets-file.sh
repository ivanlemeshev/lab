#!/usr/bin/bash

set -e

ansible-vault view --vault-password-file="./ansible/vault_password.txt" ./ansible/secrets.yml

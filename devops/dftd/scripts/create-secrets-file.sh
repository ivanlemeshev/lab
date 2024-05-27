#!/usr/bin/bash

set -e

echo "Generating user's password"

vault_password="./ansible/vault_password.txt"
password=$(pwgen --secure --capitalize --numerals --symbols 12 1)

echo "password: $password" > ./ansible/secrets.yml
echo "password_hash: $(echo "${password}" | mkpasswd --stdin --method=sha-512)" >> ./ansible/secrets.yml

ansible-vault encrypt --vault-password-file="${vault_password}" ./ansible/secrets.yml

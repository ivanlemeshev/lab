#!/usr/bin/bash

set -e

echo "Generating vault password file"

pwgen --secure --capitalize --numerals --symbols 64 1 > ./ansible/vault_password.txt

echo "The new vault password is saved in ./ansible/vault_password.txt"

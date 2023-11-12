#!/usr/bin/env bash

# Generate root CA
openssl genrsa -out ca.key 4096
openssl req -x509 -new -nodes -key ca.key -sha256 -out ca.crt -subj "/C=US/ST=State/L=City/O=Organization/CN=MyRootCA"

# Generate server side cert and key
openssl genrsa -out server.key 2048
openssl req -new -key server.key -out server.csr -subj "/C=US/ST=State/L=City/O=Organization/CN=localhost"
openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 365 -sha256
mv server.crt key/
mv server.key key/
mv server.csr key/

# Generate client side cert and key
## Generate client key
openssl genrsa -out client.key 2048
## Create a certificate signing request (CSR) for the client
openssl req -new -key client.key -out client.csr -subj "/C=US/ST=State/L=City/O=Organization/CN=localhost"
## Sign the client certificate with the root CA key and certificate
openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out client.crt -days 365 -sha256
mv client.crt key/
mv client.key key/
mv client.csr key/

mv ca.key key/
mv ca.crt key/
mv ca.csr key/

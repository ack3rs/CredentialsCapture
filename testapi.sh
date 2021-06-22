#!/bin/sh

curl  --data '{"firstname":"Mark","lastname":"Ackroyd","country": "UK", "email": "mark@ackroyd.net"}' "http://localhost:8080/save"
echo " - OK"

curl  --data '{"firstnam":"Mark","lastnam":"Ackroyd","ountry": "UK", "email": "mark@ackroyd.net"}' "http://localhost:8080/save"
echo " - Country is Missing"

curl  --data '{' "http://localhost:8080/save"
echo " - Malformed"

curl  --data '{"firstnam":"Mark","lastname":"Ackroyd","country": "UK", "email": "mark@ackroyd.net"}' "http://localhost:8080/save"
echo " - Firstname is Missing"

curl  --data '{"firstname":"Mark","lastname":"Ackroyd","country": "UK", "email": "NOT AN EMAIL"}' "http://localhost:8080/save"
echo " - Email is Invalid"
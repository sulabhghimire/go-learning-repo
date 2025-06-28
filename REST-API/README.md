How to Create a Test SSL Certificate

openssl req -x509 -newkey rsa:2048 -nodes -keyout key.pem -out cert.pem -days 365
After filling the info asked created two files
-> cert.pem
-> key.pem

.pem extends for Privacy Enhanced Mail
It is a base64 encoded DER certificates and used to storing and transferring cryptographic keys, certificates and other data.
DER means Distinguished Encoding Rules

Key Characteristics Of .pem file

- The pem format is base64 encoded making it ASCII text which is easier to read and transport in text based protocols.
- pem file has specific header and footer to identify the type of content they hold

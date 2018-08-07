# go-https-client-server

Example of Server and Client that communicate over HTTPS.

They're currently handle only `/telemetry` endpoint, but using the example you can add another ones. Client sends some (fixed) data, server accept and log it, then response to the client with another (fixed) message.

Client features:
* DNS Name resolving using local resolver first, then bypass onto specified remote DNS.
* Granular timeouts on different operations, including TLS Handshake and zombie/idle connections.
* TCP Keep-Alive is enabled by "net/http" by default, but some related timeouts are set as well.

## Installation
_Server:_
1. ```go get github.com/niki4/go-https-client-server```
2. ```cd go-https-client-server```
3. You will need to obtain security certificate files (cert.pem, key.pem). If you have OpenSSL installed in your system, then you could use following command to generate self-signed certificate:
`openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout key.pem -out cert.pem`   
4. You can enter any data on certificate fields, but Common Name - it must be your server name. For localhost it will be: 
`Common Name (e.g. server FQDN or YOUR name) []:localhost`
5. ```go run server.go```

_Client:_

Follow the same steps as for the server, but run `client.go`. It's best to have `server.go` running on the separate machine:
1. Open `client.go` with some text editor and specify your server name in "baseURL" const, so it would like `baseURL = "nameyourserver.com"`. Save and close. 
2. ```go run client.go```


##Known issues
1. There's an issue with resolving `localhost` to 127.0.0.1 because of DNS returns `NXDOMAIN` flag instead of IP. As a bypass, for local testing you may keep `baseURL = "127.0.0.1"` client config.
2. Client has config `TLSClientConfig: &tls.Config{InsecureSkipVerify: true}` to allow accepting unsigned security certificates. As it's security lack, you should avoid using it on production, but it's ok for local testing.
3. Again, self-generated certificates are not intended for production, but if you look for running server somewhere on the internet, you may be interested in [autocert](https://godoc.org/golang.org/x/crypto/acme/autocert) package which helps your server to obtain Let's Encrypt certificate automatically. 
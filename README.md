# decode jwt

A go library to help you decode a JSON Web Token (JWT)

Please read [RFC 7519](https://tools.ietf.org/html/rfc7519) for more details on JWT's

## How to use this library

You can clone this library and run the following command:

```sh
go run jwtDecode.go -jwt=awtee123455
```

If you don't provide an argument of -jwt then you will receive a usage message

## Install library in System

Run the command:

```sh
go get github.com/jbelmont/decode-jwt
```

Assuming that you have go installed and have set a `GOPATH` you can simply run:

`jwt-decode -jwt=aTokenHere`

You might need to add `$GOPATH/bin` to your `PATH` variable in .bashrc and .zshrc configuration files.

## Sample Run

```sh
go run jwtDecode.go -jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c
{
 "raw": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
 "headers": {
  "alg": "HS256",
  "typ": "JWT"
 },
 "claims": {
  "iat": 1516239022,
  "sub": "1234567890",
  "name": "John Doe"
 },
 "signature": "",
 "valid": false
}
```

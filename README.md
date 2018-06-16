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
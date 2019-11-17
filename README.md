# Template Authentication API Golang

This project is designed to provide a good starter server to build web applications. It handles user features such as authentication, email validation, password reset, ...

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

What things you need to install the software and how to install them

```
Golang > 1.1
Postgresql server
Gmail account (Enable this feature : https://myaccount.google.com/lesssecureapps)
```

### Installing

```
git clone project
cd golang-template-api-authentication
```

```
go get ./
```

```
go build
```

Set your .env variables at https://github.com/Gandalfounet/golang-template-api-authentication/blob/master/.env and start the executable

```
./golang-template-api-authentication
```

### FEATURES

```
Login
```

```
Register
```

```
Send validation email after register
```

```
Send confirmation email after validation
```

```
Forgot password => Send link to modify password by email
```

```
Reset password => Update the password
```

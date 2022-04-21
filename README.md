# SoftwareDevelopment2022-Backend
This is repository for Lancaster Country Day School TSA Group Software development 2022


## Environment
Golang 1.17 

## Build
Docker file can be found in `./dockerfile/Dockerfile`

## Router
#### login: host:port/v1/auth/login
#### register: host:port/v1/auth/reg
#### verify code: host:port/v1/auth/send_verify

## APIs
#### login:
POST: {
    "email":"xxxx@gmail.com",
    "password":"1234"
}
#### register:
POST:{
    "email": "xxxxx@gmail.com",
    "nickname": "Kevin",
    "password": "123456",
    "verify_code": 123456
}
#### verify code:
POST: {
"email": "xxxx@gmail.com",
"target": "register"
}

### Note:
config.json will be generated once you first run the server. Path: `./files/config.json`
Database must be filled in order to run.
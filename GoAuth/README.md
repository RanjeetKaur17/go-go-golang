# GoAuth

This is a Simple Go Framework to Enable Session Based Authentication in any Go Application. In this Framework, we will use postgres to save login credentials and postgres/redis to create user sessions. Redis is the better alternate for maintaining sessions, since it will have better performance for Get and has in-built support for expiry. By default, we have configured it with Redis, but that can easily be switched in `main.go` file. 

Following are the details on how you can run it independently or import/use it into your code.

## Prerequisites
To Install Go, <a href="https://golang.org/doc/install">click here</a>.<br/><br/>
Add Redis and PSQL dependencies, using
```
go get -u github.com/lib/pq
go get -u github.com/go-redis/redis
```
Add gin-gonic dependency, using 
```
go get -u github.com/gin-gonic/gin
```
Add Swagger dependency, using
```
*go get -u github.com/swaggo/gin-swagger<br/>*
*go get -u github.com/swaggo/gin-swagger/swaggerFiles<br/>*
```
Create required schema.<br/>
```
CREATE TABLE users (
  username TEXT PRIMARY KEY,
  password TEXT
);
```
Create Following table, if Sessions are to be stored in database rather than redis
```
CREATE TABLE user_login (
  username TEXT,
  device_id TEXT,
  user_secure_key TEXT,
  PRIMARY KEY(username, device_id)
);
```

Start your redis intsance and database server and configure values in `config/config.json`.
```
{
  "Database" : {
    "Host" : "localhost",
    "Port" : "5432",
    "Username" : "root",
    "Password" : "password",
    "Dbname" : "user_info"
  },
  "Redis" : "localhost:6379"
}
```


## Build
To build this project, run following command in the project directory (where main file exists).<br/>

`go build .`<br/>

A binary file by the project name would be generated.
 
 ## Testing
 
If there are any test cases in the project, following command can be used to run those test cases.<br/>

`go test ./...`<br/>

## Execution

Once the binary is generated, execute the binary to run the program.<br/>

`./GoAuth`<br/>

## Usage

Once the instance is started, you can call following APIs, for Auth related functions.

#####To SignUp a new user.
```
POST http://localhost:8080/auth/signup
BODY :
{
    "username" : "username",
    "password" : "password"
}
RESPONSE
{
    "data" : <is_successful>,
    "exception" : <error_message>
}
```

#####To Create a user session.

```
POST http://localhost:8080/auth/login
BODY :
{
    "username" : "username",
    "password" : "password"
}
RESPONSE
{
    "data" : {
                "userSecureKey" : <user_secure_key>,
                "deviceId" : <assigned_device_id>
             },
    "exception" : <error_message>
}
```

#####To Get User Info.

```
GET http://localhost:8080/hello
HEADER  Authorization : {"userSecureKey" : <user_secure_key>, "deviceId" : <assigned_device_id>}
RESPONSE
{
    "data" : "Hello <username>!",
    "exception" : <error_message>
}
```

#####To End a user session.
```
POST http://localhost:8080/auth/logout
HEADER  Authorization : {"userSecureKey" : <user_secure_key>, "deviceId" : <assigned_device_id>}
RESPONSE
{
    "data" : <is_successful>,
    "exception" : <error_message>
}
```

Alternately, you can access and test these APIs, through the swagger page `http://localhost:8080/swagger/index.html`

## Import
To Import this into you code, download the package from github. Configure your database and redis connections in `config/config.json`. Copy following code in your code, and that should automatically enable login/ logout functionality.
```
controller.SetUp(model.CONFIG_PATH)
	controller.StartRoutesV1(&controller.AuthController{
		Service:&service.AuthService{
			UserDAO:&database.UserDAO{},
			UserSessionDAO:&cache.UserSessionCacheDAO{}}})
```

For the API's that are to be protected, make them use following method while configuring their routes.
```
router.Group("xyz").Use(auth.CheckAuthorization)
```

## Maintainers
<ul><a href="https://github.com/RanjeetKaur17">RanjeetKaur17</a></ul>

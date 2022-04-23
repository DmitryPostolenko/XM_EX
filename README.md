# XM Golang Exercise

[![Go Report Card](https://goreportcard.com/badge/github.com/DmitryPostolenko/lets-go-chat)](https://goreportcard.com/report/github.com/DmitryPostolenko/lets-go-chat)

REST API microservice to handle Companies. 
Company is an entity defined by the following attributes:
- Name
- Code
- Country
- Website
- Phone

To run the application [Go](https://golang.org/doc/install), 
[Docker](https://www.docker.com/get-started) and 
[Redis](https://redis.io/docs/getting-started/) must be installed

The application can be downloaded from GitHub using
- Web (https://github.com/DmitryPostolenko/XM_EX/archive/refs/heads/master.zip)
- Git (https://github.com/DmitryPostolenko/XM_EX.git)

Before running the application, start postgresql Docker container:
<pre>
docker run -it --rm --name go-postgres -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=secret -e PGDATA=/var/lib/postgresql/data/pgdata -v ~/local-go-postgres:/var/lib/postgresql/data postgres:14.0
</pre>

And ensure Redis is running (host:localhost, port:6379)

To run the application, start terminal from the application directory and run the command:
<pre>
    $ go mod tidy
    $ go run main.go
    OR
    $ go run .
</pre>

To run all tests, start terminal from the package directory and run the command:
<pre>
    $ go test ./...
</pre>

Application is also deployed to heroku. for testing purposes us url: https://xmex.herokuapp.com/ instead of http://localhost:8080/

App allows users from Cyprus (by IP address data, ipapi) non authorized access to companies management (creation, modification, deletion, listing)
In other cases user should register and login before interaction

To check locally:

- Create user:
<pre>
curl -v POST http://localhost:8080/v0.9/user/register -H 'Content-Type: application/json' -d '{"userName":"my_login","password":"my_password"}'
</pre>
Response on success:
<pre>
{
"id": "uuid_id",
"user_name": "my_login"
}
</pre>

- Login user:
<pre>
curl -v POST http://localhost:8080/v0.9/user/login -H 'Content-Type: application/json' -d '{"userName":"my_login","password":"my_password"}'
</pre>
Response on success:
<pre>
{
"token": "jwt_access_token"
}
</pre>

- Logout user:
<pre>
curl -v DELETE http://localhost:8080/v0.9/user/logout -H 'Content-Type: application/json' -d '{"token": "jwt_access_token"}'
</pre>
Response on success:
<pre>
{
    "msg": "Success"
}
</pre>

For interaction with companies(in case user is from Cyprus, token is omitted empty):

- Create company:
<pre>
curl -v POST http://localhost:8080/v0.9/company/ -H 'Content-Type: application/json' -d '{"token":"jwt_access_token", "name":"my_test_company","code":"2332323","country":"France","website":"https://something.fr","phone":"2323323"}'
</pre>
Response on success:
<pre>
{
    "id": "uuid_id",
    "company_name": "my_test_company"
}
</pre>

- List companies:
<pre>
curl -v GET http://localhost:8080/v0.9/company/list?token=jwt_access_token
</pre>
Response on success, lists all companies as objects array:
<pre>
[
...
    {
        "id": "uuid_id",
        "name": "my_test_company",
        "code": "2332323",
        "country": "France",
        "website": "https://something.fr",
        "phone": "2323323"
    },
...
]
</pre>

- Search companies, you can search by company fields (id, name, code, country, website, phone):
<pre>
curl -v GET http://localhost:8080/v0.9/ccompany?field=code&value=23323&token=jwt_access_token
</pre>
Response on success, lists found companies as objects array:
<pre>
[
...
    {
        "id": "uuid_id",
        "name": "my_test_company",
        "code": "2332323",
        "country": "France",
        "website": "https://something.fr",
        "phone": "2323323"
    },
...
]
</pre>

- Modify company:
<pre>
curl -v PUT http://localhost:8080/v0.9/company/ -H 'Content-Type: application/json' -d '{"token": "jwt_access_token","id":"fd66824d-521f-4e6f-63d7-88da84500663","name":"www","code":"2332323","country":"Ukraine","website":"https://something.com","phone":"2323323"}'
</pre>
Response on success:
<pre>
{
    "msg": "Success"
}
</pre>

- Delete company:
<pre>
curl -v DELETE http://localhost:8080/v0.9/company/86e9860c-d11b-4317-7625-c95ee3db87c7?token=jwt_access_token
</pre>
Response on success:
<pre>
{
    "msg": "Success"
}
</pre>
# XM Golang Exercise

[![Go Report Card](https://goreportcard.com/badge/github.com/DmitryPostolenko/lets-go-chat)](https://goreportcard.com/report/github.com/DmitryPostolenko/lets-go-chat)

REST API microservice to handle Companies. Company is an entity defined by the following attributes:
- Name
- Code
- Country
- Website
- Phone

To run the application [Go](https://golang.org/doc/install), [Docker](https://www.docker.com/get-started) and [Redis](https://redis.io/docs/getting-started/) must be installed

The application can be downloaded from GitHub using
- Web (https://github.com/DmitryPostolenko/XM_EX/archive/refs/heads/master.zip)
- Git (https://github.com/DmitryPostolenko/XM_EX.git)

Before running the application, start postgresql Docker container:
<pre>
docker run -it --rm --name go-postgres -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=secret -e PGDATA=/var/lib/postgresql/data/pgdata -v ~/local-go-postgres:/var/lib/postgresql/data postgres:14.0
</pre>

To run the application, start terminal from the application directory and run the command:
<pre>
    $ go mod tidy
    $ go run main.go
    OR
    $ go run .
</pre>


Find company
http://localhost:8080/v0.9/company?field=code&value=23323
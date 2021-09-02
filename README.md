# GORMSeeder - A simple database seeder based on GORM

The GORMSeeder is a lightweight but powerful and flexible seeder tool based on GORM. Especially useful in container environments like Docker.

The goal is to create a build of your seeder setup.

## Steps for deployment

* Create a go build
* Putting build in a `FROM scratch AS bin` container for minimal size
* Deploy
* Start service or container with environment variables to perform a seed on database

## Example

See example under `example/` to see how it works.

The folder looks like 
```bash
├── Seeder
│   ├── main.go
│   ├── seed00001_create-mock-user.go
│   ├── seed00002_create-locales.go
│   └── seed00003_create-products.go
    ...
```

`seed`-files can have any names.

Example for `main.go`:
```golang
package main

import (
	"fmt"

	gs "github.com/randree/gormseeder"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(postgres.Open("host=localhost user=user password=passpass dbname=testdb port=5432 sslmode=disable"))
	if err != nil {
		fmt.Println(err.Error())
	}

	gs.InitSeeder(db, "seeders")
}
```

`seed`-files (e.g. `seed001_mock-user-list.go`) looks like:
```golang
package main

import (
	gs "github.com/randree/gormseeder"
	"gorm.io/gorm"
)

func init() {
	gs.Seed(gs.State{

		Tag: "<tag-name>",

		Perform: func(db *gorm.DB) error {
			
			...

			return err
		},
	})
}
```

## Create and use module

Steps to create a go module:
```console
$ go mod init Seeder
```
To load dependencies:
```console
$ go mod tidy
```

To run a Seeder:
```console
$ TAG=<tag-name> (go run ./... | <go-build>)
```

With `TAG=all` you can perform all seeds in one go.

To show a Seeder history:
```console
$ HISTORY=1 (go run ./... | <go-build>)

| DATETIME HISTORY                       | FILENAME                  | USER       |
| -------------------------------------- | ------------------------- | ---------- |
| 2021-08-31 13:09:47.932619 +0200 CEST  | seed003_mock-user-list.go | admin      |
| 2021-08-31 13:09:47.908876 +0200 CEST  | seed002_mock-products.go  | foo        |
| 2021-08-31 13:09:47.871357 +0200 CEST  | seed001_customers.go      | bar        |
```

To show version:
```console
$ VERSION=1 (go run ./... | <go-build>)

Gormseeder version:  0.1.0
```

### Calls

```console
$ TAG=<Tag> [HISTORY=1] [VERSION=1] (Docker Container | go build | go run ./...)
```
Docker-compose file
```yaml
...
  migrator:
    image: from_scratch_image
    environment:
      TAG: <Tag>
...

```
If you want to seed everything use `TAG=all`.

## References

- [GROM](https://gorm.io/) The GORM project.
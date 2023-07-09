# atlas-provider-beego

Load [beego](https://github.com/beego/beego) schemas into an [Atlas](https://atlasgo.io) project.

### Use-cases
1. **Declarative migrations** - use a Terraform-like `atlas schema apply --env beego` to apply your beego schema to the database.
2. **Automatic migration planning** - use `atlas migrate diff --env beego` to automatically plan a migration from  
  the current database version to the beego schema.

### Installation

Install Atlas from macOS or Linux by running:
```bash
curl -sSf https://atlasgo.sh | sh
```
See [atlasgo.io](https://atlasgo.io/getting-started#installation) for more installation options.

Install the provider by running:
```bash
go get -u ariga.io/atlas-provider-beego
``` 

#### Standalone 

If application contains a package which registers all of its beego models during initialization,
you can use the provider directly to load your beego schema into Atlas. 

In your project directory, create a new file named `atlas.hcl` with the following contents:

```hcl
data "external_schema" "beego" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "ariga.io/atlas-provider-beego",
    "load",
    "--path", "./path/to/models",
    "--dialect", "mysql", // | postgres | sqlite3
  ]
}

env "beego" {
  src = data.external_schema.beego.url
  dev = "docker://mysql/8/dev"
  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}
```

#### As Go File

If you want to use the provider as a Go file, you can use the provider as follows:

Create a new program named `loader/main.go` with the following contents:

```go
package main

import (
  "io"
  "os"

  "ariga.io/atlas-provider-beego/beegoschema"
  "github.com/<yourorg>/<yourrepo>/path/to/models"
)

func main() {
  stmts, err := beegoschema.New("mysql").Load()
  if err != nil {
    fmt.Fprintf(os.Stderr, "failed to load beego schema: %v\n", err)
    os.Exit(1)
  }
  io.WriteString(os.Stdout, stmts)
}

// If your models are already registered in an init() function elsewhere, you can simply use
// a blank import to ensure that the init() function is called. Otherwise, you can register
// your models here.
func init() {
	orm.RegisterModel(new(models.User), new(models.Group))
}
```

In your project directory, create a new file named `atlas.hcl` with the following contents:

```hcl
data "external_schema" "beego" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "./loader",
  ]
}

env "beego" {
  src = data.external_schema.beego.url
  dev = "docker://mysql/8/dev"
  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}
```

### Usage

Once you have the provider installed, you can use it to apply your `beego` schema to the database.

#### Apply

You can use the `atlas schema apply` command to plan and apply a migration of your database to
your current `beego` schema. This works by inspecting the target database and comparing it to the
desired schema and creating a migration plan. Atlas will prompt you to confirm the migration plan
before applying it to the database.

```bash
atlas schema apply --env beego -u "mysql://root:password@localhost:3306/mydb"
```
Where the `-u` flag accepts the [URL](https://atlasgo.io/concepts/url) to the
target database.

#### Diff

Atlas supports a [version migration](https://atlasgo.io/concepts/declarative-vs-versioned#versioned-migrations) 
workflow, where each change to the database is versioned and recorded in a migration file. You can use the
`atlas migrate diff` command to automatically generate a migration file that will migrate the database
from its latest revision to the current `beego` schema.

```bash
atlas migrate diff --env beego 
```

### Supported Databases

The provider supports the following databases:
* MySQL
* PostgreSQL
* SQLite

### Issues

Please report any issues or feature requests in the [ariga/atlas](https://github.com/ariga/atlas/issues) repository.

### License

This project is licensed under the [Apache License 2.0](LICENSE).
variable "dialect" {
  type = string
}

locals {
  dev_url = {
    mysql = "docker://mysql/8/dev"
    postgres = "docker://postgres/15"
    sqlite3 = "sqlite://file::memory:?cache=shared"
  }[var.dialect]
}

data "external_schema" "beego" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "ariga.io/atlas-provider-beego",
    "load",
    "--path", "./models",
    "--dialect", var.dialect,
  ]
}

env "beego" {
  src = data.external_schema.beego.url
  dev = local.dev_url
  migration {
    dir = "file://migrations/${var.dialect}"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}

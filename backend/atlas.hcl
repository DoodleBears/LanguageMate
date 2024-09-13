# The "local" environment represents our local testings.
env "local" {
  src = "file://manifest/database/tables/"
  dev = "docker://maria/latest/languagemate"
  url = "maria://languagemate:Passw0rd@localhost:3306/languagemate"
  migration {
    dir = "file://manifest/database/migrations/"
  }
}

# The "docker" environment represents our docker development.
env "docker" {
  src = "file://manifest/database/tables/"
  dev = "docker://maria/latest/languagemate"
  url = "maria://languagemate:Passw0rd@database:3306/languagemate"
  migration {
    dir = "file://manifest/database/migrations/"
  }
}
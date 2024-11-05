pkgname = "age"
pkgver = "1.2.0"
pkgrel = 6
build_style = "go"
make_build_args = ["-ldflags=-X main.Version={pkgver}", "./cmd/..."]
hostmakedepends = ["go"]
pkgdesc = "File encryption tool"
maintainer = "psykose <alice@ayaya.dev>"
license = "BSD-3-Clause"
url = "https://github.com/FiloSottile/age"
source = "{url}/archive/refs/tags/v{pkgver}.tar.gz"
sha256 = "cefe9e956401939ad86a9c9d7dcf843a43b6bcdf4ee7d8e4508864f227a3f6f0"

build_script {
  commands = [
    ["go", "build"]
  ]

  overrides {

  }
}

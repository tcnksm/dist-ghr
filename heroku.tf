variable "heroku_email" {}
variable "heroku_api_key" {}

provider "heroku" {
  email = "${var.heroku_email}"
  api_key = "${var.heroku_api_key}"
}

resource "heroku_app" "default" {
  name = "ghr"
  stack = "cedar"
  config_vars {
    BUILDPACK_URL="https://github.com/kr/heroku-buildpack-go.git"
    BASE_URL="https://github.com/tcnksm/ghr/releases"
    DIST_NAME="ghr"
    VERSION="v0.1.0"
  }
}

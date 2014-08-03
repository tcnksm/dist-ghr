re-dist-ghr
====

Redirect distribution for [tcnksm/ghr](https://github.com/tcnksm/ghr)

## Description

To enable user to downlaod packages, with **same url** and **same command**.

```
$ L=/usr/local/bin/ghr && curl -sL -A "`uname -sp`" http://localhost:3000/ghr.zip | zcat >$L && chmod +x $L
```

[http://localhost:3000]() will attempt to detect your OS and CPU architecture based on the User-Agent, then redirect you to the latest release for your platform.

## URL

Artifacts is hosted on github relase page, [https://github.com/tcnksm/ghr/releases/download/${VERSION}/ghr\_${VERSION}\_${OS}\_${ARCH}.zip]()


## Terraform

To create and update application on Heroku, use [Terraform](http://www.terraform.io/)

To test:

```bash
$ terraform plan -var heroku_email=$HEROKU_EMAIL -var heroku_api_key=$HEROKU_API_KEY
```

To create and update:

```bash
$ terraform apply -var heroku_email=$HEROKU_EMAIL -var heroku_api_key=$HEROKU_API_KEY
```

## Config

To run this application, need to set some environmetal variables:

- **BASE_URL** - Base url where binary is hosted
    - [https://github.com/tcnksm/ghr/releases]()
- **DIST_NAME** - Distribution name of binary
    - ghr
- **VERSION** - Version of binary

These variables are set by [Terraform](http://www.terraform.io/)

## Test

After setting environmental variable, just execute:

```bash
$ go run web.go
```

## Licence

[MIT](https://github.com/tcnksm/dist-ghr/blob/master/LICENCE)

## Author

[tcnksm](https://github.com/tcnksm)




# Go hello

## Starting a new project

- Create a project on `github.com`

- Clone your project

```sh
mkdir -p $GOPATH/src/github.com/<your-username>
cd $GOPATH/src/github.com/<your-username>
git clone git@github.com:<your-username>/<your-project>
cd <your-project>
```

## go getable application

If you follow the common conventions, your application is automatically `go getable`. If you commit and push this single file to `github`, anyone with a working Go installation should be able to get it:

```sh
go get github.com/<your-username>/<your-project>
$GOPATH/bin/hello
```

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

## Weather API

[API](https://openweathermap.org/appid)

`api.openweathermap.org/data/2.5/forecast?id=524901&APPID=1111111111`

A possible answer from API can look like this one:

```json
{
    "name": "Tokyo",
    "coord": {
        "lon": 139.69,
        "lat": 35.69
    },
    "weather": [
        {
            "id": 803,
            "main": "Clouds",
            "description": "broken clouds",
            "icon": "04n"
        }
    ],
    "main": {
        "temp": 296.69,
        "pressure": 1014,
        "humidity": 83,
        "temp_min": 295.37,
        "temp_max": 298.15
    }
}
```

```go
type weatherData struct {
    Name string `json:"name"`
    Main struct {
        Kelvin float64 `json:"temp"`
    } `json:"main"`
}
```

```go
type weatherData struct {
    Cod string `json:"cod"`
    City struct {
        Name string `json:"name"`
    } `json:"city"`
}
```

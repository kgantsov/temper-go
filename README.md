## TEMPer GO

Small libusb/gousb-based driver to read TEMPer USB HID devices (USB ID 0c45:7401)

Tested on this device https://www.amazon.com/TEMPer-USB-Thermometer-w-Alerts/dp/B002VA813U


### Requirements

The `libusb` and `pkg-config` are need to be installed

```bash
apt-get install -y libusb-1.0-0-dev pkg-config
```

### Instalation

```bash
github.com/kgantsov/temper-go
```


### Usage

After the package is installed it can be used in go like this:

```go
package main

import (
    "log"

    "github.com/kgantsov/temper-go/pkg"
)

func main() {
    temp, err := temper.GetTemperature()

    if err == nil {
        log.Printf(
            "Temperature: %.2f°K %.2f°F %.2f°C\n",
            temp.Temperature+273.15,
            9.0/5.0*temp.Temperature+32,
            temp.Temperature,
        )
    } else {
        log.Fatalf("Failed: %s", err)
    }
}
```


### CLI utility

Also there is a utility in `github.com/kgantsov/temper-go/cmd/temper` and
`github.com/kgantsov/temper-go/cmd/temp-logger`.
You can easily build and use them to check current temperature of log it

```bash
cd cmd/temper
go build
```

Now you can run the `./temper` and you should see similar to this result:

```bash
vagrant@temper:~/golang/src/github.com/kgantsov/temper-go/cmd/temper$ sudo ./temper
2017/09/20 15:58:22 Temperature: 293.27°K 68.22°F 20.12°C


cd cmd/tem-logger
go build
```

Now you can run the `./temp-logger` and you should see similar to this result:

```bash
vagrant@temper:~/golang/src/github.com/kgantsov/temper-go/cmd/temp-logger$ sudo ./temp-logger
2017/09/20 16:00:31 Temperature: 297.09°K 75.09°F 23.94°C
2017/09/20 16:00:32 Temperature: 296.96°K 74.86°F 23.81°C
2017/09/20 16:00:33 Temperature: 296.90°K 74.75°F 23.75°C
2017/09/20 16:00:34 Temperature: 296.90°K 74.75°F 23.75°C
2017/09/20 16:00:35 Temperature: 296.84°K 74.64°F 23.69°C
2017/09/20 16:00:36 Temperature: 296.77°K 74.53°F 23.62°C
2017/09/20 16:00:37 Temperature: 296.77°K 74.53°F 23.62°C
2017/09/20 16:00:38 Temperature: 296.77°K 74.53°F 23.62°C
2017/09/20 16:00:39 Temperature: 296.71°K 74.41°F 23.56°C
2017/09/20 16:00:40 Temperature: 296.71°K 74.41°F 23.56°C
```

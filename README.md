## TEMPer GO

Small libusb/gousb-based driver to read TEMPer USB HID devices (USB ID 0c45:7401)


### Requirements

The `libusb` and `pkg-config` are need to be installed

    apt-get install -y libusb-1.0-0-dev pkg-config

### Instalation

    github.com/kgantsov/temper-go


### Usage

After the package is installed it can be used in go like this:

    package main

    import (
    	"github.com/kgantsov/temper-go"
    	"log"
    )

    func main() {
    	temp, err := temper.GetTemperature()

    	if err == nil {
    		log.Printf("Temperature: %.2fK %.2fF %.2fC\n", temp+273.15, 9.0/5.0*temp+32, temp)
    	} else {
    		log.Fatalf("Failed: %s", err)
    	}
    }


### CLI utility

Also there is a utility in `github.com/kgantsov/temper-go/cmd/temper` and
`github.com/kgantsov/temper-go/cmd/temp-logger`.
You can easily build and use them to check current temperature of log it

    cd cmd/temper
    go build

Now you can run the `./temper` and you should see similar to this result:

    vagrant@temper:~/golang/src/github.com/kgantsov/temper-go/cmd/temper$ sudo ./temper
    2017/09/18 19:11:38 Temperature: 292.84K 67.44F 19.69C


    cd cmd/tem-logger
    go build

Now you can run the `./temp-logger` and you should see similar to this result:

    vagrant@temper:~/golang/src/github.com/kgantsov/temper-go/cmd/temp-logger$ sudo ./temp-logger
    2017/09/18 19:13:34 Temperature: 292.84K 67.44F 19.69C
    2017/09/18 19:13:35 Temperature: 292.84K 67.44F 19.69C
    2017/09/18 19:13:35 Temperature: 292.84K 67.44F 19.69C
    2017/09/18 19:13:35 Temperature: 292.84K 67.44F 19.69C
    2017/09/18 19:13:35 Temperature: 292.84K 67.44F 19.69C
    2017/09/18 19:13:35 Temperature: 292.84K 67.44F 19.69C
    2017/09/18 19:13:35 Temperature: 293.27K 68.22F 20.12C
    2017/09/18 19:13:35 Temperature: 293.27K 68.22F 20.12C
    2017/09/18 19:13:36 Temperature: 293.46K 68.56F 20.31C
    2017/09/18 19:13:36 Temperature: 293.46K 68.56F 20.31C
    2017/09/18 19:13:36 Temperature: 293.59K 68.79F 20.44C
    2017/09/18 19:13:36 Temperature: 293.59K 68.79F 20.44C
    2017/09/18 19:13:36 Temperature: 293.84K 69.24F 20.69C
    2017/09/18 19:13:36 Temperature: 293.84K 69.24F 20.69C
    2017/09/18 19:13:36 Temperature: 293.84K 69.24F 20.69C
    2017/09/18 19:13:36 Temperature: 294.09K 69.69F 20.94C
    2017/09/18 19:13:37 Temperature: 294.09K 69.69F 20.94C
    2017/09/18 19:13:37 Temperature: 294.09K 69.69F 20.94C
    2017/09/18 19:13:37 Temperature: 294.09K 69.69F 20.94C
    2017/09/18 19:13:37 Temperature: 294.40K 70.25F 21.25C
    2017/09/18 19:13:37 Temperature: 294.40K 70.25F 21.25C
    2017/09/18 19:13:37 Temperature: 294.40K 70.25F 21.25C
    2017/09/18 19:13:37 Temperature: 294.52K 70.47F 21.38C
    2017/09/18 19:13:37 Temperature: 294.52K 70.47F 21.38C
    2017/09/18 19:13:38 Temperature: 294.71K 70.81F 21.56C
    2017/09/18 19:13:38 Temperature: 294.71K 70.81F 21.56C
    2017/09/18 19:13:38 Temperature: 294.77K 70.93F 21.62C
    2017/09/18 19:13:38 Temperature: 294.77K 70.93F 21.62C
    2017/09/18 19:13:38 Temperature: 294.84K 71.04F 21.69C
    2017/09/18 19:13:38 Temperature: 294.84K 71.04F 21.69C
    2017/09/18 19:13:38 Temperature: 294.84K 71.04F 21.69C



package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/MarinX/keylogger"
	"github.com/go-vgo/robotgo"
)

func dockerize(counter int, command string, lazyRecursion int) {
	if counter < 1 {
		robotgo.TypeString(command)
		robotgo.KeyTap("enter")
		return
	}
	robotgo.TypeString(`docker run --name blabl --privileged --network="host" -it jpetazzo/dind`)
	robotgo.KeyTap("enter")
	marker := 0
	timeWait := 0
	if lazyRecursion-counter == 1 {
		timeWait = 5
	} else {
		timeWait = 24
	}
	for marker < timeWait {
		time.Sleep(time.Second)
		marker += 1
		fmt.Println("Slept for", marker, "seconds")
	}
	dockerize(counter-1, command, lazyRecursion)
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: ./containerize NUMBER COMMAND [rest of COMMAND]")
		return
	}
	devs, err := keylogger.NewDevices()
	if err != nil {
		log.Fatal(err)
	}
	numberIterations, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	finalCommand := strings.Join(os.Args[2:], " ")

	for _, val := range devs {
		fmt.Println("Id->", val.Id, "Device->", val.Name)
	}

	//keyboard device file, on your system it will be diffrent!
	rd := keylogger.NewKeyLogger(devs[3])

	in, err := rd.Read()
	if err != nil {
		log.Fatal(err)
	}

	for i := range in {

		//we only need keypress
		if i.Type == keylogger.EV_KEY {
			if i.KeyString() == "L_CTRL" {
				time.Sleep(time.Second)
				dockerize(numberIterations, finalCommand, numberIterations+1)
				return
			}
		}
	}
}

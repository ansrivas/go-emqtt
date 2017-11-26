package main

import (
	"fmt"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/mqtt"
	"time"
)

func publisher() {
	mqttAdaptor := mqtt.NewAdaptorWithAuth("tcp://url:1883", "clientid", "username", "password")

	// mqttAdaptor := mqtt
	work := func() {
		// mqttAdaptor.On("hello/world", func(msg mqtt.Message) {
		// 	fmt.Println("hello message published")
		// 	fmt.Println(msg.Payload())
		//
		// })

		gobot.Every(time.Microsecond, func() {
			data := []byte(fmt.Sprintf("Hello world from golang %d", time.Now().UTC()))
			mqttAdaptor.Publish("hello/world", data)
		})
	}

	robot := gobot.NewRobot("mqttBot",
		[]gobot.Connection{mqttAdaptor},
		work,
	)

	robot.Start()
}
func main() {
	publisher()
}

package heartbeat

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

const LOG_FILE = "/go/app/heartbeat.log"

func Pulse() {
	WriteHeartbeat()
	for {
		timer := time.After(time.Second * 10)
		<-timer
		WriteHeartbeat()
	}
}

func WriteHeartbeat() {
	t := time.Now()
	f, err := os.OpenFile(LOG_FILE, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	f.WriteString(t.Format(time.RFC3339))
}

func CheckHeartbeat() {
	t := time.Now()
	bytes, err := ioutil.ReadFile(LOG_FILE)
	if err != nil {
		panic(err)
	}

	pulseTime, err := time.Parse(time.RFC3339, string(bytes))
	if err != nil {
		panic(err)
	}
	if t.Sub(pulseTime) > time.Minute {
		panic("heartbeat stop!")
	}
}

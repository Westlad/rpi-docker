package main

import(
	"github.com/stianeikeland/go-rpio"
	"encoding/json"
	"net/http"
	"log"
	"fmt"
)

const PIN int = 18
type heaterState struct {
        State string
}

func main() {
	pin := GPIOinit(PIN)
	http.HandleFunc("/",heater(pin))
	http.ListenAndServe(":8080", nil)
}

func GPIOinit(PIN int) rpio.Pin {
	err := rpio.Open()
	if err!=nil {
		panic(err)
	}
	pin := rpio.Pin(PIN)
	pin.Output()
	return pin
}

func heater(pin rpio.Pin) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//get the response JSON
		decoder := json.NewDecoder(r.Body)

		var t heaterState
		err := decoder.Decode(&t)
		if err != nil {
			panic(err)
		}

		switchHeater(t,pin)
		//defer r.Body.Close()
		log.Println(t.State)
		j,err := json.Marshal(t)
		if err!=nil {
			panic(err)
		}
		fmt.Println(t.State)
		w.Write(j)
	}
}

func switchHeater(t heaterState, pin rpio.Pin) {
	//switch the heater on or off
	if t.State=="on" { //turn heater on
		pin.High()
	} else {
		pin.Low()
	}
}


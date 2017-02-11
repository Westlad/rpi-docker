package main

import(
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/",temperature)
	http.ListenAndServe(":8080", nil)
}

type Message struct {
        Temperature float64
}

func temperature(w http.ResponseWriter, r *http.Request) {
	//read the thermometer
	tp := 23.4
	m := Message{tp}
	j, err := json.Marshal(m)

	if err != nil {
		panic(err)
	}
	w.Write(j)
}



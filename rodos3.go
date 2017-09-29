package main

import (
	"fmt"
	"time"
	"net/http"
	"github.com/karalabe/hid"
)

var ON = []byte{0xE7, 0x00, 0x00, 0, 0, 0, 0, 0}
var OFF = []byte{0xE7, 0x19, 0x19, 0, 0, 0, 0, 0}
var handle *hid.Device
var err error = nil

func handler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path[1:] == "on" {
    fmt.Fprintf(w, "Switching relay to %s!", r.URL.Path[1:])
    handle.Write(ON)
    }

    if r.URL.Path[1:] == "off" {
    fmt.Fprintf(w, "Switching relay to %s!", r.URL.Path[1:])
    handle.Write(OFF)
    }
}

func main() {
    devices := hid.Enumerate(8352, 16755)
    if len(devices) > 0 {
      fmt.Println("Found devices:", devices)
    } else {
      fmt.Println("No rodos-3 switches found!")
    }
    rodos := devices[0]
    handle, err = rodos.Open()
    if err != nil {
	panic(err)
    }

    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
    handle.Close()
}

func test() {
    devices := hid.Enumerate(8352, 16755)
    if len(devices) > 0 {
      fmt.Println("Found devices:", devices)
    } else {
      fmt.Println("No rodos-3 switches found!")
    }
    rodos := devices[0]
    handle, err := rodos.Open()
    if err != nil {
        panic(err)
    }
    fmt.Println("Switching to ON...")
    handle.Write(ON)
    time.Sleep(5000 * time.Millisecond)
    fmt.Println("switching to OFF...")
    handle.Write(OFF)
    handle.Close()
}

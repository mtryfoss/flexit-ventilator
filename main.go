package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
)

func main() {
	var destination string
	var action string
	var value string
	flag.StringVar(&destination, "d", "", "Destination IP")
	flag.StringVar(&action, "a", "", "Fan action")
	flag.StringVar(&value, "v", "", "Value")
	flag.Parse()

	if len(destination) == 0 {
		fmt.Println("Invalid destination")
		flag.PrintDefaults()
		return
	}

	// Initial byte sequence
	cmd := []byte{0x6D, 0x6F, 0x62, 0x69, 0x6C, 0x65}
	switch action {
	case "speedlevel":
		level, err := strconv.Atoi(value)
		if err != nil || level < 22 || level > 255 {
			fmt.Printf("Invalid speed level: %d (between 22 and 255 are valid)\n", level)
			flag.PrintDefaults()
			return
		}
		// set speed level action
		cmd = append(cmd, 0x05)
		// actual speed
		cmd = append(cmd, byte(level))
	default:
		fmt.Printf("Valid actions are: speedlevel\n")
		return
	}

	// End byte sequence
	cmd = append(cmd, []byte{0x0D, 0x0A}...)

	conn, _ := net.Dial("udp", fmt.Sprintf("%s:4000", destination))
	fmt.Fprintf(conn, string(cmd))

}

package main

import (
	"fmt"
	"github.com/karalabe/hid"
	"log"
	"time"
)

func main() {

	infoList := hid.Enumerate(0x0922, 0x8003)

	for i := 0; i < len(infoList); i++ {

		info := infoList[i]

		fmt.Printf("ID %04x:%04x  %s\n",

			info.VendorID,
			info.ProductID,
			info.Path,
			// info.Release,
			// info.UsagePage,
			// info.Usage,
			// info.Interface
		)

		device, err := info.Open()

		if err != nil {
			log.Fatal(err)
		}

		buffer := make([]byte, 512)

		for {
			read, err := device.Read(buffer)
			if err != nil {
				log.Fatal(err)
			}

			if read > 0 {
				fmt.Println(buffer[4] + buffer[5]*255)
			}

			time.Sleep(1000)
		}

		defer device.Close()
	}
}

package leetcode

import "fmt"

func numberOfBeams(bank []string) int {
	beams, prevDevices, devicesInRow := 0, 0, 0

	for _, devices := range bank {
		devicesInRow = 0

		for i := 0; i < len(devices); i++ {
			if devices[i] == '1' {
				devicesInRow++
			}
		}

		fmt.Printf("Devices in a row: %s, Count: %d\n", devices, devicesInRow)

		if devicesInRow > 0 {
			beams += prevDevices * devicesInRow

			fmt.Printf("Beams between 2 row: %d x %d = %d\n", prevDevices, devicesInRow, prevDevices*devicesInRow)
			fmt.Println("-----------------------------")

			prevDevices = devicesInRow
		} else {
			fmt.Println("-----------------------------")
		}
	}
	fmt.Printf("Final Laser Beams: %d\n", beams)

	return beams
}

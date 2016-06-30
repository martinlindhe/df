package main

import (
	"fmt"
	"os"

	"github.com/ricochet2200/go-disk-usage/du"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	kb = float64(1024)
	mb = float64(kb * kb)
	gb = float64(mb * kb)

	humanMode = kingpin.Flag("human-readable", "Print sizes in powers of 1024 (e.g., 1023M).").Short('h').Bool()
	version   = kingpin.Flag("version", "Output version information and exit.").Bool()
)

func main() {

	kingpin.Parse()

	if *version {
		fmt.Println("df (go-df) 0.1")
		fmt.Println("https://github.com/martinlindhe/go-df")
		os.Exit(0)
	}

	drives := getWinDrives()

	if *humanMode {
		heading := []string{"Filesystem", "Size", "Used", "Avail", "Use%", "Mounted on"}

		renderColumnData(generateColumnData(drives, heading, gb, "G"))
	} else {

		heading := []string{"Filesystem", "1K-blocks", "Used", "Available", "Use%", "Mounted on"}

		renderColumnData(generateColumnData(drives, heading, kb, ""))
	}
}

func generateColumnData(drives []string, heading []string, divider float64, unit string) ([][]string, []bool) {
	data := [][]string{}

	rightAlignedColumns := []bool{false, true, true, true, true, false}

	data = append(data, heading)

	for _, drive := range drives {
		drive = drive + ":\\"

		usage := du.NewDiskUsage(drive)
		usageSize := fmt.Sprintf("%.f"+unit, float64(usage.Size())/divider)
		usageUsed := fmt.Sprintf("%.f"+unit, float64(usage.Used())/divider)
		usageAvail := fmt.Sprintf("%.f"+unit, float64(usage.Available())/divider)
		usagePct := fmt.Sprintf("%.f%%", usage.Usage()*100)
		mountPoint := "n/a" // XXX what to show?

		data = append(data, []string{
			drive, usageSize, usageUsed, usageAvail, usagePct, mountPoint,
		})
	}
	return data, rightAlignedColumns
}

func renderColumnData(data [][]string, rightAlignedColumns []bool) {

	columnWidths := calcColumnWidths(data)

	fmtStr := ""
	for _, row := range data {

		for idx, col := range row {
			w := columnWidths[idx] + 1
			right := rightAlignedColumns[idx]
			if right {
				fmtStr = fmt.Sprintf("%% %ds", w)
			} else {
				fmtStr = fmt.Sprintf("%% -%ds", w)
				if idx > 0 {
					fmtStr = " " + fmtStr
				}
			}
			fmt.Printf(fmtStr, col)
		}
		fmt.Println()
	}
}

func calcColumnWidths(data [][]string) []int {

	columns := len(data[0])
	widths := make([]int, columns)

	for _, row := range data {
		for idx, col := range row {
			len := len(col)
			if len > widths[idx] {
				widths[idx] = len
			}
		}
	}

	// ensure some min-widths
	if widths[0] < 13 {
		widths[0] = 13
	}
	if widths[2] < 5 {
		widths[2] = 5
	}

	return widths
}

// returns all drive letters in use on windows
func getWinDrives() (r []string) {
	for _, drive := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		_, err := os.Open(string(drive) + ":\\")
		if err == nil {
			r = append(r, string(drive))
		}
	}
	return
}

package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/gookit/color"
)

func drawGraph(value, max, width int, char rune, clr color.Color) int {
	nChars := value * width / max
	clr.Print(strings.Repeat(string(char), nChars))
	return nChars
}

func main() {
	graphWidth := 100
	interval := "2"
	if len(os.Args) == 2 {
		interval = os.Args[1]
	} else if len(os.Args) > 2 {
		interval = os.Args[1]
		graphWidthArg, err := strconv.Atoi(os.Args[2])
		if err == nil {
			graphWidth = graphWidthArg
		}
	}

	for {
		output, _ := exec.Command("vmstat", interval, "2").Output()
		lines := strings.Split(string(output), "\n")
		line := lines[len(lines)-2]
		cols := strings.Fields(line)

		us, _ := strconv.Atoi(cols[12])
		sy, _ := strconv.Atoi(cols[13])
		id, _ := strconv.Atoi(cols[14])
		wa, _ := strconv.Atoi(cols[15])
		st, _ := strconv.Atoi(cols[16])

		total := us + sy + id + wa + st

		now := time.Now()
		timeStamp := now.Format("15:04:05")
		fmt.Printf("%s ", timeStamp)
		
		curWidth := 0
		curWidth += drawGraph(us, total, graphWidth, '■', color.Green)
		curWidth += drawGraph(sy, total, graphWidth, '■', color.Red)
		curWidth += drawGraph(id, total, graphWidth, '■', color.Blue)
		curWidth += drawGraph(wa, total, graphWidth, '■', color.Yellow)
		color.White.Print(strings.Repeat("■", graphWidth-curWidth))

		fmt.Printf(" ")
		color.Green.Print("■us:" + fmt.Sprintf("%d ", us))
		color.Red.Print("■sy:" + fmt.Sprintf("%d ", sy))
		color.Blue.Print("■id:" + fmt.Sprintf("%d ", id))
		color.Yellow.Print("■wa:" + fmt.Sprintf("%d ", wa))
		color.White.Print("■st:" + fmt.Sprintf("%d\n", st))

	}
}

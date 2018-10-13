package main
import (
 "bufio"
 "fmt"
 "image"
 "image/color"
 "image/png"
 "io"
 "os"
 "path/filepath"
 "regexp"
 "strconv"
)
var m *image.NRGBA
var x int
var y int
var barWidth int

func findIP(input string) string {
 partIP := "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])"
 grammar := partIP + "\\." + partIP + "\\." + partIP + "\\." + partIP
 matchMe := regexp.MustCompile(grammar)
 return matchMe.FindString(input)
}
func plotBar(width int, height int, color color.RGBA) {
 xx := 0
 for xx < barWidth {

 yy := 0
 for yy < height {
 m.Set(xx+width, y-yy, color)
 yy = yy + 1
 }
 xx = xx + 1
 }
}

func getColor(x int) color.RGBA {
 switch {
 case x == 0:
 return color.RGBA{0, 0, 255, 255}
 case x == 1:
 return color.RGBA{255, 0, 0, 255}
 case x == 2:
 return color.RGBA{0, 255, 0, 255}
 case x == 3:
 return color.RGBA{255, 255, 0, 255}
 case x == 4:
 return color.RGBA{255, 0, 255, 255}
 case x == 5:
 return color.RGBA{0, 255, 255, 255}
 case x == 6:
 return color.RGBA{255, 100, 100, 255}
 case x == 7:
 return color.RGBA{100, 100, 255, 255}
 case x == 8:
 return color.RGBA{100, 255, 255, 255}
 case x == 9:
 return color.RGBA{255, 255, 255, 255}
 }
 return color.RGBA{0, 0, 0, 255}
}

func main() {
 var data []int
 arguments := os.Args
 if len(arguments) < 4 {
 fmt.Printf("%s X Y IP input\n", filepath.Base(arguments[0]))
 os.Exit(0)
 }
 x, _ = strconv.Atoi(arguments[1])
 y, _ = strconv.Atoi(arguments[2])
 WANTED := arguments[3]
 fmt.Println("Image size:", x, y)

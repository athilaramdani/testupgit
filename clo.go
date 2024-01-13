package main
import "fmt"

func main() {
  var a,b,c int
  fmt.Scan(&a)
  b = a + 32
  c = b + a
  fmt.Print(c)
}

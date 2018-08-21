package main
import(
  "fmt"
  "runtime"
)

func main(){
  switch os:=runtime.GOOS; os{
  case "darwin":
    fmt.Println("This is a MAC OS")
  case "linux":
fmt.Println("thisisLINUX")
default:

  fmt.Printf("%s",os)}
  os:=runtime.GOOS
  fmt.Println(os)
}

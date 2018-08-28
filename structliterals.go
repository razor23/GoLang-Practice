package main
import "fmt"

type guitar struct
{
  color string
  size int
}

var  (
  x1=guitar{color:"red"}
  x2=guitar{}
  x3=guitar{size:1}
)

func main(){
  fmt.Println(x1,x2,x3)
}

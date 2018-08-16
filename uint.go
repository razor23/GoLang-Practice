package main
import (
 "fmt"
 "math/cmplx"
)

var(
  bozeman bool=false
  Toby uint64=1<<32
  Max complex128=cmplx.Sqrt(-5+8i)

)

func main(){
  fmt.Printf("The value of %T is %v\n",bozeman,bozeman)
  fmt.Printf("The value of %T is %v\n",Toby,Toby)
  fmt.Printf("The value of %T is %v\n",Max,Max)


}

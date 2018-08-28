package main

import "fmt"

func main(){

fmt.Println("Hello Count")

for i:=0;i<10;i++{
  defer fmt.Println(i)
}

fmt.Println("Let's go")

}

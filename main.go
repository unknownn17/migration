package main

import "fmt"

func main(){
	msg:=SayHello("doston")
	fmt.Println(msg)
}

func SayHello(name string)(string){
	return fmt.Sprintf("Howdy + %s",name)
}
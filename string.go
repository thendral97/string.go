package main

import {"fmt"
		"os"
		"log"
		"io/ioutil"}
func main(){
	file,err:=os.Create("sample.tx")
	if err!=nil
		 log.fatal(err)
}
file writestring("Hi,my name is arrya and this file got created by go!!")
file.close()
steam,err:=ioutil.Readfile("Sample.txt")
if err!=nil{
	s1=string(stream)
	fmt.Println(s1)
}







}
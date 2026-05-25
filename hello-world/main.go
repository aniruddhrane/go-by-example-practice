package main
import "fmt"


func main(){
  fmt.Println("Hello Rane")
}
/*
structure of go program
1.package declaration
2.Imports
3.Functions
4.Statement inside functions

what is package?
A package is a collection of related Gofiles/code
similar to modules,namespaces,libraries in other languages
why main?
main is special package 
go searches for package main 
to find the executable program
If package was different it would have been resusable file

func main()
import "fmt"
It is go standard library for printing formatting text,input/output
more standard libraries fmt,math,strings,os,time,net/http

{ should be present infront of function name because go inserts semicolon

*/

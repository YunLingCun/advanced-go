package main

//// 1.1

//import "C"
//
//func main() {
//    println("hello cgo")
//}

//// 1.2
//
//// #include <stdio.h>
//import "C"
//
//func main() {
//    C.puts(C.CString("Hello World\n"))
//}

//// 1.3
//
///*
//#include <stdio.h>
//
//static void SayHello(const char*s) {
//    puts(s);
//}
// */
//import "C"
//
//func main() {
//    C.SayHello(C.CString("Hello World"))
//}

// 1.4

//void SayHello(const char* s);
import "C"

func main() {
    C.SayHello(C.CString("Hello World\n"))
}
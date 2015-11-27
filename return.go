package citest

/*
#include "Python.h"
int funcInt() {
  return 100;
}
*/
/*
#cgo linux pkg-config: python-2.7
#cgo darwin pkg-config: python-2.7
*/
import "C"

func returnInt() int {
	return int(C.funcInt())
}

func returnString() string {
	return "abc"
}

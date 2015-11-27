package citest

/*
int funcInt() {
  return 100;
}
*/
import "C"

func returnInt() int {
	return int(C.funcInt())
}

func returnString() string {
	return "abc"
}

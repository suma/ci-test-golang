package citest

func returnString() string {
	return "abc"
}

func errorShouldBeTheLastType() (error, string) {
	return nil, "abc"
}

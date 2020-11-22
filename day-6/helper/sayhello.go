package helper

//bisa di export
func SayHello(name string) string {
	return "Hello " + name
}

//tidak bisa di export
func sayHi(name string) string {
	return "Hi " + name
}

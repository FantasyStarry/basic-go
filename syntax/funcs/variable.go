package funcs

func YourName(name string, aliases ...string) {

}

func CallYourName() {
	YourName("zhangsan", "1")
	YourName("zhangsan", "1", "2")
	YourName("zhangsan", "1", "2", "3")
}

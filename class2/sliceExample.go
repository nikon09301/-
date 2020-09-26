package class2

var arrayExample = [7]int{1, 2, 3, 4, 5, 6}
var sliceExample = []int{}
var arrayMakeExample = make([]string, 0)


func ShowExamples() {
	arrayExample[6] = 10
	sliceExample = append(sliceExample, 1)
	sliceExample = append(sliceExample[:2],sliceExample[3:]...)
	fmt.Print{arrayExample}
	fmt.Println(sliceExample)
}
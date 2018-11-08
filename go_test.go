// array.go
package main

import (
	"fmt"
	//"math"
	//"strconv"
	//"os"
	//"net/http"
	//"runtime"
	//"time"
)

/*
type Student struct {
	id    int64
	name  string
	sex   int
	age   uint
	class string
}

func main() {
	var n [10]int
	var i, j int

	for i = 0; i < 10; i++ {
		n[i] = i + 100
	}

	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}
}
*/

/*
func main() {
	const (
		a = iota //0
		b        //1
		c        //2
		d = "ha" //独立值，iota += 1
		e        //"ha"   iota += 1
		f = 100  //iota +=1
		g        //100  iota +=1
		h = iota //7,恢复计数
		i        //8
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)
}
*/

//func main() {
//	var a int = 4
//	var b int32
//	var c float32
//	var ptr *int

//	/* 运算符实例 */
//	fmt.Printf("第 1 行 - a 变量类型为 = %T\n", a)
//	fmt.Printf("第 2 行 - b 变量类型为 = %T\n", b)
//	fmt.Printf("第 3 行 - c 变量类型为 = %T\n", c)

//	/*  & 和 * 运算符实例 */
//	ptr = &a /* 'ptr' 包含了 'a' 变量的地址 */
//	fmt.Printf("a 的值为  %d\n", a)
//	fmt.Printf("*ptr 为 %d\n", *ptr)
//	fmt.Printf("*ptr 为 0x%x\n", ptr)

//	if a == 4 {
//		fmt.Printf("a 的值为 : %d\n", a)
//	}
//}

//func main() {
//	/* 定义局部变量 */
//	var grade string = "B"
//	var marks int = 90

//	switch marks {
//	case 90:
//		grade = "A"
//	case 80:
//		grade = "B"
//	case 50, 60, 70:
//		grade = "C"
//	default:
//		grade = "D"
//	}

//	switch {
//	case grade == "A":
//		fmt.Printf("优秀!\n")
//	case grade == "B", grade == "C":
//		fmt.Printf("良好\n")
//	case grade == "D":
//		fmt.Printf("及格\n")
//	case grade == "F":
//		fmt.Printf("不及格\n")
//	default:
//		fmt.Printf("差\n")
//	}
//	fmt.Printf("你的等级是 %s\n", grade)
//}

//func main() {
//	var x interface{}

//	switch i := x.(type) {
//	case nil:
//		fmt.Printf(" x 的类型 :%T", i)
//	case int:
//		fmt.Printf("x 是 int 型")
//	case float64:
//		fmt.Printf("x 是 float64 型")
//	case func(int) float64:
//		fmt.Printf("x 是 func(int) 型")
//	case bool, string:
//		fmt.Printf("x 是 bool 或 string 型")
//	default:
//		fmt.Printf("未知型")
//	}
//}

//func main() {
//	var c1, c2, c3 chan int
//	var i1, i2 int
//	select {
//	case i1 = <-c1:
//		fmt.Printf("received ", i1, " from c1\n")
//	case c2 <- i2:
//		fmt.Printf("sent ", i2, " to c2\n")
//	case i3, ok := (<-c3): // same as: i3, ok := <-c3
//		if ok {
//			fmt.Printf("received ", i3, " from c3\n")
//		} else {
//			fmt.Printf("c3 is closed\n")
//		}
//	default:
//		fmt.Printf("no communication\n")
//	}
//}

//func main() {

//	var b int = 15
//	var a int

//	numbers := [6]int{1, 2, 3, 5}

//	/* for 循环 */
//	for a := 0; a < 10; a++ {
//		fmt.Printf("a 的值为: %d\n", a)
//	}

//	for a < b {
//		a++
//		fmt.Printf("a 的值为: %d\n", a)
//	}

//	for i, x := range numbers {
//		fmt.Printf("第 %d 位 x 的值 = %d %T\n", i, x, numbers)
//	}
//}

//func main() {
//	/* 定义局部变量 */
//	var a int = 100
//	var b int = 200
//	var ret int

//	/* 调用函数并返回最大值 */
//	ret = max(a, b)

//	fmt.Printf("最大值是 : %d\n", ret)
//}

///* 函数返回两个数的最大值 */
//func max(num1, num2 int) int {
//	/* 定义局部变量 */
//	var result int

//	if num1 > num2 {
//		result = num1
//	} else {
//		result = num2
//	}
//	return result
//}

//func swap(x, y string) (string, string) {
//	return y, x
//}

//func main() {
//	a, b := swap("Mahesh", "Kumar")
//	fmt.Println(a, b)
//}

//func main() {
//	/* 定义局部变量 */
//	var a int = 100
//	var b int = 200

//	fmt.Printf("交换前，a 的值 : %d\n", a)
//	fmt.Printf("交换前，b 的值 : %d\n", b)

//	/* 调用 swap() 函数
//	 * &a 指向 a 指针，a 变量的地址
//	 * &b 指向 b 指针，b 变量的地址
//	 */
//	swap(&a, &b)

//	fmt.Printf("交换后，a 的值 : %d\n", a)
//	fmt.Printf("交换后，b 的值 : %d\n", b)
//}

//func swap(x *int, y *int) {
//	var temp int
//	temp = *x /* 保存 x 地址上的值 */
//	*x = *y   /* 将 y 值赋给 x */
//	*y = temp /* 将 temp 值赋给 y */
//}

//func main() {
//	/* 声明函数变量 */
//	getSquareRoot := func(x float64) float64 {
//		return math.Sqrt(x)
//	}

//	/* 使用函数 */
//	fmt.Println(getSquareRoot(9))
//}

//func getSequence() func() int {
//	i := 0
//	return func() int {
//		i += 1
//		return i
//	}
//}

//func main() {
//	/* nextNumber 为一个函数，函数 i 为 0 */
//	nextNumber := getSequence()

//	/* 调用 nextNumber 函数，i 变量自增 1 并返回 */
//	fmt.Println(nextNumber())
//	fmt.Println(nextNumber())
//	fmt.Println(nextNumber())

//	/* 创建新的函数 nextNumber1，并查看结果 */
//	nextNumber1 := getSequence()
//	fmt.Println(nextNumber1())
//	fmt.Println(nextNumber1())
//}

//func main() {
//	add_func := add(1, 2)
//	fmt.Println(add_func(1, 1))
//	fmt.Println(add_func(0, 0))
//	fmt.Println(add_func(2, 2))
//}

//// 闭包使用方法
//func add(x1, x2 int) func(x3 int, x4 int) (int, int, int) {
//	i := 0
//	return func(x3 int, x4 int) (int, int, int) {
//		i++
//		return i, x1 + x2, x3 + x4
//	}
//}

///* 定义结构体 */
//type Circle struct {
//	radius float64
//}

//func main() {
//	var c1 Circle
//	c1.radius = 10.00
//	fmt.Println("Area of Circle(c1) = ", c1.getArea())
//}

////该 method 属于 Circle 类型对象中的方法
//func (c Circle) getArea() float64 {
//	//c.radius 即为 Circle 类型对象中的属性
//	return 3.14 * c.radius * c.radius
//}

//func add(a, b int) int {
//	return a + b
//}

//func multiplicationTable() {
//	for i := 1; i <= 9; i++ {
//		for j := 1; j <= i; j++ {
//			var ret string
//			if i*j < 10 && j != 1 {
//				ret = " " + strconv.Itoa(i*j)
//			} else {
//				ret = strconv.Itoa(i * j)
//			}

//			fmt.Print(j, "*", i, " =", ret, " ")
//		}

//		fmt.Print("\n")
//	}
//}

//func main() {
//	multiplicationTable()
//}

//var a int = 20

//func main() {
//	var a int = 10
//	var b int = 20
//	var c int = 0

//	fmt.Printf("main a = %d\n", a)
//	c = sum(a, b)
//	fmt.Printf("main c = %d\n", c)
//}

//func sum(a, b int) int {
//	fmt.Printf("sum a = %d\n", a)
//	fmt.Printf("sum b = %d\n", b)

//	return a + b
//}

//func main() {
//	var balance = []int{1000, 2, 3, 17, 50}
//	var avg float32

//	avg = getAverage(balance, 5)

//	fmt.Printf("平均值:%f\n", avg)
//}

//func getAverage(arr []int, size int) float32 {
//	var i, sum int
//	var avg float32

//	for i = 0; i < size; i++ {
//		sum += arr[i]
//	}

//	avg = float32(sum) / float32(size)

//	return avg
//}

//func main() {
//	var a int = 20 /* 声明实际变量 */
//	var ip *int    /* 声明指针变量 */

//	ip = &a /* 指针变量的存储地址 */

//	fmt.Printf("a 变量的地址是: %x\n", &a)

//	/* 指针变量的存储地址 */
//	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)

//	/* 使用指针访问值 */
//	fmt.Printf("*ip 变量的值: %d\n", *ip)
//}

//func main() {
//	var s, sep string
//	for i := 1; i < len(os.Args); i++ {
//		s += sep + os.Args[i]
//		sep = " "
//	}
//	fmt.Println(s)
//}

//type Books struct {
//	title   string
//	author  string
//	subject string
//	book_id int
//}

//func main() {
//	// 创建一个新的结构体
//	fmt.Println(Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})

//	// 也可以使用 key => value 格式
//	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", book_id: 6495407})

//	// 忽略的字段为 0 或 空
//	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com"})
//}

//type Books struct {
//	title   string
//	author  string
//	subject string
//	book_id int
//}

//func main() {
//	var Book1 Books /* 声明 Book1 为 Books 类型 */
//	var Book2 Books /* 声明 Book2 为 Books 类型 */

//	/* book 1 描述 */
//	Book1.title = "Go 语言"
//	Book1.author = "www.runoob.com"
//	Book1.subject = "Go 语言教程"
//	Book1.book_id = 6495407

//	/* book 2 描述 */
//	Book2.title = "Python 教程"
//	Book2.author = "www.runoob.com"
//	Book2.subject = "Python 语言教程"
//	Book2.book_id = 6495700

//	/* 打印 Book1 信息 */
//	printBook(&Book1)

//	/* 打印 Book2 信息 */
//	printBook(&Book2)
//}

//func printBook(book *Books) {
//	fmt.Printf("Book title : %s\n", book.title)
//	fmt.Printf("Book author : %s\n", book.author)
//	fmt.Printf("Book subject : %s\n", book.subject)
//	fmt.Printf("Book book_id : %d\n", book.book_id)
//}

//func main() {
//	var numbers = make([]int, 3, 5)
//	printSlice(numbers)
//}

//func printSlice(x []int) {
//	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
//}

//func main() {
//	var numbers []int
//	printSlice(numbers)

//	/* 允许追加空切片 */
//	numbers = append(numbers, 0)
//	printSlice(numbers)

//	/* 向切片添加一个元素 */
//	numbers = append(numbers, 1)
//	printSlice(numbers)

//	/* 同时添加多个元素 */
//	numbers = append(numbers, 2, 3, 4)
//	printSlice(numbers)

//	/* 创建切片 numbers1 是之前切片的两倍容量*/
//	numbers1 := make([]int, len(numbers), (cap(numbers))*2)

//	/* 拷贝 numbers 的内容到 numbers1 */
//	copy(numbers1, numbers)
//	printSlice(numbers1)
//}

//func printSlice(x []int) {
//	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
//}

//重要
//func main() {
//	//这是我们使用range去求一个slice的和。使用数组跟这个很类似
//	nums := []int{2, 3, 4}
//	sum := 0
//	for _, num := range nums {
//		sum += num
//	}
//	fmt.Println("sum:", sum)
//	//在数组上使用range将传入index和值两个变量。上面那个例子我们不需要使用该元素的序号，所以我们使用空白符"_"省略了。有时侯我们确实需要知道它的索引。
//	for i, num := range nums {
//		if num == 3 {
//			fmt.Println("index:", i)
//		}
//	}
//	//range也可以用在map的键值对上。
//	kvs := map[string]string{"a": "apple", "b": "banana"}
//	for k, v := range kvs {
//		fmt.Printf("%s -> %s\n", k, v)
//	}
//	//range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
//	for i, c := range "go" {
//		fmt.Println(i, c)
//	}
//}

//func main() {
//	var countryCapitalMap map[string]string /*创建集合 */
//	countryCapitalMap = make(map[string]string)

//	/* map插入key - value对,各个国家对应的首都 */
//	countryCapitalMap["France"] = "Paris"
//	countryCapitalMap["Italy"] = "罗马"
//	countryCapitalMap["Japan"] = "东京"
//	countryCapitalMap["India "] = "新德里"

//	/*使用键输出地图值 */
//	for country := range countryCapitalMap {
//		fmt.Println(country, "首都是", countryCapitalMap[country])
//	}

//	/*查看元素在集合中是否存在 */
//	captial, ok := countryCapitalMap["美国"] /*如果确定是真实的,则存在,否则不存在 */
//	/*fmt.Println(captial) */
//	/*fmt.Println(ok) */
//	if ok {
//		fmt.Println("美国的首都是", captial)
//	} else {
//		fmt.Println("美国的首都不存在")
//	}
//}

//type Phone interface {
//	call()
//}

//type NokiaPhone struct {
//}

//func (nokiaPhone NokiaPhone) call() {
//	fmt.Println("I am Nokia, I can call you!")
//}

//type IPhone struct {
//}

//func (iPhone IPhone) call() {
//	fmt.Println("I am iPhone, I can call you!")
//}

//func main() {
//	var phone Phone

//	phone = new(NokiaPhone)
//	phone.call()

//	phone = new(IPhone)
//	phone.call()
//}

//go 接口
//type Man interface {
//	name() string
//	age() int
//}

//type Woman struct {
//}

//func (woman Woman) name() string {
//	return "Jin Yawei"
//}
//func (woman Woman) age() int {
//	return 23
//}

//type Men struct {
//}

//func (men Men) name() string {
//	return "liweibin"
//}
//func (men Men) age() int {
//	return 27
//}

//func main() {
//	var man Man

//	man = new(Woman)
//	fmt.Println(man.name())
//	fmt.Println(man.age())
//	man = new(Men)
//	fmt.Println(man.name())
//	fmt.Println(man.age())
//}

//type helloHandler struct{}

//func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	w.Write([]byte("Hello, world!"))
//}

//func main() {
//	http.Handle("/", &helloHandler{})
//	http.ListenAndServe(":12345", nil)
//}

//func main() {
//	c1 := make(chan string, 1)
//	go func() {
//		time.Sleep(time.Second * 2)
//		c1 <- "result 1"
//	}()
//	select {
//	case res := <-c1:
//		fmt.Println(res)
//	case <-time.After(time.Second * 1):
//		fmt.Println("timeout 1")
//	}
//}

//func main() {
//	str := "hello,世界"
//	n := len(str)
//	for i := 0; i < n; i++ {
//		ch := str[i]
//		fmt.Println(i, ch)
//	}
//}

//func main() {
//	mySlice := make([]int, 5, 10)
//	fmt.Println("len(mySlice):", len(mySlice))
//	fmt.Println("cap(mySlice):", cap(mySlice))

//	mySlice = append(mySlice, 1, 2, 3)
//}

//type PersonInfo struct {
//	ID      string
//	Name    string
//	Address string
//}

//func main() {
//	//var personDB map[string]PersonInfo
//	personDB := make(map[string]PersonInfo)

//	personDB["12345"] = PersonInfo{"12345", "Tom", "Room 203,..."}
//	personDB["1"] = PersonInfo{"1", "Jack", "Room 101,..."}

//	person, ok := personDB["1234"]

//	if ok {
//		fmt.Println("Found person", person.Name, "WITH ID 1234.")
//	} else {
//		fmt.Println("Did not find person with ID 1234.")
//	}

//	for i, v := range personDB {
//		fmt.Println("key:", i, "value:", v)
//	}
//}

//Go语言中的闭包
//func main() {
//	var j int = 5
//	a := func() func() {
//		var i int = 10
//		return func() {
//			fmt.Printf("i, j: %d, %d\n", i, j)
//		}
//	}()

//	a()

//	j *= 2

//	a()
//}

//func main() {
//	r := [...]int{99: -1}

//	for i := 0; i < len(r); i++ {
//		fmt.Printf("%d ", r[i])
//		if i%10 == 0 {
//			fmt.Printf("\n")
//		}
//	}
//}

//go 方法
//type Point struct{ X, Y float64 }

////	traditional	function
//func Distance(p, q Point) float64 {
//	return math.Hypot(q.X-p.X, q.Y-p.Y)
//}

////	same thing,	but	as a method	of the Point type
//func (p Point) Distance(q Point) float64 {
//	return math.Hypot(q.X-p.X, q.Y-p.Y)
//}

//type Rectangle struct {
//	width, height float64
//}

//type Circle struct {
//	radius float64
//}

//func (r Rectangle) area() float64 {
//	return r.width * r.height
//}

//func (c Circle) area() float64 {
//	return c.radius * c.radius * math.Pi
//}

//func main() {
//	r1 := Rectangle{12, 2}
//	r2 := Rectangle{9, 4}
//	c1 := Circle{10}
//	c2 := Circle{25}

//	fmt.Println("Area of r1 is: ", r1.area())
//	fmt.Println("Area of r2 is: ", r2.area())
//	fmt.Println("Area of c1 is: ", c1.area())
//	fmt.Println("Area of c2 is: ", c2.area())
//}

//const (
//	WHITE = iota
//	BLACK
//	BLUE
//	RED
//	YELLOW
//)

//type Color byte

//type Box struct {
//	width, height, depth float64
//	color                Color
//}

//type BoxList []Box //a slice of boxes

//func (b Box) Volume() float64 {
//	return b.width * b.height * b.depth
//}

//func (b *Box) SetColor(c Color) {
//	b.color = c
//}

//func (bl BoxList) BiggestColor() Color {
//	v := 0.00
//	k := Color(WHITE)
//	for _, b := range bl {
//		if bv := b.Volume(); bv > v {
//			v = bv
//			k = b.color
//		}
//	}
//	return k
//}

//func (bl BoxList) PaintItBlack() {
//	for i, _ := range bl {
//		bl[i].SetColor(BLACK)
//	}
//}

//func (c Color) String() string {
//	strings := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
//	return strings[c]
//}

//func main() {
//	boxes := BoxList{
//		Box{4, 4, 4, RED},
//		Box{10, 10, 1, YELLOW},
//		Box{1, 1, 20, BLACK},
//		Box{10, 10, 1, BLUE},
//		Box{10, 30, 1, WHITE},
//		Box{20, 20, 20, YELLOW},
//	}

//	fmt.Printf("We have %d boxes in our set\n", len(boxes))
//	fmt.Println("The volume of the first one is", boxes[0].Volume(), "cm³")
//	fmt.Println("The color of the last one is", boxes[len(boxes)-1].color.String())
//	fmt.Println("The biggest one is", boxes.BiggestColor().String())

//	fmt.Println("Let's paint them all black")
//	boxes.PaintItBlack()
//	fmt.Println("The color of the second one is", boxes[1].color.String())

//	fmt.Println("Obviously, now, the biggest one is", boxes.BiggestColor().String())
//}

//method继承
type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human  //匿名字段
	school string
}

type Employee struct {
	Human   //匿名字段
	company string
}

//在human上面定义了一个method
func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

func main() {
	mark := Student{Human{"Mark", 25, "222-222-YYYY"}, "MIT"}
	sam := Employee{Human{"Sam", 45, "111-888-XXXX"}, "Golang Inc"}

	mark.SayHi()
	sam.SayHi()
}

////method重写
//type Human struct {
//	name  string
//	age   int
//	phone string
//}

//type Student struct {
//	Human  //匿名字段
//	school string
//}

//type Employee struct {
//	Human   //匿名字段
//	company string
//}

////Human定义method
//func (h *Human) SayHi() {
//	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
//}

////Employee的method重写Human的method
//func (e *Employee) SayHi() {
//	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
//		e.company, e.phone) //Yes you can split into 2 lines here.
//}

//func main() {
//	mark := Student{Human{"Mark", 25, "222-222-YYYY"}, "MIT"}
//	sam := Employee{Human{"Sam", 45, "111-888-XXXX"}, "Golang Inc"}

//	mark.SayHi()
//	sam.SayHi()
//}

//goroutine
//func newTask() {
//	i := 0
//	for {
//		i++
//		fmt.Printf("new goroutine: i = %d\n", i)
//		time.Sleep(1 * time.Second)
//	}
//}

//func main() {
//	go newTask()

//	j := 0

//	for {
//		j++
//		fmt.Printf("main goroutine: j = %d\n", j)
//		time.Sleep(1 * time.Second)
//	}
//}

//func main() {
//	go func(s string) {
//		for i := 0; i < 2; i++ {
//			fmt.Println(s)
//		}
//	}("world")

//	for i := 0; i < 2; i++ {
//		runtime.Gosched()
//		fmt.Println("hello")
//	}
//}

//func main() {
//	c := make(chan int)

//	go func() {
//		defer fmt.Println("子协程结束")

//		fmt.Println("子协程正在运行……")

//		c <- 666
//	}()

//	num := <-c

//	fmt.Println("num =", num)
//	fmt.Println("main协程结束")
//}

//人
//type Person struct {
//	name string
//	sex  byte
//	age  int
//}

////学生
//type Student struct {
//	Person // 匿名字段，那么默认Student就包含了Person的所有字段
//	id     int
//	addr   string
//	name   string //和Person中的name同名
//}

//func main() {
//	var s Student //变量声明

//	//给Student的name，还是给Person赋值？
//	s.name = "mike"

//	//{Person:{name: sex:0 age:0} id:0 addr: name:mike}
//	fmt.Printf("%+v\n", s)

//	//默认只会给最外层的成员赋值
//	//给匿名同名成员赋值，需要显示调用
//	s.Person.name = "yoyo"
//	//Person:{name:yoyo sex:0 age:0} id:0 addr: name:mike}
//	fmt.Printf("%+v\n", s)
//}

//结构体指针类型
//type Person struct { //人
//	name string
//	sex  byte
//	age  int
//}

//type Student struct { //学生
//	*Person // 匿名字段，结构体指针类型
//	id      int
//	addr    string
//}

//func main() {
//	//初始化
//	s1 := Student{&Person{"mike", 'm', 18}, 1, "bj"}

//	//{Person:0xc0420023e0 id:1 addr:bj}
//	fmt.Printf("%+v\n", s1)
//	//mike, m, 18
//	fmt.Printf("%s, %c, %d\n", s1.name, s1.sex, s1.age)

//	//声明变量
//	var s2 Student
//	s2.Person = new(Person) //分配空间
//	s2.name = "yoyo"
//	s2.sex = 'f'
//	s2.age = 20

//	s2.id = 2
//	s2.addr = "sz"

//	//yoyo 102 20 2 20
//	fmt.Println(s2.name, s2.sex, s2.age, s2.id, s2.age)
//}

//函数类型
//type FuncType func(int, int) int //声明一个函数类型, func后面没有函数名

////函数中有一个参数类型为函数类型：f FuncType
//func Calc(a, b int, f FuncType) (result int) {
//	result = f(a, b) //通过调用f()实现任务
//	return
//}

//func Add(a, b int) int {
//	return a + b
//}

//func Minus(a, b int) int {
//	return a - b
//}

//func main() {
//	//函数调用，第三个参数为函数名字，此函数的参数，返回值必须和FuncType类型一致
//	result := Calc(1, 1, Add)
//	fmt.Println(result) //2

//	var f FuncType = Minus
//	fmt.Println("result = ", f(10, 2)) //result =  8
//}

//字符串操作
//func Contains(s, substr string) bool
//功能：字符串s中是否包含substr，返回bool值

//func Join(a []string, sep string) string
//功能：字符串链接，把slice a通过sep链接起来

//func Index(s, sep string) int
//功能：在字符串s中查找sep所在的位置，返回位置值，找不到返回-1

//func Repeat(s string, count int) string
//功能：重复s字符串count次，最后返回重复的字符串

//func Replace(s, old, new string, n int) string
//功能：在s字符串中，把old字符串替换为new字符串，n表示替换的次数，小于0表示全部替换

//func Split(s, sep string) []string
//功能：把s字符串按照sep分割，返回slice

//func Trim(s string, cutset string) string
//功能：在s字符串的头部和尾部去除cutset指定的字符串

//func Fields(s string) []string
//功能：去除s字符串的空格符，并且按照空格分割返回slice

//func main() {
//	fmt.Println(strings.Contains("seafood", "foo"))
//	fmt.Println(strings.Contains("seafood", "bar"))
//	fmt.Println(strings.Contains("seafood", ""))
//	fmt.Println(strings.Contains("", ""))
//	//运行结果:
//	//true
//	//false
//	//true
//	//true

//	s := []string{"foo", "bar", "baz"}
//	fmt.Println(strings.Join(s, ", "))
//	//运行结果:foo, bar, baz

//	fmt.Println(strings.Index("chicken", "ken"))
//	fmt.Println(strings.Index("chicken", "dmr"))
//	//运行结果:
//	//    4
//	//    -1

//	fmt.Println("ba" + strings.Repeat("na", 2))
//	//运行结果:banana

//	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
//	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))
//	//运行结果:
//	//oinky oinky oink
//	//moo moo moo

//	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
//	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
//	fmt.Printf("%q\n", strings.Split(" xyz ", ""))
//	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))
//	//运行结果:
//	//["a" "b" "c"]
//	//["" "man " "plan " "canal panama"]
//	//[" " "x" "y" "z" " "]
//	//[""]

//	fmt.Printf("[%q]", strings.Trim(" !!! Achtung !!! ", "! "))
//	//运行结果:["Achtung"]

//	fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   "))
//	//运行结果:Fields are: ["foo" "bar" "baz"]
//}


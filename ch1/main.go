package main

func main() {
	// a := []int{1,2,3,4,5}
	// fmt.Println(a)
	// fmt.Println(len(a))

	// a[4] = 99
	// fmt.Println(a)

	// fmt.Println(a[0:2])
	// fmt.Println(a[:3])
	// fmt.Println(a[:len(a) - 1])
	// fmt.Println(a[:len(a) - 2])

	// a := [3]int{1,2,3}
	// fmt.Printf("%T\n", a)
	// fmt.Printf("%T\n", a[0:])

	// me := map[string]int{"height": 180}
	// fmt.Println(me["height"])

	// me["weight"] = 70
	// fmt.Println(me)

	// hungry := true
	// sleepy := false

	// fmt.Println(!hungry)
	// fmt.Println(hungry && sleepy)
	// fmt.Println(hungry || sleepy)

	// for _, v := range []int{1,2,3}{
	// 	fmt.Println(v)
	// }

	m := NewMan("David")
	m.Hello()
	m.Goodbye()
}

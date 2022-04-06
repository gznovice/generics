package main


import(
	"fmt"
)

/*
1. usage of any type
2. print info ot cerntain type 
3. relationship with interface
*/

type animal interface{
	//func say() //can't have func
	say()
}

type people struct{
}

func(p people) say(){
	fmt.Println("HI")
}

type putin struct{
}

func(p putin) say(){
	fmt.Println("KILL")
}

//type constraint 
type Addable interface{
	~int| ~int8| ~int16| ~int32| ~int64| ~uint| ~uint8| ~uint16| ~uint32| ~uint64| ~string
}

func add[T Addable](a, b T) T{
	return a + b
}

//func printType[T Addable](a T) {//.\TestGenerics.go:45:14: cannot use type switch on type parameter value a (variable of type T constrained by Addable)
func printType(a interface{}) {
	switch  a.(type) {
    case int, int8, int16, int32, int64:
        fmt.Println("int")
    case uint, uint8, uint16, uint32, uint64:
        fmt.Println("uint")
    case string:
        fmt.Println("string")
    }
}

func main(){
	myObjectSlice := []animal{
		people{}, putin{}}
	
	//for v := range myObjectSlice { ///omit the _
	for _, v := range myObjectSlice {
		v.say()
	}
	
	
	myObjectSlice2 := []any{//any type support since 1.18...
		people{}, putin{}}
	
	//for v := range myObjectSlice { ///omit the _
	for _, v := range myObjectSlice2 {		
		//v.say() v.say undefined (type any has no field or method say)
		v = v
	}
		
	fmt.Println("add(1, 3) =", add(1, 3))
	fmt.Println("add(\"hello\", \"world\") =", add("hello", "world"))
	//fmt.Println("add(1, \"world\") =", add(1, "world")) default type string of "world" does not match inferred type int for T
	//fmt.Println("add(1.1, 3.3) =", add(1.1, 3.3))float64 does not implement Addable
	printType(3)
	printType(uint(3))
	printType("hello")
}
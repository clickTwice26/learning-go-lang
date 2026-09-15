package main
import "fmt"
// maps =>  associative data structure
func main(){

	// creating map
	m := make(map[string]string)

	// setting an element
	m["name"] = "go-lang"

	// get an element
	fmt.Println(m["name"])

	// add more element
	m["area"] = "backend"

	// accssing non existing key
	fmt.Println(m["phone"]) // it does not throw an instant error, return empty(nil)
	m1 := make(map[string]int)
	fmt.Println(m1["test"])

	//len
	fmt.Println(len(m))
	delete(m, "area") // delete a key
	fmt.Println(m)
	clear(m) // clear the entire map

	m2 := map[string]int{
		"price" : 10,
		"discount" : 5,
	}
	fmt.Println(m2)
	_k, ok := m2["price"]
	if ok{
		fmt.Println(_k, ok)
	}else{
		fmt.Println("key not found", ok)
	}
}
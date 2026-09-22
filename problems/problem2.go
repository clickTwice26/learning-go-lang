package main
import "fmt"
func main(){
	price_catalog := map[string]float32{
		"apple": 0.99,
    	"banana": 0.59,
     	"milk": 3.49,
      	"bread": 2.50,
       	"eggs": 4.25,
	}
	var sum float32 = 0.0
	notExistProducts := []string{}
	customer_cart := []string{"apple", "banana", "apple", "chocolate", "milk", "coffee"}
	fmt.Println(price_catalog)

	for _, value := range customer_cart {
		if price, isExist := price_catalog[value]; isExist {
			// fmt.Println(price, isExist)
			sum += price
		}else{
			// fmt.Println(value, "is not exist in the catalog")
			notExistProducts = append(notExistProducts, value)
		}
	}

	fmt.Printf("Total Bill : $%.2f\n", sum)
	fmt.Println("Unstocked Items : ", notExistProducts)
}
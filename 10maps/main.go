package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")

	Currencies := make(map[string]string)
	Currencies["Rs"] = "India"
	Currencies["AED"] = "UAE"
	Currencies["Ryl"] = "KSA"
	Currencies["Ut"] = "B"

	fmt.Println("Currencies and Countries: ", Currencies)
	fmt.Println("Currencies['AED'] : " , Currencies["AED"])

	delete(Currencies, "Ut")
	fmt.Println("Map: ", Currencies)

	for key, value := range Currencies {
		fmt.Printf("key: %v -> value: %v\n",key,value)
	}

	for _, value := range Currencies {
		fmt.Printf("%v \t \n",value)
	}

	for i := range 5{
		fmt.Print(i, " ")
	}
}
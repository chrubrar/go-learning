package datastructures

import (
	"fmt"
	"sort"
)

func main() {

	Salary := make(map[string]int)

	Salary["Jennifer"] = 50000
	Salary["Chru"] = 250000
	Salary["ben"] = 1000

	fmt.Println("Salary map contents=", Salary)

	m := map[string]string{"route": "66"}

	p := m["route"]
	fmt.Println("p is =", p)

	q := m["rute"]

	fmt.Println("q is =", q)

	//p, ok := m["route"]

	//_, ok := m["route"]

	powderpuff := map[string]string{"name": "powderpuff", "breed": "terrier", "color": "white/brown", "owner": "Taig"}

	fmt.Printf(" The map is = %v", powderpuff)

	for key, value := range powderpuff {
		fmt.Println("value=", value, "key=", key)
	}

	fluffythings, ok := powderpuff["owner"]

	if ok {
		fmt.Println("fluffythings=", fluffythings)
	} else {
		fmt.Println("fluffythings not found")
	}

	keysOnly := []string{}

	for key := range powderpuff {
		keysOnly = append(keysOnly, key)
	}

	fmt.Println("keysOnly=", keysOnly)

	powderpuff["Vet"] = "Dr. Dolittle"

	fmt.Println("powderpuff=", powderpuff)

	usernames := map[string]string{
		"ben":  "bendysingh",
		"joe":  "joeysingh",
		"taig": "taigbrar",
		"jess": "jessbrar",
		"chru": "chrubrar",
	}

	fmt.Println("usernames=", usernames)

	delete(usernames, "ben")

	fmt.Println("usernames=", usernames)
	usernames["Tanya"] = "Tanyabrar"

	fmt.Println("usernames=", usernames)

	salesClosed := map[string]string{
		"ben":  "2000",
		"joe":  "1500",
		"taig": "1000",
		"jess": "500",
		"chru": "250",
	}

	fmt.Println("salesClosed=", salesClosed)

	for key, value := range salesClosed {
		fmt.Println("value=", value, "key=", key)
	}

	salesClosed["chru"] = "25000"

	fmt.Println("salesClosed=", salesClosed)

	keys := make([]string, 0, len(salesClosed))

	for k := range salesClosed {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, salesClosed[k])
	}
}

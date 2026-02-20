package main

import (
	"fmt"
	"maps"
)

func mapAsSet() {
	intSet := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}

	for _, v := range vals {
		intSet[v] = true
	}
	fmt.Println(intSet)
	fmt.Println(len(vals), len(intSet))
	fmt.Println(intSet[5])
	fmt.Println(intSet[500])
	if intSet[100] {
		fmt.Println("100 is in the set")
	}
}


func Units() map[string]int {
	units := map[string]int{
        "quarter_of_a_dozen":3,
        "half_of_a_dozen":6,
        "dozen":12,
        "small_gross":120,
        "gross":144,
    	"great_gross":1728,
    }

    return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {                                          
      val, ok := units[unit]                                                                                  
      if !ok {                                                                                                
          return false                                                                                        
      }                                                                                                       
      bill[item] += val                                                                                       
      return true                                                                                             
  } 

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {                                       
      currentQty, exists := bill[item]                                                                        
      if !exists {                                                                                            
          return false                                                                                        
      }                                                                                                       
                                                                                                              
      unitValue, ok := units[unit]                                                                            
      if !ok {                                                                                                
          return false                                                                                        
      }                                                                                                       
                                                                                                              
      newQty := currentQty - unitValue                                                                        
      if newQty < 0 {                                                                                         
          return false                                                                                        
      }                                                                                                       
                                                                                                              
      if newQty == 0 {                                                                                        
          delete(bill, item)                                                                                  
      } else {                                                                                                
          bill[item] = newQty                                                                                 
      }                                                                                                       
                                                                                                              
      return true                                                                                             
  } 

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	itemValue,ok := bill[item]

    return itemValue,ok
}

func main() {
	var nilMap map[string]int //  nil map can be read but panics on write
	// nilMap["two"] = 2 //will cause panic
	totalWins := map[string]int{}
	totalWins["player_one"] = 111 // can read and write to a map assigned an empty map literal.
	schoolInfo := map[string][]string{
		"teachers": {"Zakir", "Asif"},
		"students": {"labib", "yousha"},
	}
	fmt.Println(nilMap)
	fmt.Println(totalWins)
	fmt.Println(schoolInfo)
	fmt.Println(len(schoolInfo))
	totalWins["Orcas"] = 1
	totalWins["Lions"] = 2
	fmt.Println(totalWins["Orcas"])
	fmt.Println(totalWins["Kittens"])
	totalWins["Kittens"]++
	fmt.Println(totalWins["Kittens"])
	totalWins["Lions"] = 3
	fmt.Println(totalWins["Lions"])

	m := map[string]int{
		"hello": 5,
		"world": 0,
		"code":  9,
	}

	//The comma ok Idiom - to check if a key is present in the map
	v, ok := m["hello"]
	fmt.Println(v, ok) // 5 true
	v, ok = m["goodbye"]
	fmt.Println(v, ok) //0 false

	// Deleting from Maps
	delete(m, "hello")
	fmt.Println(m)

	fmt.Println(m, len(m))
	clear(m) //Emptying a Map
	fmt.Println(m, len(m))

	m = map[string]int{
		"hello": 5,
		"world": 10,
	}

	n := map[string]int{
		"world": 10,
		"hello": 5,
	}

	fmt.Println(maps.Equal(m, n)) //Comparing Maps

	mapAsSet()
	fmt.Println("----------")
	
}

//idea for algorithm:
//For each element in nums, theres exists only 1 other number which can add to the target
//This value is currElement - target
//The niave approach must iterate through nums again looking for currElement-target making it O(n^2)
//Create a hashmap where key is value which stores the index in nums O(1) lookup

//Lets try a simple approach first and using  arrays and modulo
//Store each value in nums at index value%lengthOfHashTable

//first idea create array of length of nums*1.3 rounded up to nearest int
//store each element of nums in the hashTable at index = it's value%size of hashmap
//store collisions at next index
//check for exmpty slot by checking of value in hashTable is nil or 0

//ran into problem where init an empty array assigns empty values as 0
//but 0 is a valid value to encounter in nums
//Go is statically typed cannot mix int and nil values (used to dynamic langs like python)
//cloud create a new type struct but could get too complex for this problem

//Lets try this using built in maps type
//we can use multiple assignment with map[key] to store the value and a boolean if its present https://go.dev/doc/effective_go#maps
//Don't need to populate the hashMap before, lets just store the value as we check it

package twoSum_hashTable

func TwoSum(nums []int, target int) []int {
	hashTable := make(map[int]int)
	for currIndex, currValue := range nums {
		if checkIndex, present := hashTable[target-currValue]; present { //“comma ok” idiom
			return []int{checkIndex, currIndex}
		}
		hashTable[currValue] = currIndex
	}
	return []int{}
}

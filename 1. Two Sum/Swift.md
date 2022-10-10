$O(n^2)$
```Swift
func twoSum(_ nums: [Int], _ target: Int) -> [Int] {
    // Iterate through every element in numbers
    for (i, number1) in nums.enumerated() {
        /**
        For every iteration, 
        iterate through numbers,
        where i does not equal j.
        */
        for (j, number2) in nums.enumerated() where i != j {
   
            // Get the sum of number1 and number2
            let sum = number1 + number2
        
            // Check if sum equals the target
            if sum == target {
      
              // If so, return the indicies or numbers
              return [i, j] // or [number1, number2]
            }
        }
    }
    
    /**
    Return an empty array if unable to
    find a pair that adds up to the target.
    */
    return []
}
```
```Swift
func twoSum(_ nums: [Int], _ target: Int) -> [Int] {
    /**
    Init an empty dictionary.
    The value will be the array element.
    The key will be the index of the element.
    */
    var hashMap: [Int: Int] = [:]
    // Iterate through every element in numbers
    for (i, element) in numbers.enumerated() {
    
        /**
        Check if the hashMap contains an element
        that with the current element, sums up to k.
        */
        if let mapped = hashMap[target - element] {
           return [i, mapped]
        }
        
        // Add the element to the dictionary
        hashMap[element]
    }
    
    /**
    Return an empty array if unable to
    find a pair that adds up to the target.
    */
    return []
}
```
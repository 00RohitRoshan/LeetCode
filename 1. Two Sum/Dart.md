```Dart
class Solution {
  List<int> twoSum(List<int> nums, int target) {

   List<int> indices  = [];

  
  for(int i = 0; i< nums.length; i++){
    if(i != nums.length -1){
        for(int j=i+1;j< nums.length;j++){
            if(j < nums.length){
                if(nums[i] + nums[j] == target){
                 indices = [i,j];
                }
            }
        }
    }
  }
  return indices;    
  }
}
```
```Dart
class Solution {
  List<int> twoSum(List<int> nums, int target) {

  if (nums.isEmpty ||nums == null || target == null) return [];

  var map = {};
  List<int> result = [];

  for (int number in nums) {
    if (map.containsKey(number)) {
      result.add(map[number]);
      result.add(nums.indexOf(number));
    } else {
      map[target - number] = nums.indexOf(number);
    }
  }

  return result;    
  }
}
```
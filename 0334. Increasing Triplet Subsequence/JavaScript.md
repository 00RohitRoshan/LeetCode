```JS
var increasingTriplet = function(nums) {
    var f = Number.MAX_SAFE_INTEGER;
    var s = Number.MAX_SAFE_INTEGER;
    for (num in nums){
        if(num <= f){
            f = num ;
        }else if (num <=s){
            s = num ;
        }else if(num > f &&um > s){
            return true;
        }
    }
    return false;
};
```
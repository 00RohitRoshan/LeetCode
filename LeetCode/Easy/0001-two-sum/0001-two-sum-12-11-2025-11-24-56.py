class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        has = {}
        for i,v in enumerate(nums) :
            required = target-v
            if required in has :
                return [has[required],i]
            has[v] = i
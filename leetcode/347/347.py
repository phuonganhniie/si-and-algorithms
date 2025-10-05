class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        bucket = [[] for _ in range(len(nums) + 1)]
        frequencyMap = {}
        
        for n in nums:
            if n not in frequencyMap:
                frequencyMap[n] = 1
            else:
                frequencyMap[n] += 1
                
        for key, frequency in frequencyMap.items():
            bucket[frequency].append(key)
            
        result = []
        
        for i in reversed(range(len(bucket))):
            if bucket[i]:
                for value in bucket[i]:
                    if len(result) < k:
                        result.append(value)
                    else:
                        return result
        
        return result
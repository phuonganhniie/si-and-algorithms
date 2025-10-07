from typing import List

class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        track_map = {}
        longest_len = 0
        
        for num in nums:
            track_map[num] = True
            
        for num in track_map:
            current_length = 1

            next_num = num + 1
            previous_num = num - 1
            
            if previous_num not in track_map:
                while next_num in track_map:
                    current_length += 1
                    next_num += 1
                
                longest_len = max(current_length, longest_len)
                
        return longest_len    
from typing import List

"""
    2125. Number of Laser Beams in a Bank
    https://leetcode.com/problems/number-of-laser-beams-in-a-bank/description/
"""
class Solution:
    def numberOfBeams(self, bank: List[str]) -> int:
        ans = prev = 0
        
        for s in bank:
            c = s.count('1')
            if c:
                ans += prev * c
                prev = c

        return ans
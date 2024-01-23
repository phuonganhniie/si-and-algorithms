"""
    455. Assign Cookies
    https://leetcode.com/problems/assign-cookies/description/
"""

from typing import List

class Solution:
    def findContentChildren(self, g: List[int], s: List[int]) -> int:
        g.sort() # Sort the greed factors
        s.sort() # Sort the cookie sizes
        child_i = 0 # Initialize pointer for children
        cookie_j = 0 # Initialize pointer for cookies
        content_child = 0 # Initialize count of content children
        
        while child_i < len(g) and cookie_j < len(s):
            if s[cookie_j] >= g[child_i]:
                # The current cookie can content the child
                content_child += 1
                child_i += 1 # Move to the next child
            cookie_j += 1 # Move to the next cookie, regardless of whether the current child was content
        
        return content_child
class Solution:
    def isSubsequence(self, s: str, t: str) -> bool:
        ptrS, ptrT = 0, 0
        
        while ptrS < len(s) and ptrT < len(t):
            if s[ptrS] == t[ptrT]:
                ptrS += 1
                ptrT += 1
            else:
                ptrT += 1
        
        if len(s) == ptrS:
            return True
        return False
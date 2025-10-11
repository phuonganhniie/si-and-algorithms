from typing import List

class Solution:
    def maxPartitionFactor(self, points: List[List[int]]) -> int:
        n = len(points)
        if n == 2:
            return 0
        
        distances = []
        for i in range(n):
            for j in range(i + 1, n):
                dist = abs(points[i][0] - points[j][0]) + abs(points[i][1] - points[j][1])
                distances.append(dist)
        
        distances = sorted(set(distances))
        
        left, right = 0, len(distances) - 1
        result = 0
        
        while left <= right:
            mid = (left + right) // 2
            candidate = distances[mid]
            
            if self._can_achieve(points, candidate):
                result = candidate
                left = mid + 1
            else:
                right = mid - 1
        
        return result
    
    def _can_achieve(self, points: List[List[int]], min_factor: int) -> bool:
        n = len(points)
        graph = [[] for _ in range(n)]
        
        for i in range(n):
            for j in range(i + 1, n):
                dist = abs(points[i][0] - points[j][0]) + abs(points[i][1] - points[j][1])
                if dist < min_factor:
                    graph[i].append(j)
                    graph[j].append(i)
        
        color = [-1] * n
        
        def bfs(start):
            from collections import deque
            queue = deque([start])
            color[start] = 0
            
            while queue:
                node = queue.popleft()
                for neighbor in graph[node]:
                    if color[neighbor] == -1:
                        color[neighbor] = 1 - color[node]
                        queue.append(neighbor)
                    elif color[neighbor] == color[node]:
                        return False
            return True
        
        for i in range(n):
            if color[i] == -1:
                if not bfs(i):
                    return False
        
        has_edge = any(len(graph[i]) > 0 for i in range(n))
        if not has_edge:
            color[0] = 1
        
        count0 = sum(1 for c in color if c == 0)
        count1 = sum(1 for c in color if c == 1)
        
        return count0 > 0 and count1 > 0


def test_large_case():
    solution = Solution()
    points1 = [[3874,-65073],[17744,-88825],[51615,-6352],[64357,-9827],[97388,-493],[7489,-47612],[-63730,11026],[-52374,-55516],[-6310,90173],[22797,-39450],[-22314,-94950],[-15225,79748],[69933,-32347],[41440,-9644],[62951,-58618],[72200,23010]]
    result1 = solution.maxPartitionFactor(points1)
    print(f"Large case 1 (n=16) result: {result1}")
    
    points2 = [[-85401,-17368],[46585,-69810],[-36644,-97571],[-29903,-81002],[-41885,70137],[-96898,-36499],[-92734,-61360],[-954,36272],[39099,36602],[1986,58859],[97516,22190],[-56062,22948],[-50995,12805],[-13759,35414],[36630,18091],[83194,78635]]
    result2 = solution.maxPartitionFactor(points2)
    print(f"Large case 2 (n=16) result: {result2}")
    
    return result1, result2

def test_examples():
    solution = Solution()
    
    points1 = [[0,0],[0,2],[2,0],[2,2]]
    result1 = solution.maxPartitionFactor(points1)
    print(f"Example 1: {result1} (expected: 4)")
    
    points2 = [[0,0],[0,1],[10,0]]
    result2 = solution.maxPartitionFactor(points2)
    print(f"Example 2: {result2} (expected: 11)")

if __name__ == "__main__":
    print("Testing solution...")
    test_examples()
    print()
    test_large_case()
    print("\nAll tests completed!")
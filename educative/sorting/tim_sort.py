MIN_RUN = 32

def insertion_sort(arr, left, right):
    for i in range (left+1, right+1):
        temp = arr[i]
        j = i - 1
        while j >= left and arr[j] > temp: # trường hợp thằng trước lớn hơn thằng sau -> swap
            arr[j+1] = arr[j]
            j -= 1
        arr[j+1] = temp # trường hợp thằng sau đã lớn hơn thằng trước -> continue
    
def merge_sort(arr, left, mid, right): 
    len1, len2 = mid - left + 1, right - mid
    left_part, right_part = arr[left:mid+1], arr[mid+1:right+1] # galloping mode 
    
    i, j, k = 0, 0, left
    
    while i < len1 and j < len2:
        if left_part[i] <= right_part[j]:
            arr[k] = left_part[i]
            i += 1
        else:
            arr[k] = right_part[j]
            j += 1
        k +=1
            
    while i < len1:
        arr[k] = left_part[i]
        i += 1
        k += 1
        
    while j < len2:
        arr[k] = left_part[j]
        j += 1
        k += 1
    
def tim_sort(arr):
    n = len(arr)
    
    # 1. Sorting runs by insertion sort
    for i in range (0, n, MIN_RUN):
        insertion_sort(arr, i, min(i + MIN_RUN - 1, n - 1))
        
    # 2. Merge runs
    size = MIN_RUN
    while size < n:
        for left in range(0, n, 2*size): # galloping mode (implement bằng insertion binary search)
            mid = min(n-1, left+size - 1)
            right = min((left + 2 * size - 1), (n - 1))
            if mid < right:
                merge_sort(arr, left, mid, right)
        size *= 2
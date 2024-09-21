import unittest
from tim_sort import tim_sort

class TestTimSort(unittest.TestCase):
    def test_sorted_array(self):
        arr = [1,2,3,4,5,6]
        expected = [1, 2, 3, 4, 5,6]
        tim_sort(arr)
        self.assertEqual(arr, expected)
    
    def test_reverse_sorted_array(self):
        arr = [5, 4, 3, 2, 1]
        expected = [1, 2, 3, 4, 5]
        tim_sort(arr)
        self.assertEqual(arr, expected)

    def test_unsorted_array(self):
        arr = [10, 15, 20, 5, 30, 40, 50, 1, 60, 70]
        expected = [1, 5, 10, 15, 20, 30, 40, 50, 60, 70]
        tim_sort(arr)
        self.assertEqual(arr, expected)

    def test_empty_array(self):
        arr = []
        expected = []
        tim_sort(arr)
        self.assertEqual(arr, expected)

    def test_single_element_array(self):
        arr = [42]
        expected = [42]
        tim_sort(arr)
        self.assertEqual(arr, expected)

    def test_duplicate_elements(self):
        arr = [5, 3, 8, 3, 5, 9, 3]
        expected = [3, 3, 3, 5, 5, 8, 9]
        tim_sort(arr)
        self.assertEqual(arr, expected)
        
if __name__ == '__main__':
    unittest.main()

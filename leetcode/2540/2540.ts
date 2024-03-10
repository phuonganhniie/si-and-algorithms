// function getCommon(nums1: number[], nums2: number[]): number {
//   let i,
//     j = 0;

//   while (i < nums1.length && j < nums2.length) {
//     if (nums1[i] == nums2[j]) {
//       return nums1[i];
//     }

//     if (nums1[i] < nums2[j]) {
//       i++;
//     } else {
//       j++;
//     }
//   }
//   return -1;
// }

// Binary Search
function binarySearch(arr: number[], target: number): boolean {
  let left = 0,
    right = arr.length - 1;

  while (left <= right) {
    const mid = left + Math.floor((right - left) / 2);
    if (arr[mid] === target) {
      return true; // Target found
    }

    if (arr[mid] < target) {
      left = mid + 1; // Search right half
    } else {
      right = mid - 1; // Search left half
    }
  }
  return false; // Target not found
}

function getCommon(nums1: number[], nums2: number[]): number {
  if (nums1.length < nums2.length) {
    [nums1, nums2] = [nums2, nums1]; // Ensure nums2 is the smaller array
  }

  for (let num of nums1) {
    if (binarySearch(nums2, num)) {
      return num; // Common element found
    }
  }
  return -1; // No common element found
}

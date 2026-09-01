// #88 Merge Sorted Array [Easy]
// Tags: Array, Two Pointers, Sorting
//
// You are given two integer arrays nums1 and nums2, sorted in non-decreasing order, and two integers m and n, representing the number of elements in nums1 and nums2 respectively. Merge nums1 and nums2 into a single sorted array stored in nums1.
//
// Example 1: Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
//          Output: [1,2,2,3,5,6]
// Example 2: Input: nums1 = [1], m = 1, nums2 = [], n = 0
//          Output: [1]

function merge(nums1: number[], m: number, nums2: number[], n: number): void {

  let p1 = 0;
  let p2 = 0;
  const result: number[] = [];

  while (p1 < m && p2 < n) {
    if (nums1[p1] <= nums2[p2]) {
      result.push(nums1[p1]);
      p1++;
    } else {
      result.push(nums2[p2]);
      p2++;
    }
  }

  // 補回剩餘
  while (p1 < m) {
    result.push(nums1[p1]);
    p1++;
  }
  while (p2 < n) {
    result.push(nums2[p2]);
    p2++;
  }

  // 3. 蓋回原陣列 nums1
  for (let i = 0; i < result.length; i++) {
    nums1[i] = result[i];
  }
}

// --- Tests ---
console.log(JSON.stringify(merge([1,2,3,0,0,0], 3, [2,5,6], 3))); // → [1,2,2,3,5,6]
console.log(JSON.stringify(merge([1], 1, [], 0))); // → [1]

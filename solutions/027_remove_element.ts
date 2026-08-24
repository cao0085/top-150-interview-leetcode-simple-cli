// #27 Remove Element [Easy]
// Tags: Array, Two Pointers
//
// Given an integer array nums and an integer val, remove all occurrences of val in nums in-place. Return k, the number of elements in nums which are not equal to val.
//
// Example 1: Input: nums = [3,2,2,3], val = 3
//          Output: 2
// Example 2: Input: nums = [0,1,2,2,3,0,4,2], val = 2
//          Output: 5

function removeElement(nums: number[], val: number): number {    
 let du = 0;
 for(const num of nums) {
  if (num === val) du ++
 }

 return nums.length - du
}

// --- Tests ---
console.log(JSON.stringify(removeElement([3,2,2,3], 3))); // → 2
console.log(JSON.stringify(removeElement([0,1,2,2,3,0,4,2], 2))); // → 5

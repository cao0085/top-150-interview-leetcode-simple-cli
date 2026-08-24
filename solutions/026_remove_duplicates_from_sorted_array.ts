// #26 Remove Duplicates from Sorted Array [Easy]
// Tags: Array, Two Pointers
//
// Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that each unique element appears only once. Return k after placing the first k elements in nums.
//
// Example 1: Input: nums = [1,1,2]
//          Output: 2, nums = [1,2,_]
// Example 2: Input: nums = [0,0,1,1,1,2,2,3,3,4]
//          Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]

function removeDuplicates(nums: number[]): number {    
  let r = 1;
  let target = nums[0];

  for (const num of nums){
   if (target !== num) {
    target = num;
    r ++
   }
  }

  return r;
}

// --- Tests ---
console.log(JSON.stringify(removeDuplicates([1,1,2]))); // → 2
console.log(JSON.stringify(removeDuplicates([0,0,1,1,1,2,2,3,3,4]))); // → 5
console.log(JSON.stringify(removeDuplicates([1]))); // → 1

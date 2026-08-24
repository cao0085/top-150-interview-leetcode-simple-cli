// #1 Two Sum [Easy]
// Tags: Array, Hash Table
//
// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
//
// Example 1: Input: nums = [2,7,11,15], target = 9
//          Output: [0,1]
// Example 2: Input: nums = [3,2,4], target = 6
//          Output: [1,2]
// Example 3: Input: nums = [3,3], target = 6
//          Output: [0,1]

function twoSum(nums: number[], target: number): number[] {
 const n: number = nums.length; 
 for (let i = 0; i < n; i ++) {
    for (let j = i +:1; j < n; j++) {
      let a: number = nums[i];
      let b: number = nums[j];

      if(a+b === target) {
        return [i,j]
      }
    }
  }
}

// --- Tests ---
console.log(JSON.stringify(twoSum([2,7,11,15], 9))); // → [0,1]
console.log(JSON.stringify(twoSum([3,2,4], 6))); // → [1,2]
console.log(JSON.stringify(twoSum([3,3], 6))); // → [0,1]

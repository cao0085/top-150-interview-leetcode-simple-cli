// #15 3Sum [Medium]
// Tags: Array, Two Pointers, Sorting
//
// Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, j != k, and nums[i] + nums[j] + nums[k] == 0. The solution set must not contain duplicate triplets.
//
// Example 1: Input: nums = [-1,0,1,2,-1,-4]
//          Output: [[-1,-1,2],[-1,0,1]]
// Example 2: Input: nums = [0,1,1]
//          Output: []
// Example 3: Input: nums = [0,0,0]
//          Output: [[0,0,0]]

function threeSum(nums: number[]): number[][] {
 // 題目僅保證同組解答不含相同 INDEX
 // 相加等於零，可以理解為必定有零、負數值、a = -(b+c)
 nums.sort((a,b) => a - b);
 const result: number[][] = [];
 const n = nums.length;

 for(let i = 0; i < n; i ++) {
   let target: number = nums[i];
   if (target > 0) break;

   let p1: number = i+1;
   let p2: number = n-1;
    
   while(p1 < p2){
    sum = nums[p1] + nums[p2];

    total = target + sum;

    if (total > 0) {
      p1 ++
    } else if (total < 0) {
      p2 -= 1
    } else {
      result.push([target,nums[p1],nums[p2]])
      break;
    }
   }
 }
  
 return result;
}

// --- Tests ---
console.log(JSON.stringify(threeSum([-1,0,1,2,-1,-4]))); // → [[-1,-1,2],[-1,0,1]]
console.log(JSON.stringify(threeSum([0,0,0]))); // → [[0,0,0]]
console.log(JSON.stringify(threeSum([0,1,1]))); // → []

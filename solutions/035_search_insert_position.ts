// #35 Search Insert Position [Easy]
// Tags: Array, Binary Search
//
// Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order. Must run in O(log n).
//
// Example 1: Input: nums = [1,3,5,6], target = 5
//          Output: 2
// Example 2: Input: nums = [1,3,5,6], target = 2
//          Output: 1
// Example 3: Input: nums = [1,3,5,6], target = 7
//          Output: 4

function searchInsert(nums: number[], target: number): number {
    
}

// --- Tests ---
console.log(JSON.stringify(searchInsert([1,3,5,6], 5))); // → 2
console.log(JSON.stringify(searchInsert([1,3,5,6], 2))); // → 1
console.log(JSON.stringify(searchInsert([1,3,5,6], 7))); // → 4
console.log(JSON.stringify(searchInsert([1,3,5,6], 0))); // → 0

// #228 Summary Ranges [Easy]
// Tags: Array
//
// You are given a sorted unique integer array nums. A range [a,b] is the set of all integers from a to b. Return the smallest sorted list of ranges that cover all the numbers in the array exactly.
//
// Example 1: Input: nums = [0,1,2,4,5,7]
//          Output: ["0->2","4->5","7"]
// Example 2: Input: nums = [0,2,3,4,6,8,9]
//          Output: ["0","2->4","6","8->9"]

function summaryRanges(nums: number[]): string[] {
  const result: string = [];
  let current: [number, number] = [-1,-1];

  for(let i = 0; i < nums.length - 1; i ++) {
    const nn = nums[i]
   if(current[0] === -1) {
     current[0] = nn;
   } else {
    if ((nn + 1) !== nums[i+1]){
      current[1] = nn;
      result.push(`${current[0]}->${current[1]}`);
      current = [-1,-1];
    }
   }
  }

  return result;
}

// --- Tests ---
console.log(JSON.stringify(summaryRanges([0,1,2,4,5,7]))); // → ["0->2","4->5","7"]
console.log(JSON.stringify(summaryRanges([0,2,3,4,6,8,9]))); // → ["0","2->4","6","8->9"]
console.log(JSON.stringify(summaryRanges([]))); // → []

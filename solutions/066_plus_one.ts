// #66 Plus One [Easy]
// Tags: Array, Math
//
// You are given a large integer represented as an integer array digits, where each digits[i] is the ith digit of the integer. Increment the large integer by one and return the resulting array of digits.
//
// Example 1: Input: digits = [1,2,3]
//          Output: [1,2,4]
// Example 2: Input: digits = [4,3,2,1]
//          Output: [4,3,2,2]
// Example 3: Input: digits = [9]
//          Output: [1,0]

function plusOne(digits: number[]): number[] {
 const plusN = 1;
 let bePlusN = plusN;
 let p = digits.length - 1;
 const result: number = [...digits];

// 只要還有進位，且指標還在陣列範圍內
  while (bePlusN !== 0 && p >= 0) {
    const curr = result[p] + bePlusN;
    
    if (curr > 9) {
      result[p] = curr - 10;
      bePlusN = 1;
    } else {
      result[p] = curr;
      bePlusN = 0;
    }
    
    p--; // 正確往前移動指標
  }

  // 如果跑完迴圈 bePlusN 還是 1，代表全部進位（例如 [9, 9] 變成 [0, 0]）
  // 必須在最前面補 1
  if (bePlusN === 1) {
    result.unshift(1);
  }

 return result;
}

// --- Tests ---
console.log(JSON.stringify(plusOne([1,2,3]))); // → [1,2,4]
console.log(JSON.stringify(plusOne([9]))); // → [1,0]
console.log(JSON.stringify(plusOne([9,9]))); // → [1,0,0]

// #392 Is Subsequence [Easy]
// Tags: Two Pointers, String, Dynamic Programming
//
// Given two strings s and t, return true if s is a subsequence of t, or false otherwise. A subsequence of a string is a new string generated from the original string with some characters deleted without changing the relative order.
//
// Example 1: Input: s = "abc", t = "ahbgdc"
//          Output: true
// Example 2: Input: s = "axc", t = "ahbgdc"
//          Output: false

function isSubsequence(s: string, t: string): boolean {

 // a1, a2,
 
 i = 0;
 j = 0;

 while(i <= s.length - 1 && j <= t.length - 1) {
    const target = s[i];
    const cur = t[j];
    if (cur === target) {
      i ++
    }
    j ++
 }

 return (i === s.length)
}

// --- Tests ---
console.log(JSON.stringify(isSubsequence("abc", "ahbgdc"))); // → true
console.log(JSON.stringify(isSubsequence("axc", "ahbgdc"))); // → false
console.log(JSON.stringify(isSubsequence("", "ahbgdc"))); // → true

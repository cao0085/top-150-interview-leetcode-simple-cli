// #14 Longest Common Prefix [Easy]
// Tags: String, Trie
//
// Write a function to find the longest common prefix string amongst an array of strings. If there is no common prefix, return an empty string "".
//
// Example 1: Input: strs = ["flower","flow","flight"]
//          Output: "fl"
// Example 2: Input: strs = ["dog","racecar","car"]
//          Output: ""

function longestCommonPrefix(strs: string[]): string {
 if (strs.length === 0) {
   return ""
 } else if (strs.length === 1){
    return strs[0];
 }
 // 任意取一個為基準
 const first = strs[0];
 for (let i = 0; i < first.length; i++) {
  for (const str of strs) {

    if (str[i] !== first[i]) {
      return first.slice(0,i);
    }
  }
 }
}

// --- Tests ---
console.log(JSON.stringify(longestCommonPrefix(["flower","flow","flight"]))); // → "fl"
console.log(JSON.stringify(longestCommonPrefix(["dog","racecar","car"]))); // → ""
console.log(JSON.stringify(longestCommonPrefix(["interview"]))); // → "interview"

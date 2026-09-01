// #58 Length of Last Word [Easy]
// Tags: String
//
// Given a string s consisting of words and spaces, return the length of the last word in the string.
//
// Example 1: Input: s = "Hello World"
//          Output: 5
// Example 2: Input: s = "   fly me   to   the moon  "
//          Output: 4
// Example 3: Input: s = "luffy is still joyboy"
//          Output: 6

function lengthOfLastWord(s: string): number {

  // find last start & next empty
  let lastC = 0;
  let flag: boolean = false;

  for(let i = s.length - 1; i >= 0; i--) {
    let c = s[i]
    if (c === " " && flag === true) {
      return lastC
    } else if (c !== " ") {
      flag = true;
      lastC ++
    }
  }

  return lastC;
}

// --- Tests ---
console.log(JSON.stringify(lengthOfLastWord("Hello World"))); // → 5
console.log(JSON.stringify(lengthOfLastWord("   fly me   to   the moon  "))); // → 4
console.log(JSON.stringify(lengthOfLastWord("luffy is still joyboy"))); // → 6

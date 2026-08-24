// #20 Valid Parentheses [Easy]
// Tags: String, Stack
//
// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid. An input string is valid if open brackets are closed by the same type and in the correct order.
//
// Example 1: Input: s = "()"
//          Output: true
// Example 2: Input: s = "()[]{}"
//          Output: true
// Example 3: Input: s = "(]"
//          Output: false

function isValid(s: string): boolean {
  const stack: string[] = [];
  const dic: Record<string,string> = {
    "(":")",
    "{":"}",
    "[":"]"
  } as const;

  for (const ss of s) {

   const toS = dic[ss] ?? "";

   if (toS) {
    stack.push(toS)
   } else {
    const ps = stack.pop()
    if(ss !== ps) {
      return false
    }
   }
  }

  return true;
}

// --- Tests ---
console.log(JSON.stringify(isValid("()"))); // → true
console.log(JSON.stringify(isValid("()[]{}"))); // → true
console.log(JSON.stringify(isValid("(]"))); // → false
console.log(JSON.stringify(isValid("([])"))); // → true
console.log(JSON.stringify(isValid("{[]}"))); // → true

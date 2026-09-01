// #13 Roman to Integer [Easy]
// Tags: Hash Table, Math, String
//
// Given a roman numeral, convert it to an integer. Roman numerals are represented by seven different symbols: I(1), V(5), X(10), L(50), C(100), D(500), M(1000). Subtraction rules apply for IV(4), IX(9), XL(40), XC(90), CD(400), CM(900).
//
// Example 1: Input: s = "III"
//          Output: 3
// Example 2: Input: s = "LVIII"
//          Output: 58
// Example 3: Input: s = "MCMXCIV"
//          Output: 1994

function romanToInt(s: string): number {
  const RM = {
    I: 1,
    IV: 4,
    V: 5,
    IX: 9,
    X: 10,
    XL: 40,
    L: 50,
    XC: 90,
    C: 100,
    CD: 400,
    D: 500,
    CM: 900,
    M: 1000,
  } as const;

  type RomanKey = keyof typeof RM;

  let result = 0;
  let i = 0;

  while (i < s.length) {
    const twoChars = s.slice(i, i + 2) as RomanKey;

    if (i + 1 < s.length && twoChars in RM) {
      result += RM[twoChars];
      i += 2; // 跳過兩個字元
    } else {
      const singleChar = s[i] as RomanKey;
      result += RM[singleChar];
      i += 1; // 進到下一個字元
    }
  }

  return result
}

// --- Tests ---
console.log(JSON.stringify(romanToInt("III"))); // → 3
console.log(JSON.stringify(romanToInt("LVIII"))); // → 58
console.log(JSON.stringify(romanToInt("MCMXCIV"))); // → 1994

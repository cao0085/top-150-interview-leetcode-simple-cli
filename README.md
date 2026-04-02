# LeetCode Top 150 Practice CLI

本地練習工具，不需要登入帳號，支援 TypeScript / JavaScript / Go / C# 解答。

## 安裝需求

- [Node.js](https://nodejs.org/) （必要，用來執行 CLI）
- Go / .NET SDK（可選，依你要寫的語言決定）

## 快速開始

```powershell
cd C:\tony\leetcode
node leet.js
```

## 指令

| 指令 | 說明 |
|------|------|
| `node leet.js list` | 列出全部 150 題 |
| `node leet.js list easy` | 篩選 Easy 題目 |
| `node leet.js list medium` | 篩選 Medium 題目 |
| `node leet.js list hard` | 篩選 Hard 題目 |
| `node leet.js show <id>` | 顯示題目描述、範例、Constraints |
| `node leet.js pick <id> [lang]` | 建立解答檔並顯示題目 |
| `node leet.js run <id>` | 執行解答並跑測資 |

## 選題並開始練習

```powershell
# 1. 選一題，建立解答檔
node leet.js pick 1 ts

# 2. 用 vim 打開解答
vim solutions/001_two_sum.ts

# 3. 寫完後執行測試
node leet.js run 1
```

## 支援語言

| 語言 | 副檔名 | 執行方式 |
|------|--------|---------|
| TypeScript | `.ts` | 自動用 `npx tsx`（首次會安裝） |
| JavaScript | `.js` | `node` |
| Go | `.go` | `go run <file>`（需安裝 Go） |
| C# | `.cs` | `dotnet script <file>`（需安裝 dotnet-script） |

```powershell
node leet.js pick 1 ts    # TypeScript
node leet.js pick 1 go    # Go
node leet.js pick 1 cs    # C#
```

## 解答檔格式

每個解答檔包含：
1. 題目描述（註解）
2. Function signature（填入解法）
3. 測資（`console.log` / `fmt.Println` / `Console.WriteLine`）

**TypeScript 範例：**
```typescript
// #70 Climbing Stairs [Easy]
// Tags: Math, Dynamic Programming

function climbStairs(n: number): number {
    // 在這裡寫解法
}

// --- Tests ---
console.log(JSON.stringify(climbStairs(1))); // → 1
console.log(JSON.stringify(climbStairs(2))); // → 2
console.log(JSON.stringify(climbStairs(10))); // → 89
```

執行 `node leet.js run 70` 後，輸出結果與 `// →` 後面的值比對即可。

## 檔案結構

```
leetcode/
├── leet.js          # CLI 主程式
├── problems.json    # 150 題題目資料（描述、範例、測資、模板）
├── solutions/       # 你的解答（自動建立）
│   ├── 001_two_sum.ts
│   ├── 070_climbing_stairs.ts
│   └── ...
└── README.md
```

## 進度追蹤

`list` 指令會顯示已建立解答檔的題目（✓ 標記）：

```
 ✓  1     Two Sum                    Easy      Array, Hash Table
    3     Longest Substring...       Medium    Hash Table, String
```

## C# 執行環境設定（可選）

```powershell
# 安裝 dotnet-script
dotnet tool install -g dotnet-script
```

## Go 執行環境設定（可選）

從 https://go.dev/dl/ 下載安裝即可。

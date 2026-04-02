#!/usr/bin/env node
'use strict';

const fs = require('fs');
const path = require('path');
const { spawnSync } = require('child_process');

const ROOT = __dirname;
const SOLUTIONS_DIR = path.join(ROOT, 'solutions');
const PROBLEMS = require('./problems.json');

const $ = {
  r: '\x1b[0m', b: '\x1b[1m',
  green: '\x1b[32m', red: '\x1b[31m', yellow: '\x1b[33m',
  blue: '\x1b[34m', cyan: '\x1b[36m', gray: '\x1b[90m',
};

const diffColor = d => d === 'Easy' ? $.green : d === 'Medium' ? $.yellow : $.red;

function solveFile(p, lang) {
  const n = String(p.id).padStart(3, '0');
  const slug = p.title.toLowerCase().replace(/[^a-z0-9]+/g, '_').replace(/_+$/, '');
  return path.join(SOLUTIONS_DIR, `${n}_${slug}.${lang}`);
}

function cmd_list(filter) {
  let ps = PROBLEMS;
  if (filter && ['easy','medium','hard'].includes(filter.toLowerCase()))
    ps = ps.filter(p => p.difficulty.toLowerCase() === filter.toLowerCase());

  console.log(`\n${$.b}LeetCode Top 150 Interview Questions${$.r}\n`);
  console.log($.gray + '     ID   ' + 'Title'.padEnd(46) + 'Diff'.padEnd(10) + 'Tags' + $.r);
  console.log($.gray + '─'.repeat(95) + $.r);

  let solved = 0;
  for (const p of ps) {
    const has = ['ts','js','go','cs'].some(l => fs.existsSync(solveFile(p, l)));
    if (has) solved++;
    const mark = has ? $.green + '✓' + $.r : ' ';
    const id   = String(p.id).padEnd(6);
    const ttl  = p.title.padEnd(46);
    const df   = diffColor(p.difficulty) + p.difficulty.padEnd(10) + $.r;
    const tg   = $.gray + (p.tags||[]).slice(0,3).join(', ') + $.r;
    console.log(` ${mark}  ${id}${ttl}${df}${tg}`);
  }
  console.log(`\n${$.gray}${solved}/${ps.length} solved${$.r}\n`);
}

function cmd_show(id) {
  const p = PROBLEMS.find(x => x.id === id);
  if (!p) return console.log($.red + `Problem #${id} not found` + $.r);

  console.log(`\n${$.b}${$.blue}#${p.id} ${p.title}${$.r}  ${diffColor(p.difficulty)}[${p.difficulty}]${$.r}  ${$.gray}${(p.tags||[]).join(', ')}${$.r}`);
  console.log();
  console.log(p.description || $.gray + '(No description — look it up on leetcode.com)' + $.r);

  if (p.examples?.length) {
    console.log(`\n${$.b}Examples:${$.r}`);
    p.examples.forEach((ex, i) => {
      console.log(`\n  Example ${i + 1}:`);
      console.log(`    Input:  ${ex.input}`);
      console.log(`    Output: ${ex.output}`);
      if (ex.explanation) console.log(`    ${$.gray}${ex.explanation}${$.r}`);
    });
  }
  if (p.constraints?.length) {
    console.log(`\n${$.b}Constraints:${$.r}`);
    p.constraints.forEach(c => console.log(`  • ${c}`));
  }
  console.log();
}

function goLit(v) {
  if (v === null) return 'nil';
  if (typeof v === 'boolean') return String(v);
  if (typeof v === 'number') return String(v);
  if (typeof v === 'string') return `"${v}"`;
  if (Array.isArray(v)) {
    if (v.length === 0) return '[]int{}';
    if (typeof v[0] === 'number') return `[]int{${v.join(',')}}`;
    if (typeof v[0] === 'string') return `[]string{${v.map(s => `"${s}"`).join(',')}}`;
  }
  return JSON.stringify(v);
}

function makeTemplate(p, lang) {
  const fnName = p.function_name || 'solution';
  const descLines = (p.description || '').split('\n').map(l => '// ' + l).join('\n');
  const exLines = (p.examples||[]).map((e,i) =>
    `// Example ${i+1}: Input: ${e.input}\n//          Output: ${e.output}`
  ).join('\n');

  if (lang === 'ts' || lang === 'js') {
    const sig = lang === 'ts'
      ? (p.ts_template || `function ${fnName}(...args: any[]): any {\n    \n}`)
      : (p.js_template || p.ts_template?.replace(/:\s*\w[\w\[\]<>,\s]*/g, '') || `function ${fnName}(...args) {\n    \n}`);

    const tests = (p.test_cases||[]).map(tc => {
      const args = Array.isArray(tc.input)
        ? tc.input.map(a => JSON.stringify(a)).join(', ')
        : JSON.stringify(tc.input);
      return `console.log(JSON.stringify(${fnName}(${args}))); // → ${JSON.stringify(tc.expected)}`;
    }).join('\n');

    return [
      `// #${p.id} ${p.title} [${p.difficulty}]`,
      `// Tags: ${(p.tags||[]).join(', ')}`,
      `//`,
      descLines,
      `//`,
      exLines,
      ``,
      sig,
      ``,
      tests ? `// --- Tests ---\n${tests}` : `// Add test calls here`,
      ``
    ].join('\n');
  }

  if (lang === 'go') {
    const sig = p.go_template || `func ${fnName}() {\n    // TODO\n}`;
    const tests = (p.test_cases||[]).map(tc => {
      const args = Array.isArray(tc.input)
        ? tc.input.map(a => goLit(a)).join(', ')
        : goLit(tc.input);
      return `\tfmt.Println(${fnName}(${args})) // → ${JSON.stringify(tc.expected)}`;
    }).join('\n') || '\t// Add test calls here';

    return `package main\n\nimport "fmt"\n\n// #${p.id} ${p.title} [${p.difficulty}]\n// Tags: ${(p.tags||[]).join(', ')}\n\n${sig}\n\nfunc main() {\n${tests}\n}\n`;
  }

  if (lang === 'cs') {
    const sig = p.cs_template || `    public object Solution() {\n        throw new NotImplementedException();\n    }`;
    const tests = (p.test_cases||[]).map(tc =>
      `        // Input: ${JSON.stringify(tc.input)} → Expected: ${JSON.stringify(tc.expected)}`
    ).join('\n') || '        // Add test calls here';

    return `// #${p.id} ${p.title} [${p.difficulty}]\n// Tags: ${(p.tags||[]).join(', ')}\nusing System;\nusing System.Linq;\nusing System.Collections.Generic;\n\npublic class Solution {\n${sig}\n\n    static void Main() {\n        var s = new Solution();\n${tests}\n    }\n}\n`;
  }

  return `// #${p.id} ${p.title}\n// TODO\n`;
}

function cmd_pick(id, lang = 'ts') {
  const p = PROBLEMS.find(x => x.id === id);
  if (!p) return console.log($.red + `Problem #${id} not found` + $.r);

  cmd_show(id);

  if (!fs.existsSync(SOLUTIONS_DIR)) fs.mkdirSync(SOLUTIONS_DIR, { recursive: true });

  const file = solveFile(p, lang);
  if (!fs.existsSync(file)) {
    fs.writeFileSync(file, makeTemplate(p, lang));
    console.log($.green + `✓ Created: solutions/${path.basename(file)}` + $.r);
  } else {
    console.log($.gray + `Already exists: solutions/${path.basename(file)}` + $.r);
  }

  const rel = `solutions/${path.basename(file)}`;
  const runCmd = lang === 'ts' || lang === 'js'
    ? `node leet.js run ${id}`
    : lang === 'go' ? `go run ${rel}` : `dotnet script ${rel}`;

  console.log(`\n  ${$.b}Edit:${$.r}  vim ${rel}`);
  console.log(`  ${$.b}Run:${$.r}   ${runCmd}\n`);
}

function cmd_run(id) {
  const p = PROBLEMS.find(x => x.id === id);
  if (!p) return console.log($.red + `Problem #${id} not found` + $.r);

  let file = null, lang = null;
  for (const l of ['ts', 'js', 'go', 'cs']) {
    const f = solveFile(p, l);
    if (fs.existsSync(f)) { file = f; lang = l; break; }
  }

  if (!file) {
    console.log($.red + `No solution found. Try: node leet.js pick ${id} ts` + $.r);
    return;
  }

  console.log(`\n${$.b}▶ Running #${p.id} ${p.title}${$.r}\n`);

  if (lang === 'ts') {
    let res = spawnSync('npx', ['--yes', 'tsx', file], { shell: true, stdio: 'inherit' });
    if (res.status !== 0 && res.error) {
      console.log($.red + 'Failed to run tsx. Make sure Node.js is installed.' + $.r);
    }
  } else if (lang === 'js') {
    spawnSync('node', [file], { shell: true, stdio: 'inherit' });
  } else if (lang === 'go') {
    const res = spawnSync('go', ['run', file], { shell: true, stdio: 'inherit' });
    if (res.status !== 0) console.log($.yellow + 'Make sure Go is installed: https://go.dev' + $.r);
  } else if (lang === 'cs') {
    const res = spawnSync('dotnet', ['script', file], { shell: true, stdio: 'inherit' });
    if (res.status !== 0) console.log($.yellow + 'Install dotnet-script: dotnet tool install -g dotnet-script' + $.r);
  }
  console.log();
}

// --- Entry ---
const [,, cmd, a1, a2] = process.argv;
if      (cmd === 'list') cmd_list(a1);
else if (cmd === 'show') cmd_show(parseInt(a1));
else if (cmd === 'pick') cmd_pick(parseInt(a1), a2);
else if (cmd === 'run')  cmd_run(parseInt(a1));
else console.log(`
${$.b}leet${$.r} — LeetCode Top 150 Practice CLI

  ${$.cyan}node leet.js list [easy|medium|hard]${$.r}
  ${$.cyan}node leet.js show <id>${$.r}
  ${$.cyan}node leet.js pick <id> [ts|js|go|cs]${$.r}
  ${$.cyan}node leet.js run  <id>${$.r}

Examples:
  node leet.js list easy
  node leet.js show 1
  node leet.js pick 1 ts     → creates solutions/001_two_sum.ts
  node leet.js run 1         → runs with npx tsx (auto-installs)
`);

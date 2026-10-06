#!/usr/bin/env node
/*
ZERO machine-language reference runtime.
The host is JavaScript only; the language being executed is ZERO.LANGUAGE.
No predefined task answer, semantic dictionary, model, Python or Rust.
*/

const fs = require("fs");

function tokenize(input) {
  return input.split(/\s+/).filter(Boolean);
}

function trace(state, token) {
  if (!state.seen.has(token)) state.seen.add(token);
  state.history.push({op:"trace", token});
}

function step(state, token) {
  const out = [];
  if (token.startsWith("◇") && token.endsWith("◇")) {
    const x = token.slice(1,-1);
    trace(state,x);
    out.push("appear:"+x);
  } else if (token.includes("!")) {
    const [a,b] = token.split("!",2);
    trace(state,a); trace(state,b);
    const collision = a === b ? "same" : "different";
    out.push("collision:"+collision+":"+a+":"+b);
  } else if (token.includes("~")) {
    const [a,b] = token.split("~",2);
    trace(state,a); trace(state,b);
    out.push("variation:"+a+":"+b);
  } else if (token.includes("|")) {
    const [a,b] = token.split("|",2);
    trace(state,a); trace(state,b);
    out.push("neighbor:"+a+":"+b);
  } else if (token.startsWith("<") && token.endsWith(">")) {
    const x = token.slice(1,-1);
    if (state.seen.has(x)) out.push("reinject:"+x);
    else out.push("unknown-reinject:"+x);
  } else if (token === "...") {
    out.push("unknown");
  } else {
    out.push("opaque:"+token);
  }
  state.history.push({op:"step", token, out});
  return out;
}

function run(input) {
  const state = {seen:new Set(), history:[]};
  const output=[];
  for (const token of tokenize(input)) output.push(...step(state,token));
  return {output, distinct:state.seen.size, history:state.history.length};
}

if (require.main === module) {
  const input = process.argv.slice(2).join(" ");
  if (!input) {
    console.error("usage: node zero-runtime.js '<ZERO program>'");
    process.exit(2);
  }
  process.stdout.write(JSON.stringify(run(input), null, 2)+"\n");
}

module.exports = {run, tokenize};

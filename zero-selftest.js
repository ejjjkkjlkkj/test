const assert = require("assert");
const {run} = require("./zero-runtime");

function expect(program, fragment) {
  const r = run(program);
  assert(r.output.some(x => x.includes(fragment)),
    program+" did not produce "+fragment+"; got "+JSON.stringify(r.output));
}

expect("◇a◇ a|b a!b <a>", "appear:a");
expect("◇a◇ a|b a!b <a>", "neighbor:a:b");
expect("◇a◇ a|b a!b <a>", "collision:different:a:b");
expect("◇a◇ a|b a!b <a>", "reinject:a");
expect("◇?◇ ?~u ...", "variation:?:u");

const transferA = run("◇a7◇ k2|a7 a7!m4 <k2>");
const transferB = run("◇q3◇ n8|q3 q3!t1 <n8>");
assert(transferA.output.some(x => x.startsWith("collision:")));
assert(transferB.output.some(x => x.startsWith("collision:")));

console.log("ZERO SELFTEST PASS");

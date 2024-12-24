import { part1 } from "./part1.ts";
import { part2 } from "./part2.ts";

const startTime = performance.now();
console.log("part1:", part1("input.txt"));
const endTime = performance.now();
console.log(`Execution time: ${endTime - startTime} milliseconds`);

const startTime2 = performance.now();
console.log("part2:", part2("input.txt"));
const endTime2 = performance.now();
console.log(`Execution time: ${endTime2 - startTime2} milliseconds`);

import { assertEquals } from "@std/assert";
import { part1 } from "./part1.ts";
import { getTrailHeads } from "./helpers/getTrailHeads.ts";
import { getTopoMapMatrix } from "./helpers/getTopoMapMatrix.ts";
import { getTrailHeadScore } from "./helpers/getTrailHeadScore.ts";
import { part2 } from "./part2.ts";

const topoMapMatrix = getTopoMapMatrix("test2.txt");

Deno.test(function getTrailHeadsTest() {
  assertEquals(getTrailHeads({ topoMapMatrix: topoMapMatrix }), [
    { x: 1, y: 0 },
    { x: 5, y: 6 },
  ]);
});

Deno.test(function getTrailHeadScoreTest() {
  assertEquals(
    getTrailHeadScore({
      topoMapMatrix,
      currentPosition: { x: 1, y: 0, elevation: 0 },
    }),
    1,
  );
});

Deno.test(function part1Test3() {
  assertEquals(part1("test3.txt"), 2);
});

Deno.test(function part1Test4() {
  assertEquals(part1("test4.txt"), 4);
});

Deno.test(function part1Test5() {
  assertEquals(part1("test5.txt"), 3);
});

Deno.test(function part1Test() {
  assertEquals(part1("test.txt"), 36);
});

Deno.test(function part1Answer() {
  assertEquals(part1("input.txt"), 482);
});

// =====================================================
// PART 2
// =====================================================

Deno.test(function part2Test1() {
  assertEquals(part2("test6.txt"), 3);
});

Deno.test(function part2Test2() {
  assertEquals(part2("test7.txt"), 13);
});

Deno.test(function part2Test3() {
  assertEquals(part2("test.txt"), 81);
});

Deno.test(function part2Answer() {
  assertEquals(part2("input.txt"), 1094);
});

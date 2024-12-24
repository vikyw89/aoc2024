import { assertEquals } from "@std/assert";
import { part1 } from "./part1.ts";
import { getTrailHeads } from "./helpers/getTrailHeads.ts";
import { getTopoMapMatrix } from "./helpers/getTopoMapMatrix.ts";
import { getTrailHeadScore } from "./helpers/getTrailHeadScore.ts";

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

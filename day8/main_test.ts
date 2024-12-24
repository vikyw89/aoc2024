import { assertEquals } from "@std/assert";
import { part1 } from "./part1.ts";
import { part2 } from "./part2.ts";

Deno.test(function testPart1() {
  assertEquals(part1("test.txt"), 14);
});

Deno.test(function answerPart1() {
  assertEquals(part1("input.txt"), 313);
});

Deno.test(function testPart2V2() {
  assertEquals(part2("test2.txt"), 9);
});

Deno.test(function testPart2() {
  assertEquals(part2("test.txt"), 34);
});

Deno.test(function answerPart2() {
  assertEquals(part2("input.txt"), 1064);
});

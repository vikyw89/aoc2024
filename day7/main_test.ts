import { assertEquals } from "@std/assert";
import { part1 } from "./part1.ts";
import { part2 } from "./part2.ts";

Deno.test(function part1Test() {
  assertEquals(part1("test.txt"), 3749);
});

Deno.test(function part1Answer() {
  assertEquals(part1("input.txt"), 663613490587);
});

Deno.test(function part2Test() {
  assertEquals(part2("test.txt"), 11387);
});

Deno.test(function part2Answer() {
  assertEquals(part2("input.txt"), 110365987435001);
});

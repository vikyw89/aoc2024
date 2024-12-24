import { assertEquals, assertGreater } from "@std/assert";
import { part1 } from "./part1.ts";
import { part2 } from "./part2.ts";
import { getCheckSum } from "./helpers/getChecksum.ts";
import { rollFileBlockToFront } from "./helpers/rollFileBlockToFront.ts";

Deno.test(function testPart1() {
  assertEquals(part1("test.txt"), 1928);
});

Deno.test(function answerPart1() {
  assertGreater(part1("input.txt"), 100);
});

Deno.test(function getCheckSumTest() {
  assertEquals(
    getCheckSum(
      "00992111777.44.333....5555.6666.....8888..".split("").map((v) => {
        if (v === ".") return -1;
        return Number(v);
      }),
    ),
    2858,
  );
});

Deno.test(function rollFileBlockToFrontTest() {
  assertEquals(
    rollFileBlockToFront({
      decompressedDiskmap: "00...111...2...333.44.5555.6666.777.888899"
        .split("")
        .map((v) => {
          if (v === ".") return -1;
          return Number(v);
        }),
    }),
    "00992111777.44.333....5555.6666.....8888..".split("").map((v) => {
      if (v === ".") return -1;
      return Number(v);
    }),
  );
});

Deno.test(function testPart2() {
  assertEquals(part2("test.txt"), 2858);
});

Deno.test(function answerPart2() {
  assertEquals(part2("input.txt"), 6327174563252);
});

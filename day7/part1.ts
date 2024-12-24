import { getOperatorsToSolve } from "./helpers/solver.ts";

export const part1 = (filePath: string): number => {
  const lines = Deno.readTextFileSync(filePath).split("\n");
  const totalAndNumbers = lines.map((line) => {
    const [total, numbers] = line.split(": ");
    return {
      total: Number(total),
      numbers: numbers.split(" ").map(Number),
    };
  });

  let results = 0;
  for (const { total, numbers } of totalAndNumbers) {
    const result = getOperatorsToSolve({
      numbers: numbers,
      target: total,
      availableOperators: ["+", "*"],
    });

    if (result) {
      results += total;
    }
  }

  return results;
};

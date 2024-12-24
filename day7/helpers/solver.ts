export type TotalAndNumbers = {
  total: number;
  numbers: number[];
};

export const evaluate = ({
  ops,
  numbers,
  target,
}: {
  ops: string[];
  numbers: number[];
  target: number;
}): boolean => {
  let result = numbers[0];
  for (let i = 0; i < ops.length; i++) {
    if (ops[i] === "+") {
      result += numbers[i + 1];
    } else if (ops[i] === "*") {
      result *= numbers[i + 1];
    } else if (ops[i] === "||") {
      result = Number(`${result}${numbers[i + 1]}`);
    }
  }
  return result === target;
};

function* generateOps({
  n,
  availableOperators,
}: {
  n: number;
  availableOperators: string[];
}): Generator<string[]> {
  // base case
  if (n === 0) {
    yield [];
    return;
  }

  //   recursive case
  for (const prevOps of generateOps({ n: n - 1, availableOperators })) {
    for (const op of availableOperators) {
      yield [...prevOps, op];
    }
  }
}

export const getOperatorsToSolve = ({
  target,
  numbers,
  availableOperators,
}: {
  target: number;
  numbers: number[];
  availableOperators: string[];
}): string[] | undefined => {
  const nOps = numbers.length - 1;
  for (const ops of generateOps({ n: nOps, availableOperators })) {
    if (evaluate({ ops, numbers, target })) {
      return ops;
    }
  }

  return;
};

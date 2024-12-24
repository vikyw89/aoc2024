export const rollLastCharToFront = (input: number[]): number[] => {
  // 2 pointers
  let leftIndex = 0;
  let rightIndex = input.length - 1;

  const stringArr = [...input];

  while (leftIndex < rightIndex) {
    const rightChar = stringArr[rightIndex];
    const leftChar = stringArr[leftIndex];

    // if left char is not empty, increment and continue
    if (leftChar >= 0) {
      leftIndex++;
      continue;
    }

    // if right chat is empty, decrement and continue
    if (rightChar <= 0) {
      rightIndex--;
      continue;
    }

    // swap
    stringArr[leftIndex] = rightChar;
    stringArr[rightIndex] = leftChar;

    leftIndex++;
    rightIndex--;
  }

  return stringArr;
};

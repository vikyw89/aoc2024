export const getCheckSum = (input: number[]): number => {
  let checkSum = 0;
  for (let i = 0; i < input.length; i++) {
    const fileId = input[i];

    // if fileId is -1 (space)
    if (fileId < 0) {
      continue;
    }

    const value = fileId * i;

    checkSum += value;
  }

  return checkSum;
};

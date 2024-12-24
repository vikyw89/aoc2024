export const decompressDiskmap = (input: string): number[] => {
  const decompressedArr: number[] = [];

  let currentFileId = 0;
  for (let i = 0; i < input.length; i = i + 2) {
    const fileBlock = Number(input[i]);
    const space = input[i + 1] ? Number(input[i + 1]) : 0;

    // process character
    for (let j = 0; j < fileBlock; j++) {
      decompressedArr.push(currentFileId);
    }

    // process space
    for (let j = 0; j < space; j++) {
      decompressedArr.push(-1);
    }

    // increment fileId
    currentFileId++;
  }

  return decompressedArr;
};

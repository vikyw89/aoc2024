export const getFileSizeMap = (input: string): Map<number, number> => {
  const fileSizeMap = new Map<number, number>();

  let currentFileId = 0;
  for (let i = 0; i < input.length; i = i + 2) {
    const fileBlock = Number(input[i]);

    // process character
    fileSizeMap.set(currentFileId, fileBlock);

    currentFileId++;
  }

  return fileSizeMap;
};

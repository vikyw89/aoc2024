export type Space = {
  startIndex: number;
  size: number;
};

export const getSpaceSizeArr = (input: string): Space[] => {
  const spaceSizeArr: Space[] = [];
  let index = 0;
  for (let i = 0; i < input.length; i = i + 2) {
    const fileBlockSize = Number(input[i]);
    index += fileBlockSize;

    const space = input[i + 1] ? Number(input[i + 1]) : 0;

    // process space
    spaceSizeArr.push({
      startIndex: i,
      size: space,
    });
  }
  return spaceSizeArr;
};

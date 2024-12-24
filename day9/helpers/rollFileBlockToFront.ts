export const rollFileBlockToFront = ({
  decompressedDiskmap,
}: {
  decompressedDiskmap: number[];
}): number[] => {
  // get last file block
  const diskmap = [...decompressedDiskmap];

  let currentIndex = diskmap.length - 1;

  const fileIdMap = new Set<number>();

  // we go from the end to the start
  while (currentIndex > 0) {
    const char = diskmap[currentIndex];

    // if char is a space, decrement
    if (char === -1) {
      currentIndex--;
      continue;
    }

    // retrieve file block size, and fileblock id
    const fileBlock = getFileBlock({
      diskMapArr: diskmap,
      endIndex: currentIndex,
    });

    const fileBlockSize = fileBlock.endIndex - fileBlock.startIndex + 1;

    if (fileIdMap.has(char)) {
      currentIndex -= fileBlockSize;
      continue;
    }

    fileIdMap.add(char);

    // find the space
    const spaceIndex = getSpaceIndex({
      diskMapArr: diskmap,
      spaceSize: fileBlockSize,
      maxIndex: currentIndex,
    });

    // if spaceIndex is undefined, break
    if (spaceIndex === undefined) {
      currentIndex -= fileBlockSize;
      continue;
    }

    // swap file block with space
    // wipe file block
    for (let i = fileBlock.startIndex; i <= fileBlock.endIndex; i++) {
      diskmap[i] = -1;
    }

    // put file block at spaceIndex
    for (let i = spaceIndex; i < fileBlockSize + spaceIndex; i++) {
      diskmap[i] = fileBlock.fileId;
    }

    // decrement currentIndex
    currentIndex -= fileBlockSize;
  }

  return diskmap;
};

export const getSpaceIndex = ({
  diskMapArr,
  spaceSize,
  maxIndex,
}: {
  diskMapArr: number[];
  spaceSize: number;
  maxIndex: number;
}): number | undefined => {
  let currentIndex = 0;

  while (currentIndex < maxIndex) {
    const char = diskMapArr[currentIndex];

    if (char !== -1) {
      currentIndex++;
      continue;
    }

    const currentSpaceSize = getSpaceSize({
      diskMapArr,
      startIndex: currentIndex,
    });

    if (currentSpaceSize >= spaceSize) {
      return currentIndex;
    }

    // increment
    currentIndex += currentSpaceSize;
  }

  //   no index found
  return;
};

export type Space = {
  startIndex: number;
  lastIndex: number;
};

export const getSpaceSize = ({
  diskMapArr,
  startIndex,
}: {
  diskMapArr: number[];
  startIndex: number;
}): number => {
  let currentIndex = startIndex + 1;

  while (currentIndex < diskMapArr.length) {
    const char = diskMapArr[currentIndex];

    // if char is a space, increment
    if (char !== -1) {
      break;
    }

    currentIndex++;
  }

  return currentIndex - startIndex;
};

export type FileBlock = {
  startIndex: number;
  endIndex: number;
  fileId: number;
};

export const getFileBlock = ({
  diskMapArr,
  endIndex,
}: {
  diskMapArr: number[];
  endIndex: number;
}): FileBlock => {
  const fileId = diskMapArr[endIndex];

  let currentIndex = endIndex - 1;

  while (currentIndex > 0) {
    const char = diskMapArr[currentIndex];

    // if char doesn't have the same fileId, break
    if (char !== fileId) {
      break;
    }

    currentIndex--;
  }

  return {
    startIndex: currentIndex + 1,
    endIndex,
    fileId,
  };
};

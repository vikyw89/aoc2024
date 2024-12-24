import { decompressDiskmap } from "./helpers/decompressDiskmap.ts";
import { getCheckSum } from "./helpers/getChecksum.ts";

import { rollFileBlockToFront } from "./helpers/rollFileBlockToFront.ts";

export const part2 = (filePath: string): number => {
  const fileString = Deno.readTextFileSync(filePath);
  const decrompressedDiskmap = decompressDiskmap(fileString);
  const rolledDiskMap = rollFileBlockToFront({
    decompressedDiskmap: decrompressedDiskmap,
  });
  const checkSum = getCheckSum(rolledDiskMap);
  return checkSum;
};

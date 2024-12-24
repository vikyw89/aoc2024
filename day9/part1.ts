import { decompressDiskmap } from "./helpers/decompressDiskmap.ts";
import { rollLastCharToFront } from "./helpers/rollLastCharToFront.ts";
import { getCheckSum } from "./helpers/getChecksum.ts";

export const part1 = (filePath: string): number => {
  const fileString = Deno.readTextFileSync(filePath);
  const decompressedDiskmap = decompressDiskmap(fileString);
  const rolledDiskmap = rollLastCharToFront(decompressedDiskmap);
  const checkSum = getCheckSum(rolledDiskmap);

  return checkSum;
};

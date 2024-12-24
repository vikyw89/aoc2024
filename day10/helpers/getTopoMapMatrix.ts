export const getTopoMapMatrix = (filePath: string): number[][] => {
  const fileString = Deno.readTextFileSync(filePath);
  return fileString
    .split("\n")
    .map((line) => line.split("").map((v) => Number(v)));
};

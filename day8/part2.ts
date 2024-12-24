import {
  calculateAntinodeCoordinatesV2,
  getAntennas,
} from "./helpers/index.ts";

export const part2 = (filePath: string): number => {
  const fileString = Deno.readTextFileSync(filePath);
  const antennas = getAntennas(fileString);
  const lines = fileString.trim().split("\n");
  const grid = lines.map((line) => line.split(""));

  const visitedAntinodes = new Set<string>();

  for (let i = 0; i < antennas.length; i++) {
    for (let j = i + 1; j < antennas.length; j++) {
      const antenna1 = antennas[i];
      const antenna2 = antennas[j];

      if (antenna1.frequency === antenna2.frequency) {
        // add antena as antinode
        const antenna1Key = `${antenna1.coordinate.x},${antenna1.coordinate.y}`;
        const antenna2Key = `${antenna2.coordinate.x},${antenna2.coordinate.y}`;
        if (!visitedAntinodes.has(antenna1Key)) {
          visitedAntinodes.add(antenna1Key);
        }
        if (!visitedAntinodes.has(antenna2Key)) {
          visitedAntinodes.add(antenna2Key);
        }
        const antinodesArr = calculateAntinodeCoordinatesV2(
          antenna1,
          antenna2,
          grid,
        );
        for (const antinode of antinodesArr) {
          const key = `${antinode.x},${antinode.y}`;
          if (!visitedAntinodes.has(key)) {
            visitedAntinodes.add(key);
          }
        }
      }
    }
  }

  return visitedAntinodes.size;
};

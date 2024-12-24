import {
  calculateAntinodeCoordinates,
  Coordinate,
  getAntennas,
  isValidCoordinate,
} from "./helpers/index.ts";

export const part1 = (filePath: string): number => {
  const input = Deno.readTextFileSync(filePath);
  const antennas = getAntennas(input);
  const lines = input.trim().split("\n");
  const grid = lines.map((line) => line.split(""));

  const antinodeCoordinates: Coordinate[] = [];
  const visitedAntinodes = new Set<string>();

  for (let i = 0; i < antennas.length; i++) {
    for (let j = i + 1; j < antennas.length; j++) {
      const antenna1 = antennas[i];
      const antenna2 = antennas[j];

      if (antenna1.frequency === antenna2.frequency) {
        const [antinode1, antinode2] = calculateAntinodeCoordinates(
          antenna1,
          antenna2,
        );

        if (isValidCoordinate(antinode1, grid)) {
          const key = `${antinode1.x},${antinode1.y}`;
          if (!visitedAntinodes.has(key)) {
            visitedAntinodes.add(key);
            antinodeCoordinates.push(antinode1);
          }
        }

        if (isValidCoordinate(antinode2, grid)) {
          const key = `${antinode2.x},${antinode2.y}`;
          if (!visitedAntinodes.has(key)) {
            visitedAntinodes.add(key);
            antinodeCoordinates.push(antinode2);
          }
        }
      }
    }
  }

  return antinodeCoordinates.length;
};

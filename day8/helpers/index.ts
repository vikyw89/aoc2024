export const countAntinodes = (grid: string[]): number => {
  const rows = grid.length;
  if (rows === 0) return 0;
  const cols = grid[0].length;

  const antennaPositions: Map<string, { row: number; col: number }[]> =
    new Map();

  // 1. Identify antenna positions
  for (let row = 0; row < rows; row++) {
    for (let col = 0; col < cols; col++) {
      const char = grid[row][col];
      if (/[A-Za-z]/.test(char)) {
        if (!antennaPositions.has(char)) {
          antennaPositions.set(char, []);
        }
        antennaPositions.get(char)!.push({ row, col });
      }
    }
  }

  const antinodePositions = new Set<string>();

  // 2. Iterate through all grid positions and check for antinodes
  for (let row = 0; row < rows; row++) {
    for (let col = 0; col < cols; col++) {
      for (const [_frequency, positions] of antennaPositions) {
        let inLineCount = 0;
        for (const pos of positions) {
          if (
            row === pos.row ||
            col === pos.col ||
            Math.abs(row - pos.row) === Math.abs(col - pos.col)
          ) {
            inLineCount++;
          }
        }
        if (inLineCount >= 2) {
          antinodePositions.add(`${row},${col}`);
          break; // Move to the next position if an antinode is found
        }
      }
    }
  }

  return antinodePositions.size;
};

export type Coordinate = {
  x: number;
  y: number;
};

export type Antenna = {
  frequency: string;
  coordinate: Coordinate;
};

export const getAntennas = (input: string): Antenna[] => {
  const lines = input.trim().split("\n");
  const antennas: Antenna[] = [];

  for (let y = 0; y < lines.length; y++) {
    const line = lines[y];
    for (let x = 0; x < line.length; x++) {
      const char = line[x];
      if (/[a-zA-Z0-9]/.test(char)) {
        antennas.push({
          frequency: char,
          coordinate: { x, y },
        });
      }
    }
  }
  return antennas;
};

export const calculateAntinodeCoordinates = (
  antenna1: Antenna,
  antenna2: Antenna,
): Coordinate[] => {
  const { x: x1, y: y1 } = antenna1.coordinate;
  const { x: x2, y: y2 } = antenna2.coordinate;

  const dx = x2 - x1;
  const dy = y2 - y1;

  const antinode1X = x1 - dx;
  const antinode1Y = y1 - dy;

  const antinode2X = x2 + dx;
  const antinode2Y = y2 + dy;

  return [
    { x: antinode1X, y: antinode1Y },
    { x: antinode2X, y: antinode2Y },
  ];
};

export const calculateAntinodeCoordinatesV2 = (
  antenna1: Antenna,
  antenna2: Antenna,
  grid: string[][],
): Coordinate[] => {
  const { x: x1, y: y1 } = antenna1.coordinate;
  const { x: x2, y: y2 } = antenna2.coordinate;

  const gridHeight = grid.length;
  const gridWidth = grid[0].length;

  const dx = x2 - x1;
  const dy = y2 - y1;

  const antinodes: Coordinate[] = [];
  let multiplier = 1;

  while (true) {
    const antinodeX = x1 - dx * multiplier;
    const antinodeY = y1 - dy * multiplier;

    if (
      antinodeX >= 0 &&
      antinodeX < gridWidth &&
      antinodeY >= 0 &&
      antinodeY < gridHeight
    ) {
      antinodes.push({ x: antinodeX, y: antinodeY });
    } else {
      break;
    }

    multiplier++;
  }

  multiplier = 1;

  while (true) {
    const antinodeX = x2 + dx * multiplier;
    const antinodeY = y2 + dy * multiplier;

    if (
      antinodeX >= 0 &&
      antinodeX < gridWidth &&
      antinodeY >= 0 &&
      antinodeY < gridHeight
    ) {
      antinodes.push({ x: antinodeX, y: antinodeY });
    } else {
      break;
    }

    multiplier++;
  }
  return antinodes;
};

export const isValidCoordinate = (
  coordinate: Coordinate,
  grid: string[][],
): boolean => {
  return (
    coordinate.y >= 0 &&
    coordinate.y < grid.length &&
    coordinate.x >= 0 &&
    coordinate.x < grid[0].length
  );
};

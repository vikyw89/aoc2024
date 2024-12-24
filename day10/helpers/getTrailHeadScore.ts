export type Position = {
  x: number;
  y: number;
  elevation: number;
};
export const getTrailHeadScore = ({
  topoMapMatrix,
  currentPosition,
  topMap = new Set<string>(),
}: {
  topoMapMatrix: number[][];
  currentPosition: Position;
  topMap?: Set<string>;
}): number => {
  // base case
  if (currentPosition.elevation === 9) {
    if (!topMap.has(`${currentPosition.x},${currentPosition.y}`)) {
      topMap.add(`${currentPosition.x},${currentPosition.y}`);
      return 1;
    }
    return 0;
  }

  // recursive case
  // get posible path
  const posiblePaths = getPossiblePaths({ currentPosition, topoMapMatrix });

  // get score
  let scores = 0;

  for (const path of posiblePaths) {
    scores += getTrailHeadScore({
      topoMapMatrix,
      currentPosition: path,
      topMap,
    });
  }

  return scores;
};

export const getPossiblePaths = ({
  currentPosition,
  topoMapMatrix,
}: {
  currentPosition: Position;
  topoMapMatrix: number[][];
}) => {
  const allPaths = [];

  //   allow left only if not at left edge
  if (currentPosition.x > 0) {
    allPaths.push({
      x: currentPosition.x - 1,
      y: currentPosition.y,
      elevation: topoMapMatrix[currentPosition.y][currentPosition.x - 1],
    });
  }

  //   allow right only if not at right edge
  if (currentPosition.x < topoMapMatrix[0].length - 1) {
    allPaths.push({
      x: currentPosition.x + 1,
      y: currentPosition.y,
      elevation: topoMapMatrix[currentPosition.y][currentPosition.x + 1],
    });
  }

  //   allow up only if not at top edge
  if (currentPosition.y > 0) {
    allPaths.push({
      x: currentPosition.x,
      y: currentPosition.y - 1,
      elevation: topoMapMatrix[currentPosition.y - 1][currentPosition.x],
    });
  }

  //   allow down only if not at bottom edge
  if (currentPosition.y < topoMapMatrix.length - 1) {
    allPaths.push({
      x: currentPosition.x,
      y: currentPosition.y + 1,
      elevation: topoMapMatrix[currentPosition.y + 1][currentPosition.x],
    });
  }

  //   filter path by elevation
  //   only +1 is allowed
  const elevationFilteredPaths = allPaths.filter((path) => {
    return path.elevation === currentPosition.elevation + 1;
  });

  return elevationFilteredPaths;
};

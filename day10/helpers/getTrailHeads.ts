export type TrailHead = {
  x: number;
  y: number;
};

export const getTrailHeads = ({
  topoMapMatrix,
}: {
  topoMapMatrix: number[][];
}): TrailHead[] => {
  const trailHeads = [];

  for (let y = 0; y < topoMapMatrix.length; y++) {
    for (let x = 0; x < topoMapMatrix[y].length; x++) {
      if (topoMapMatrix[y][x] === 0) {
        trailHeads.push({ x, y });
      }
    }
  }

  return trailHeads;
};

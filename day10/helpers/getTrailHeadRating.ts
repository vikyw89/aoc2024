import { getPossiblePaths, Position } from "./getTrailHeadScore.ts";

export const getTrailHeadRating = ({
  topoMapMatrix,
  currentPosition,
}: {
  topoMapMatrix: number[][];
  currentPosition: Position;
}) => {
  // base case
  if (currentPosition.elevation === 9) {
    return 1;
  }

  // recursive case
  // get posible path
  const posiblePaths = getPossiblePaths({ currentPosition, topoMapMatrix });

  // get score
  let ratings = 0;

  for (const path of posiblePaths) {
    ratings += getTrailHeadRating({
      topoMapMatrix,
      currentPosition: path,
    });
  }

  return ratings;
};

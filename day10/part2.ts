import { getTopoMapMatrix } from "./helpers/getTopoMapMatrix.ts";
import { getTrailHeadRating } from "./helpers/getTrailHeadRating.ts";
import { getTrailHeads } from "./helpers/getTrailHeads.ts";

export const part2 = (filePath: string): number => {
  const topoMapMatrix = getTopoMapMatrix(filePath);

  const trailHeads = getTrailHeads({ topoMapMatrix });
  let ratings = 0;

  for (const trailHead of trailHeads) {
    ratings += getTrailHeadRating({
      topoMapMatrix,
      currentPosition: {
        x: trailHead.x,
        y: trailHead.y,
        elevation: topoMapMatrix[trailHead.y][trailHead.x],
      },
    });
  }

  return ratings;
};

import { getTopoMapMatrix } from "./helpers/getTopoMapMatrix.ts";
import { getTrailHeads } from "./helpers/getTrailHeads.ts";
import { getTrailHeadScore } from "./helpers/getTrailHeadScore.ts";

export const part1 = (filePath: string): number => {
  const topoMapMatrix = getTopoMapMatrix(filePath);

  const trailHeads = getTrailHeads({ topoMapMatrix });
  let score = 0;

  for (const trailHead of trailHeads) {
    score += getTrailHeadScore({
      topoMapMatrix,
      currentPosition: {
        x: trailHead.x,
        y: trailHead.y,
        elevation: topoMapMatrix[trailHead.y][trailHead.x],
      },
      topMap: new Set<string>(),
    });
  }

  return score;
};

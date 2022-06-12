import type { GaltonBall } from './GaltonBall';

export interface GaltonBoard {
  DecisionDepth: number;
  Columns: Array<Array<GaltonBall>>;
}

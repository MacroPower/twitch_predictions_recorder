import Model from "./Model";
import Predictor from "./Predictor";

interface Outcome {
  id: string;
  color: string;
  title: string;
  badge_version: string;
  badge_set_id: string;
  total_points: number;
  total_users: number;
  result_type: string;
  top_predictors: Predictor[];
}

class Outcome extends Model {}

export function getOutcomeSum(outcomes: Outcome[]) {
  const sum = {
    title: "Total",
    total_points: 0,
    total_users: 0,
  };
  outcomes.forEach((o) => {
    sum.total_points += o.total_points;
    sum.total_users += o.total_users;
  });
  return sum;
}

export default Outcome;

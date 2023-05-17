import Model from "./Model";

interface Outcome {
  color: string;
  title: string;
  badge_version: string;
  badge_set_id: string;
  total_points: number;
  total_users: number;
  result_type: string;
}

class Outcome extends Model {}

export default Outcome;

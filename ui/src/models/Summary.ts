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

interface Summary {
  id: string;
  channel_name: string;
  prediction_window_seconds: number;
  title: string;
  status: string;
  outcomes: Outcome[];
}

class Summary extends Model {
  resource() {
    return "api/v1/summary";
  }
}

export default Summary;

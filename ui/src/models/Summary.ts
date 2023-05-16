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
  timestamp: string;
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

  getDate() {
    // Returns a Date object from the timestamp string
    return new Date(this.timestamp);
  }

  getOutcomes(): Outcome[] {
    this.outcomes.sort((a, b) => {
      // Sorts so that the outcomes are in order of badge_version alphabetically
      if (a.badge_version < b.badge_version) {
        return -1;
      }
      if (a.badge_version > b.badge_version) {
        return 1;
      }
      return 0;
    });
    return this.outcomes;
  }

  outcomeSum() {
    // Returns an Outcome object with the sum of all outcomes
    const sum: Outcome = {
      color: "",
      title: "Total",
      badge_version: "",
      badge_set_id: "",
      total_points: 0,
      total_users: 0,
      result_type: "",
    };
    this.outcomes.forEach((outcome) => {
      sum.total_points += outcome.total_points;
      sum.total_users += outcome.total_users;
    });
    return sum;
  }
}

export default Summary;

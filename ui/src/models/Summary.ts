import Model from "./Model";
import Outcome from "./Outcome";

interface Summary {
  id: string;
  timestamp: string;
  channel_name: string;
  created_at: string;
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
    return new Date(this.created_at);
  }

  getEndDate() {
    const created = this.getDate();
    created.setSeconds(created.getSeconds() + this.prediction_window_seconds);
    return created;
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
    const sum = {
      title: "Total",
      total_points: 0,
      total_users: 0,
    };
    this.outcomes.forEach((outcome) => {
      sum.total_points += outcome.total_points;
      sum.total_users += outcome.total_users;
    });
    return sum;
  }
}

export default Summary;

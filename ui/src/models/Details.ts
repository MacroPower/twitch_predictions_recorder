import Model from "./Model";
import Outcome from "./Outcome";

interface Details {
  id: string;
  channel_name: string;
  created_at: string;
  prediction_window_seconds: number;
  title: string;
  event_series: EventSeries[];
}

interface EventSeries {
  id: string;
  timestamp: string;
  status: string;
  outcomes: Outcome[];
}

export interface DetailsSeries {
  details: Outcome;
  values: Array<{
    timestamp: Date;
    status: string;
    points: number;
    users: number;
  }>;
}

class Details extends Model {
  resource() {
    return "api/v1/details";
  }

  sort() {
    this.event_series.sort((a, b) => {
      if (a.timestamp < b.timestamp) {
        return -1;
      }
      if (a.timestamp > b.timestamp) {
        return 1;
      }
      return 0;
    });
  }

  getTimeSeries(): DetailsSeries[] {
    const outcomes: {
      [Key: string]: DetailsSeries;
    } = {};

    this.event_series.forEach((e) => {
      e.outcomes.forEach((o) => {
        if (!(o.title in outcomes)) {
          outcomes[o.title] = {
            details: o,
            values: [],
          };
        }

        const ts = new Date(e.timestamp);
        outcomes[o.title].values.push({
          timestamp: ts,
          status: e.status,
          points: o.total_points,
          users: o.total_users,
        });
      });
    });

    return Object.values(outcomes);
  }
}

export default Details;

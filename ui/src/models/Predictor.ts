import Model from "./Model";
import User from "./User";

interface Predictor {
  user: User;
  points: number;
}

class Predictor extends Model {}

export default Predictor;

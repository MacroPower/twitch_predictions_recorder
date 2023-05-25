import Model from "./Model";

interface User {
  user_id: string;
  user_display_name: string;
}

class User extends Model {}

export default User;

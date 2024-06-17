export interface IPost {
  uuid: string;
  content: string;
  username: string;
  user_uuid: string;
  visibility_type_id: number;
  update_at: string; // Assuming update_at is a string representation of a date
}

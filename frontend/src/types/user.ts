export interface User {
  id: string;
  name: string;
  username: string;
  email: string;
  created_at: string;
  updated_at: string;
}

export interface UserUpdateRequest {
  name: string;
  username: string;
  email: string;
}

export interface UserCreateRequest {
  name: string;
  username: string;
  email: string;
  password: string;
}

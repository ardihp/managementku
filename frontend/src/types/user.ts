export interface User {
  id: string;
  name: string;
  username: string;
  email: string;
  role: string;
  created_at: string;
  updated_at: string;
}

export interface UserUpdateRequest {
  name: string;
  username: string;
  email: string;
  role: string;
}

export interface UserCreateRequest {
  name: string;
  username: string;
  email: string;
  role: string;
  password: string;
}

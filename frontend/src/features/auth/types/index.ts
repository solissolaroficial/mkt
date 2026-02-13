export interface LoginCredentials {
  email: string;
  password: string;
}

export interface LoginResponse {
  access_token: string;
  refresh_token: string;
  user: {
    id: string;
    email: string;
    name: string;
    role: string;
    profile_photo_key?: string;
    profile_photo_url?: string;
  };
}

export interface AuthUser {
  id: string;
  email: string;
  name: string;
  role: string;
  profilePhotoKey?: string;
  profilePhotoURL?: string;
}

export interface AuthError {
  message: string;
  code?: string;
}
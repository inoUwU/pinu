import api from "@/lib/api/api";

type LoginPayload = {
  userId: string;
  password: string;
};

type LoginResponse = {
  access_token: string;
  refresh_token: string;
};

export async function login(payload: LoginPayload) {
  return api.post("auth/login", { json: payload }).json<LoginResponse>();
}

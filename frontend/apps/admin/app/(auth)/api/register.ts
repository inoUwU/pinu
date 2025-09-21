import api from "@/lib/api/api";

type RegisterPayload = {
  userId: string;
  password: string;
};

type RegisterResponse = {
  access_token: string;
  refresh_token: string;
};

export async function register(payload: RegisterPayload) {
  return api.post("auth/register", { json: payload }).json<RegisterResponse>();
}

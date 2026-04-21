export const ACCESS_TOKEN_STORAGE_KEY = "access_token";
export const REFRESH_TOKEN_STORAGE_KEY = "refresh_token";

type AuthTokens = {
  access_token: string;
  refresh_token: string;
};

export const getAccessToken = () => {
  if (typeof window === "undefined") {
    return null;
  }

  return localStorage.getItem(ACCESS_TOKEN_STORAGE_KEY);
};

export const persistAuthTokens = (tokens: AuthTokens) => {
  localStorage.setItem(ACCESS_TOKEN_STORAGE_KEY, tokens.access_token);
  localStorage.setItem(REFRESH_TOKEN_STORAGE_KEY, tokens.refresh_token);
};

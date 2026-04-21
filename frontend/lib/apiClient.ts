import ky from "ky";
import { getAccessToken } from "@/lib/auth/token-storage";

const apiClient = ky.create({
  // 環境変数からAPIのベースURLを設定
  prefixUrl: process.env.NEXT_PUBLIC_API_BASE_URL,
  hooks: {
    // リクエストが送信される前に実行されるフック
    beforeRequest: [
      request => {
        // クライアントサイドでのみトークンを取得・設定
        const token = getAccessToken();
        if (token) {
          request.headers.set("Authorization", `Bearer ${token}`);
        }
      },
    ],
    // レスポンスを受け取った後に実行されるフック
    afterResponse: [
      async (_request, _options, response) => {
        // 401 Unauthorizedエラーの場合、ログインページへリダイレクトするなどの共通処理
        if (response.status === 401 && typeof window !== "undefined") {
          console.error("Authentication error. Redirecting to login page.");
          // 実際のアプリケーションではここでリダイレクト処理を行う
          // window.location.href = '/login';
        }
        return response;
      },
    ],
  },
});

export default apiClient;

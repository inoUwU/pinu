import { HttpResponse, http } from "msw";

// 単純な GET エンドポイント
export const handlers = [
  http.get("/api/users", () => {
    // 必要なら async にして request.body など見れる
    return HttpResponse.json([
      { id: 1, name: "Alice" },
      { id: 2, name: "Bob" },
    ]);
  }),

  // POST の例
  http.post("/api/login", async ({ request }) => {
    const body = (await request.json()) as {
      username: string;
      password: string;
    };
    const { username, password } = body;
    // 単純な判定の例
    if (username === "admin" && password === "password") {
      return HttpResponse.json({ token: "fake-jwt-token" });
    }
    return new Response("Unauthorized", { status: 401 });
  }),
];

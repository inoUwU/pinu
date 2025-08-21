"use client";

import useSWR from "swr";
import apiClient from "@/lib/apiClient";

// SWRに渡すfetcher関数。apiClientがkyインスタンスなので、.json()でパースされたデータが返る
const fetcher = (url: string) => apiClient.get(url).json();

// ログインしているユーザー情報を取得するためのカスタムフック
export const useMe = () => {
  // SWRのキーには、apiClientのprefixUrlに続くパスを指定
  const { data, error, isLoading, mutate } = useSWR("/auth/me", fetcher);

  return {
    user: data, // APIから取得したユーザーデータ
    error, // エラーオブジェクト
    isLoading, // データ取得中の状態
    isLoggedIn: !error && data, // ログイン状態の簡易的な判定
    mutate, // キャッシュの手動更新用関数
  };
};

import type { Category, MenuItem } from "./types";

/**
 * ダミーデータ - カテゴリー
 */
export const dummyCategories: Category[] = [
  {
    id: "cat-1",
    name: "麺類",
    displayOrder: 1,
    imageUrl: "https://placehold.jp/300x200.png",
  },
  {
    id: "cat-2",
    name: "ご飯物",
    displayOrder: 2,
    imageUrl: "https://placehold.jp/300x200.png",
  },
  {
    id: "cat-3",
    name: "一品料理",
    displayOrder: 3,
    imageUrl: "https://placehold.jp/300x200.png",
  },
];

/**
 * ダミーデータ - メニューアイテム
 */
export const dummyMenuItems: MenuItem[] = [
  // 麺類
  {
    id: "menu-1",
    name: "ラーメン",
    description: "醤油ベースの定番ラーメン",
    price: 800,
    isSoldOut: false,
    categoryId: "cat-1",
    displayOrder: 1,
    options: [
      { id: "opt-1", name: "大盛り", price: 100, description: "麺量1.5倍" },
      {
        id: "opt-2",
        name: "チャーシュー追加",
        price: 200,
        description: "チャーシュー3枚追加",
      },
    ],
  },
  {
    id: "menu-2",
    name: "つけ麺",
    description: "濃厚豚骨つけ汁の太麺",
    price: 900,
    isSoldOut: false,
    categoryId: "cat-1",
    displayOrder: 2,
    options: [
      { id: "opt-3", name: "大盛り", price: 100, description: "麺量1.5倍" },
      { id: "opt-4", name: "特盛り", price: 200, description: "麺量2倍" },
    ],
  },
  {
    id: "menu-3",
    name: "冷やし中華",
    description: "夏限定の冷たい中華そば",
    price: 850,
    isSoldOut: true,
    categoryId: "cat-1",
    displayOrder: 3,
    options: [],
  },

  // ご飯物
  {
    id: "menu-4",
    name: "チャーハン",
    description: "パラパラに炒めた定番チャーハン",
    price: 700,
    isSoldOut: false,
    categoryId: "cat-2",
    displayOrder: 1,
    options: [
      { id: "opt-5", name: "大盛り", price: 100, description: "ご飯量1.5倍" },
    ],
  },
  {
    id: "menu-5",
    name: "親子丼",
    description: "ふわとろ卵の親子丼",
    price: 750,
    isSoldOut: false,
    categoryId: "cat-2",
    displayOrder: 2,
    options: [
      { id: "opt-6", name: "大盛り", price: 100, description: "ご飯量1.5倍" },
      {
        id: "opt-7",
        name: "みそ汁セット",
        price: 150,
        description: "みそ汁付き",
      },
    ],
  },

  // 一品料理
  {
    id: "menu-6",
    name: "餃子",
    description: "手作り焼き餃子（6個）",
    price: 450,
    isSoldOut: false,
    categoryId: "cat-3",
    displayOrder: 1,
    options: [
      {
        id: "opt-8",
        name: "水餃子に変更",
        price: 0,
        description: "茹で餃子に変更",
      },
    ],
  },
  {
    id: "menu-7",
    name: "唐揚げ",
    description: "ジューシー鶏唐揚げ（5個）",
    price: 500,
    isSoldOut: false,
    categoryId: "cat-3",
    displayOrder: 2,
    options: [
      {
        id: "opt-9",
        name: "マヨネーズ追加",
        price: 50,
        description: "マヨネーズソース付き",
      },
    ],
  },
];

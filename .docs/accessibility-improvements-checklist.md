# アクセシビリティ改善チェックリスト

このドキュメントは、フロントエンドアクセシビリティ監査レポートに基づく、実装可能な改善タスクのチェックリストです。

## 📋 即座に対応すべき項目（高優先度）

### セマンティックHTML

- [ ] **メイン要素の追加**
  - [ ] `frontend/apps/client/app/category/page.tsx` - `<div>`を`<main>`に変更
  - [ ] `frontend/apps/client/app/order/page.tsx` - `<div>`を`<main>`に変更
  - [ ] `frontend/apps/client/app/[categoryid]/page.tsx` - `<div>`を`<main>`に変更
  - [ ] `frontend/apps/client/app/checkout/page.tsx` - `<div>`を`<main>`に変更

- [ ] **ナビゲーション要素の追加**
  - [ ] `frontend/apps/client/app/components/TabBar.tsx` - `<nav>`でラップし、`aria-label`を追加

### ARIA属性

- [ ] **ボタンのラベル追加**
  - [ ] `frontend/apps/client/app/components/PageHeader.tsx` - 戻るボタンに`aria-label="前のページに戻る"`を追加
  - [ ] `frontend/apps/admin/components/Navbar.tsx` - テーマ切り替えボタンのスクリーンリーダーテキスト確認
  - [ ] `frontend/apps/admin/app/(authenticated)/menu/components/MenuCard.tsx` - 編集ボタンに`aria-label="メニューを編集"`を追加

- [ ] **アイコンのaria-hidden追加**
  - [ ] 全コンポーネント - テキストラベルと併用されるアイコンに`aria-hidden="true"`を追加
  - 対象ファイル例:
    - `frontend/apps/client/app/components/TabBar.tsx` (Utensils, ShoppingCart, JapaneseYen)
    - `frontend/apps/admin/components/Navbar.tsx` (Coins, UtensilsCrossed, Sun, Moon)
    - `frontend/apps/admin/components/Sidebar.tsx` (全アイコン)

### キーボードナビゲーション

- [ ] **クリック可能なカードへのキーボード対応**
  - [ ] `frontend/apps/client/app/category/components/CategoryCard.tsx`
    ```tsx
    - tabIndex={0}を追加
    - role="button"を追加
    - aria-label={`${category.name}カテゴリーを表示`}を追加
    - onKeyDownハンドラーでEnter/Spaceキー対応を追加
    ```
  - [ ] `frontend/apps/admin/app/(authenticated)/bill/components/TableCard.tsx` - 同様の対応

- [ ] **スキップリンクの追加**
  - [ ] `frontend/apps/client/app/layout.tsx`
  - [ ] `frontend/apps/admin/app/layout.tsx`
  ```tsx
  <a 
    href="#main-content" 
    className="sr-only focus:not-sr-only focus:absolute focus:top-4 focus:left-4 focus:z-50 focus:bg-primary focus:text-primary-foreground focus:px-4 focus:py-2 focus:rounded"
  >
    メインコンテンツへスキップ
  </a>
  ```

### 画像と代替テキスト

- [ ] **装飾的画像の適切なマーク**
  - [ ] `frontend/apps/client/app/category/components/CategoryCard.tsx` - alt属性を意味のあるものに変更
  - [ ] `frontend/apps/admin/app/(authenticated)/menu/components/MenuCard.tsx` - 装飾的な画像に`alt=""`と`aria-hidden="true"`を追加

- [ ] **意味のある代替テキスト**
  - [ ] カテゴリー画像: `alt={`${category.name}のカテゴリー画像`}`
  - [ ] メニュー画像: `alt={menuItem.name}`（または装飾目的なら`alt=""`）

### チャートのアクセシビリティ

- [ ] **グラフへのaria-label追加**
  - [ ] `frontend/apps/admin/components/charts/DailySalesChart.tsx`
  - [ ] `frontend/apps/admin/components/charts/CategorySalesChart.tsx`
  - [ ] `frontend/apps/admin/components/charts/TopMenusChart.tsx`
  ```tsx
  <div role="img" aria-label="[グラフの説明]">
    <ResponsiveContainer>...</ResponsiveContainer>
  </div>
  ```

- [ ] **データテーブル代替の追加**
  - [ ] 各チャートコンポーネントに折りたたみ可能なテーブルビューを追加

---

## 📊 1-2週間以内に対応すべき項目（中優先度）

### フォームのアクセシビリティ

- [ ] **エラー処理の改善**
  - [ ] `frontend/apps/admin/app/(auth)/login/page.tsx`
    - エラー発生時のフォーカス管理を追加
    - エラーメッセージに`aria-live="assertive"`を追加
  - [ ] `frontend/apps/admin/app/(authenticated)/user/components/form.tsx` - 同様の対応

- [ ] **必須フィールドの明示**
  - [ ] 全フォームコンポーネント - 必須フィールドに視覚的マーカーと`aria-required="true"`を追加

### 動的コンテンツ

- [ ] **ローディング状態の改善**
  - [ ] `frontend/apps/client/app/page.tsx`
  ```tsx
  if (isLoading) return (
    <div role="status" aria-live="polite" aria-busy="true">
      <span>読み込み中...</span>
    </div>
  );
  ```
  - [ ] `frontend/apps/admin/components/charts/KPISummaryCards.tsx` - 同様の対応

- [ ] **エラー状態の改善**
  - [ ] `frontend/apps/client/app/page.tsx`
  ```tsx
  if (error) return (
    <div role="alert" aria-live="assertive">
      <p>データの読み込みに失敗しました。</p>
    </div>
  );
  ```

- [ ] **ライブリージョンの追加**
  - [ ] `frontend/apps/admin/app/(authenticated)/operation/components/OrderCard.tsx` - 経過時間表示に`aria-live="polite"`を追加

### ビジュアルデザイン

- [ ] **色以外の情報伝達手段の追加**
  - [ ] `frontend/apps/admin/app/(authenticated)/operation/components/OrderCard.tsx`
    - 時間表示にアイコンを追加
    - スクリーンリーダー用の状態説明を追加

- [ ] **カラーコントラストの検証と修正**
  - [ ] 全コンポーネント - Lighthouse/axeでコントラスト比を検証
  - [ ] 必要に応じて色を調整（最小4.5:1を確保）

### 見出し構造

- [ ] **見出し階層の整理**
  - [ ] 各ページの見出し構造を`<h1>` → `<h2>` → `<h3>`と論理的に整理
  - 対象: 全ページコンポーネント

### ダイアログ

- [ ] **ダイアログコンテンツの改善**
  - [ ] `frontend/apps/admin/app/(authenticated)/bill/components/BillDialog.tsx`
    - DialogDescriptionを日本語の適切な内容に変更
    - 重要なアクションに`role="alertdialog"`の使用を検討

---

## 🔧 継続的改善項目（低優先度）

### テスト自動化

- [ ] **eslint-plugin-jsx-a11yの導入**
  ```bash
  pnpm add -D eslint-plugin-jsx-a11y
  ```
  - [ ] ESLint設定に追加
  - [ ] 既存コードのリンティング実行と修正

- [ ] **axe-coreの導入**
  ```bash
  pnpm add -D @axe-core/react jest-axe
  ```
  - [ ] テストセットアップ
  - [ ] 主要コンポーネントのテスト追加

- [ ] **Lighthouse CIの導入**
  - [ ] GitHub Actionsワークフロー作成
  - [ ] アクセシビリティスコアの閾値設定

### 多言語対応

- [ ] **UIテキストの日本語統一**
  - [ ] `frontend/apps/admin/components/Sidebar.tsx` - "Home", "User", "Menu"などを日本語化
  - [ ] `frontend/apps/admin/app/(auth)/login/page.tsx` - "Welcome back"を日本語化

- [ ] **i18nの検討**（将来的な拡張性のため）
  - [ ] next-i18nextまたは類似ライブラリの評価
  - [ ] 実装計画の策定

### モバイルアクセシビリティ

- [ ] **タッチターゲットサイズの検証**
  - [ ] 全インタラクティブ要素が44x44px以上であることを確認
  - [ ] 必要に応じてpaddingを調整

### ドキュメント

- [ ] **アクセシビリティガイドライン作成**
  - [ ] 新規コンポーネント作成時のチェックリスト
  - [ ] コードレビュー時の確認項目
  - [ ] 開発チーム向けのベストプラクティス

- [ ] **コンポーネントのアクセシビリティドキュメント**
  - [ ] 各主要コンポーネントのアクセシビリティ機能を文書化
  - [ ] 使用例とベストプラクティスの提供

---

## 🛠️ 実装支援スニペット

### クリック可能なカードの完全な実装例

```tsx
// CategoryCard.tsx の改善版
import { Card, CardContent, CardTitle } from "@workspace/ui/components/card";
import { useRouter } from "nextjs-toploader/app";
import type { KeyboardEvent } from "react";

export default function CategoryCard({
  category,
  showImage = false,
}: CategoryCardProps) {
  const router = useRouter();

  const handleClick = () => {
    router.push(`/menu/${category.menuId}`);
  };

  const handleKeyDown = (e: KeyboardEvent<HTMLDivElement>) => {
    if (e.key === 'Enter' || e.key === ' ') {
      e.preventDefault();
      handleClick();
    }
  };

  return (
    <Card 
      onClick={handleClick}
      onKeyDown={handleKeyDown}
      tabIndex={0}
      role="button"
      aria-label={`${category.name}カテゴリーを表示`}
      className="cursor-pointer focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2"
    >
      <CardContent>
        <div className='flex gap-2'>
          {showImage && category.imageUrl && (
            <img
              src={category.imageUrl}
              alt={`${category.name}のカテゴリー画像`}
              className='w-24 h-24 object-cover rounded-md mb-2'
            />
          )}
          <CardTitle className='text-2xl break-all'>{category.name}</CardTitle>
        </div>
      </CardContent>
    </Card>
  );
}
```

### ローディング/エラー状態の実装例

```tsx
// page.tsx の改善版
export default function Page() {
  const { data, error, isLoading } = useSWR("/api/users", fetcher);

  if (error) {
    return (
      <div role="alert" aria-live="assertive" className="p-4">
        <p className="text-destructive">
          データの読み込みに失敗しました。ページを再読み込みしてください。
        </p>
      </div>
    );
  }

  if (isLoading) {
    return (
      <div role="status" aria-live="polite" aria-busy="true" className="p-4">
        <span>読み込み中...</span>
        <span className="sr-only">データを読み込んでいます。しばらくお待ちください。</span>
      </div>
    );
  }

  return <div>hello {data[0].name}!</div>;
}
```

### チャートの代替表現の実装例

```tsx
// DailySalesChart.tsx の改善版
export function DailySalesChart({ data, isLoading }: DailySalesChartProps) {
  // ... 既存のコード ...

  return (
    <Card className='col-span-2'>
      <CardHeader>
        <CardTitle>日次売上トレンド（過去30日）</CardTitle>
      </CardHeader>
      <CardContent>
        <div 
          role="img" 
          aria-label={`日次売上トレンドグラフ：過去30日間の注文数と売上額の推移を表示。最新の売上は${chartData[chartData.length - 1]?.revenue.toLocaleString()}円です。`}
          className='h-[300px]'
        >
          <ResponsiveContainer width='100%' height='100%'>
            {/* Chart content */}
          </ResponsiveContainer>
        </div>
        
        {/* データテーブル代替 */}
        <details className="mt-4">
          <summary className="cursor-pointer text-sm text-muted-foreground hover:text-foreground">
            データをテーブルで表示
          </summary>
          <div className="mt-2 overflow-auto">
            <table className="w-full text-sm">
              <thead>
                <tr>
                  <th className="text-left p-2">日付</th>
                  <th className="text-right p-2">注文数</th>
                  <th className="text-right p-2">売上</th>
                </tr>
              </thead>
              <tbody>
                {chartData.map((item, index) => (
                  <tr key={index} className="border-t">
                    <td className="p-2">{item.day}</td>
                    <td className="text-right p-2">{item.ordersCount}</td>
                    <td className="text-right p-2">¥{item.revenue.toLocaleString()}</td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        </details>
      </CardContent>
    </Card>
  );
}
```

---

## 📈 進捗管理

### 週次チェックポイント

**第1週:**
- [ ] セマンティックHTML改善
- [ ] ボタンとリンクのARIA属性追加
- [ ] 画像の代替テキスト整備

**第2週:**
- [ ] キーボードナビゲーション対応
- [ ] チャートのアクセシビリティ改善
- [ ] フォームエラー処理改善

**第3週:**
- [ ] 動的コンテンツの改善
- [ ] カラーコントラスト検証
- [ ] テストツール導入

**第4週:**
- [ ] 包括的テストの実行
- [ ] ドキュメント整備
- [ ] 残課題の洗い出し

---

## 🎯 成功指標

- [ ] Lighthouse アクセシビリティスコア: 90以上
- [ ] axe-core 自動テスト: 0件の違反
- [ ] eslint-plugin-jsx-a11y: 0件のエラー
- [ ] キーボードのみで全機能が操作可能
- [ ] スクリーンリーダーで全コンテンツがアクセス可能

---

**最終更新日:** 2026年1月13日

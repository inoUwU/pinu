# フロントエンドアクセシビリティ監査レポート

## 実施日
2026年1月13日

## 対象プロジェクト
- **クライアントアプリ** (`frontend/apps/client`)
- **管理アプリ** (`frontend/apps/admin`)
- **共通UIコンポーネント** (`frontend/packages/ui`)

## エグゼクティブサマリー

本レポートは、Pinuプロジェクトのフロントエンド側のアクセシビリティに関する包括的な監査結果をまとめたものです。WCAG 2.1レベルAAを基準として評価を行いました。

### 総合評価

| カテゴリー | 評価 | 主な問題数 |
|----------|------|----------|
| セマンティックHTML | ⚠️ 要改善 | 8件 |
| キーボードナビゲーション | ⚠️ 要改善 | 6件 |
| ARIA属性 | ⚠️ 要改善 | 10件 |
| フォーカス管理 | ⚠️ 要改善 | 5件 |
| カラーコントラスト | ✅ 概ね良好 | 2件 |
| 代替テキスト | ❌ 要対応 | 7件 |
| フォーム | ⚠️ 要改善 | 4件 |

---

## 1. セマンティックHTML・マークアップの問題

### 🔴 重大度：高
#### 1.1 `<main>` 要素の欠如
**影響するファイル:**
- `frontend/apps/client/app/category/page.tsx`
- `frontend/apps/client/app/order/page.tsx`
- `frontend/apps/client/app/[categoryid]/page.tsx`

**問題点:**
メインコンテンツ領域が`<div>`でマークアップされており、`<main>`要素が使用されていません。スクリーンリーダーユーザーがメインコンテンツに直接ジャンプできません。

**推奨対応:**
```tsx
// Before
<div className='basis-11/12 grow overflow-auto bg-gray-50 pt-2'>
  <CategorySection />
</div>

// After
<main className='basis-11/12 grow overflow-auto bg-gray-50 pt-2'>
  <CategorySection />
</main>
```

### 🟡 重大度：中
#### 1.2 ナビゲーション要素のセマンティック不足
**影響するファイル:**
- `frontend/apps/client/app/components/TabBar.tsx`

**問題点:**
タブバーが`<nav>`要素でラップされておらず、ナビゲーション目的であることが明示されていません。

**推奨対応:**
```tsx
<nav aria-label="メインナビゲーション">
  <ToggleGroup type='single' size='lg' className='w-lvw h-full p-0 m-0'>
    {/* タブアイテム */}
  </ToggleGroup>
</nav>
```

#### 1.3 見出し階層の問題
**影響するファイル:**
- `frontend/apps/admin/app/(authenticated)/bill/components/TableCard.tsx`

**問題点:**
`<h2>`がカード内に直接配置されていますが、ページ全体の見出し階層が明確ではありません。

**推奨対応:**
ページ全体の見出し構造を`<h1>` → `<h2>` → `<h3>`と論理的に設計する。

---

## 2. ARIA属性の問題

### 🔴 重大度：高
#### 2.1 ボタンのaria-labelの欠如
**影響するファイル:**
- `frontend/apps/client/app/components/PageHeader.tsx` (BackButton)
- `frontend/apps/admin/components/Navbar.tsx` (テーマ切り替えボタン)

**問題点:**
アイコンのみのボタンに明確なラベルがありません。`<span className='sr-only'>`は存在しますが、全てのボタンに適用されていません。

**推奨対応:**
```tsx
// Before
<Button onClick={onClick} className='p-6 w-20 relative' variant='ghost'>
  <ChevronLeft className='absolute left-3 top-1/2 transform -translate-y-1/2' />
</Button>

// After
<Button 
  onClick={onClick} 
  className='p-6 w-20 relative' 
  variant='ghost'
  aria-label="前のページに戻る"
>
  <ChevronLeft className='absolute left-3 top-1/2 transform -translate-y-1/2' />
</Button>
```

#### 2.2 リンクとボタンの混在
**影響するファイル:**
- `frontend/apps/client/app/components/TabBar.tsx`
- `frontend/apps/client/app/category/components/CategoryCard.tsx`

**問題点:**
ToggleGroupItem内にLinkを配置しており、インタラクティブ要素がネストされています。これはアクセシビリティとセマンティクスの問題を引き起こします。

**推奨対応:**
```tsx
// Before
<ToggleGroupItem value='category'>
  <Utensils className='w-6 h-6' />
  <Link href='/category'>メニュー</Link>
</ToggleGroupItem>

// After
<ToggleGroupItem value='category' asChild>
  <Link href='/category'>
    <Utensils className='w-6 h-6' aria-hidden="true" />
    <span>メニュー</span>
  </Link>
</ToggleGroupItem>
```

### 🟡 重大度：中
#### 2.3 ダイアログのアクセシビリティ
**影響するファイル:**
- `frontend/apps/admin/app/(authenticated)/bill/components/BillDialog.tsx`

**問題点:**
DialogDescriptionに英語のプレースホルダーテキストが残っています。また、破壊的アクションに対する適切な警告が不足しています。

**推奨対応:**
```tsx
<DialogDescription>
  この操作は取り消せません。{table}番テーブルの会計処理を完了します。
</DialogDescription>
```

さらに、重要なアクション（会計確定など）には`role="alertdialog"`の使用を検討。

#### 2.4 ライブリージョンの欠如
**影響するファイル:**
- `frontend/apps/admin/app/(authenticated)/operation/components/OrderCard.tsx`

**問題点:**
経過時間が動的に更新されますが、スクリーンリーダーにその変更が通知されません。

**推奨対応:**
```tsx
<span 
  className={`font-bold ${timeColorClass}`}
  aria-live="polite"
  aria-atomic="true"
>
  {elapsedMinutes}分
</span>
```

#### 2.5 チャートのアクセシビリティ
**影響するファイル:**
- `frontend/apps/admin/components/charts/DailySalesChart.tsx`
- `frontend/apps/admin/components/charts/CategorySalesChart.tsx`
- `frontend/apps/admin/components/charts/TopMenusChart.tsx`

**問題点:**
グラフが視覚表現のみで、代替テキストやデータテーブルが提供されていません。

**推奨対応:**
```tsx
<div role="img" aria-label="日次売上トレンドグラフ：過去30日間の注文数と売上額の推移を表示">
  <ResponsiveContainer width='100%' height='100%'>
    {/* chart content */}
  </ResponsiveContainer>
</div>
{/* グラフの下に要約テキストまたは折りたたみ可能なテーブルを追加 */}
<details className="mt-4">
  <summary>データテーブルで表示</summary>
  <table>
    {/* テーブル形式でデータを提供 */}
  </table>
</details>
```

---

## 3. キーボードナビゲーションの問題

### 🔴 重大度：高
#### 3.1 クリック可能なカードへのキーボードアクセス不足
**影響するファイル:**
- `frontend/apps/client/app/category/components/CategoryCard.tsx`
- `frontend/apps/admin/app/(authenticated)/bill/components/TableCard.tsx`

**問題点:**
`onClick`ハンドラーを持つCardコンポーネントがキーボードでアクセスできません。

**推奨対応:**
```tsx
// Before
<Card onClick={handleClick}>
  <CardContent>
    {/* content */}
  </CardContent>
</Card>

// After
<Card 
  onClick={handleClick}
  onKeyDown={(e) => {
    if (e.key === 'Enter' || e.key === ' ') {
      e.preventDefault();
      handleClick();
    }
  }}
  tabIndex={0}
  role="button"
  aria-label={`${category.name}カテゴリーを表示`}
>
  <CardContent>
    {/* content */}
  </CardContent>
</Card>
```

### 🟡 重大度：中
#### 3.2 フォーカスインジケーターの視認性
**影響するファイル:**
- `frontend/packages/ui/src/components/button.tsx`
- `frontend/packages/ui/src/components/input.tsx`

**問題点:**
フォーカスリングのスタイルは定義されていますが、一部のコンポーネントで視認性が低い可能性があります。

**推奨対応:**
カラーコントラストのテストを実施し、必要に応じてフォーカスリングの太さや色を調整。

#### 3.3 スキップリンクの欠如
**影響するファイル:**
- `frontend/apps/client/app/layout.tsx`
- `frontend/apps/admin/app/layout.tsx`

**問題点:**
メインコンテンツへスキップするリンクがありません。

**推奨対応:**
```tsx
<body>
  <a 
    href="#main-content" 
    className="sr-only focus:not-sr-only focus:absolute focus:top-4 focus:left-4 focus:z-50 focus:bg-primary focus:text-primary-foreground focus:px-4 focus:py-2 focus:rounded"
  >
    メインコンテンツへスキップ
  </a>
  {/* rest of content */}
</body>
```

---

## 4. フォームのアクセシビリティ問題

### 🟡 重大度：中
#### 4.1 フォームのエラー処理
**影響するファイル:**
- `frontend/apps/admin/app/(auth)/login/page.tsx`
- `frontend/apps/admin/app/(authenticated)/user/components/form.tsx`

**問題点:**
エラーメッセージは表示されますが、フォーカス管理やaria-live通知が不足しています。

**推奨対応:**
```tsx
// エラー発生時にフォーカスを最初のエラーフィールドに移動
const firstErrorField = Object.keys(errors)[0];
if (firstErrorField) {
  document.querySelector(`[name="${firstErrorField}"]`)?.focus();
}

// エラーメッセージ領域にaria-live追加
<div aria-live="assertive" aria-atomic="true">
  {error && <p className="text-destructive">{error}</p>}
</div>
```

#### 4.2 必須フィールドの明示
**影響するファイル:**
- 全フォームコンポーネント

**問題点:**
必須フィールドの視覚的マーカー（*など）とaria-required属性が一部欠如。

**推奨対応:**
Zod schemaで`.required()`を使用している場合、自動的に`aria-required="true"`を追加するラッパーを作成。

---

## 5. 画像とメディアの代替テキスト

### 🔴 重大度：高
#### 5.1 装飾的画像にaria-hiddenがない
**影響するファイル:**
- `frontend/apps/client/app/category/components/CategoryCard.tsx`
- `frontend/apps/admin/app/(authenticated)/menu/components/MenuCard.tsx`

**問題点:**
プレースホルダー画像や装飾的な画像にalt属性はありますが、スクリーンリーダーで読み上げられる必要のない画像が適切にマークされていません。

**推奨対応:**
```tsx
// 意味のある画像
<img
  src={category.imageUrl}
  alt={`${category.name}のカテゴリー画像`}
  className='w-24 h-24 object-cover rounded-md mb-2'
/>

// 装飾的な画像
<img
  src={menuItem.imageUrl}
  alt=""  // 空のalt属性で装飾画像を示す
  aria-hidden="true"
  className='w-full h-20 object-cover rounded-md'
  loading='lazy'
/>
```

#### 5.2 アイコンのアクセシビリティ
**影響するファイル:**
- 全コンポーネント（lucide-reactアイコンの使用箇所）

**問題点:**
多くのアイコンにaria-hidden属性が欠如しており、スクリーンリーダーで読み上げられる可能性があります。

**推奨対応:**
```tsx
// テキストラベルと併用する場合
<Utensils className='w-6 h-6' aria-hidden="true" />
<span>メニュー</span>

// アイコンのみで意味を伝える場合
<AlertCircle className='h-4 w-4' aria-label="売り切れ警告" role="img" />
```

---

## 6. カラーコントラストとビジュアルデザイン

### 🟡 重大度：中
#### 6.1 時間表示の色分け
**影響するファイル:**
- `frontend/apps/admin/app/(authenticated)/operation/components/OrderCard.tsx`

**問題点:**
経過時間を色（緑、黄、赤）で表現していますが、色だけに依存しています。

**推奨対応:**
```tsx
// アイコンまたはテキストラベルを追加
<span className={`font-bold ${timeColorClass} flex items-center gap-1`}>
  {elapsedMinutes >= TIME_THRESHOLDS.RED_MINUTES && (
    <AlertTriangle className="h-4 w-4" aria-hidden="true" />
  )}
  {elapsedMinutes}分
  <span className="sr-only">
    {elapsedMinutes >= TIME_THRESHOLDS.RED_MINUTES ? '（急ぎ）' : 
     elapsedMinutes >= TIME_THRESHOLDS.YELLOW_MINUTES ? '（注意）' : 
     '（余裕あり）'}
  </span>
</span>
```

#### 6.2 リンクの視認性
**影響するファイル:**
- `frontend/apps/admin/app/(auth)/login/page.tsx`

**問題点:**
リンクが青色とアンダーラインで表示されていますが、周囲のテキストとのコントラストを確認する必要があります。

**推奨対応:**
自動テストツール（axe-coreなど）でコントラスト比を検証。最小4.5:1を確保。

---

## 7. 言語とローカリゼーション

### ✅ 良好な点
- `<html lang='ja'>`が適切に設定されています（両アプリ）

### 🟡 改善可能な点
#### 7.1 混在する英語テキスト
**影響するファイル:**
- `frontend/apps/admin/components/Sidebar.tsx`
- `frontend/apps/admin/app/(auth)/login/page.tsx`

**問題点:**
UIの一部に英語テキスト（"Welcome back", "Home", "User"など）が残っています。

**推奨対応:**
全てのユーザー向けテキストを日本語に統一、または多言語対応の仕組みを導入。

---

## 8. 動的コンテンツとステート管理

### 🟡 重大度：中
#### 8.1 ローディング状態のアクセシビリティ
**影響するファイル:**
- `frontend/apps/client/app/page.tsx`
- `frontend/apps/admin/components/charts/KPISummaryCards.tsx`

**問題点:**
ローディング中のテキスト（"loading..."）が適切にスクリーンリーダーに通知されません。

**推奨対応:**
```tsx
// Before
if (isLoading) return <div>loading...</div>;

// After
if (isLoading) return (
  <div role="status" aria-live="polite" aria-busy="true">
    <span>読み込み中...</span>
    <span className="sr-only">データを読み込んでいます。しばらくお待ちください。</span>
  </div>
);
```

#### 8.2 エラー状態の通知
**影響するファイル:**
- `frontend/apps/client/app/page.tsx`

**問題点:**
エラーメッセージ（"failed to load"）がaria-liveなしで表示されています。

**推奨対応:**
```tsx
if (error) return (
  <div role="alert" aria-live="assertive">
    <p>データの読み込みに失敗しました。ページを再読み込みしてください。</p>
  </div>
);
```

---

## 9. モバイルアクセシビリティ

### 🟡 重大度：中
#### 9.1 タッチターゲットのサイズ
**影響するファイル:**
- `frontend/apps/client/app/components/TabBar.tsx`
- 全ボタンコンポーネント

**問題点:**
一部のボタンやリンクのタッチターゲットが44x44pxの推奨サイズを下回る可能性があります。

**推奨対応:**
```css
/* 最小タッチターゲットサイズを保証 */
.touch-target {
  min-width: 44px;
  min-height: 44px;
  padding: 8px;
}
```

---

## 10. テスト自動化の推奨事項

### 現状
- アクセシビリティテストツールが導入されていません
- `package.json`に`eslint-plugin-jsx-a11y`、`axe-core`、`@axe-core/react`などのツールが見当たりません

### 推奨対応

#### 10.1 開発時のリンティング
```bash
pnpm add -D eslint-plugin-jsx-a11y
```

`.eslintrc.json`:
```json
{
  "extends": [
    "plugin:jsx-a11y/recommended"
  ],
  "plugins": ["jsx-a11y"]
}
```

#### 10.2 自動テスト
```bash
pnpm add -D @axe-core/react jest-axe @testing-library/react
```

テスト例:
```tsx
import { axe } from 'jest-axe';
import { render } from '@testing-library/react';

test('should not have accessibility violations', async () => {
  const { container } = render(<MyComponent />);
  const results = await axe(container);
  expect(results).toHaveNoViolations();
});
```

#### 10.3 CI/CDパイプラインでの検証
- GitHub Actionsでのaxe-core実行
- Lighthouse CIの導入

---

## 優先度別対応計画

### 🔴 高優先度（即座に対応）
1. **ボタンとリンクのaria-label追加**（2時間）
2. **クリック可能なカードのキーボードアクセス対応**（3時間）
3. **画像の代替テキスト整備**（2時間）
4. **`<main>`要素の追加**（1時間）
5. **チャートの代替テキスト追加**（4時間）

**推定工数: 12時間**

### 🟡 中優先度（1-2週間以内）
1. **スキップリンクの実装**（2時間）
2. **フォームエラー処理改善**（3時間）
3. **ライブリージョンの追加**（2時間）
4. **見出し階層の整理**（3時間）
5. **装飾的アイコンのaria-hidden追加**（2時間）
6. **カラーコントラスト検証と修正**（4時間）

**推定工数: 16時間**

### 🟢 低優先度（継続的改善）
1. **アクセシビリティテストツールの導入**（8時間）
2. **ダークモードのコントラスト検証**（4時間）
3. **多言語対応の検討**（規模による）
4. **包括的なドキュメント作成**（4時間）

**推定工数: 16時間以上**

---

## まとめ

### 強みとして評価できる点
✅ Shadcn UIの使用により、基本的なアクセシビリティが確保されているコンポーネントが多い  
✅ `lang='ja'`属性が適切に設定されている  
✅ フォーカス可視化のスタイルが定義されている  
✅ React Hook Formとの統合により、フォームバリデーションの基盤がある

### 主要な改善領域
❌ ARIA属性の不足（特にボタン、リンク、ライブリージョン）  
❌ キーボードナビゲーション対応の不足（クリック可能なカード）  
❌ 画像とアイコンの代替テキスト不足  
❌ セマンティックHTML要素（`<main>`, `<nav>`）の活用不足  
❌ チャートの代替表現がない  
❌ 自動テストツールの未導入

### 推奨される次のステップ
1. 高優先度項目の実装（12時間）
2. アクセシビリティテストツールの導入（8時間）
3. 中優先度項目の段階的な実装（16時間）
4. 定期的な監査とテストの実施

**総推定工数: 約36-52時間**

---

## 参考リソース

- [WCAG 2.1 Guidelines](https://www.w3.org/WAI/WCAG21/quickref/)
- [Radix UI Accessibility](https://www.radix-ui.com/primitives/docs/overview/accessibility)
- [eslint-plugin-jsx-a11y](https://github.com/jsx-eslint/eslint-plugin-jsx-a11y)
- [axe-core](https://github.com/dequelabs/axe-core)
- [WebAIM Contrast Checker](https://webaim.org/resources/contrastchecker/)

---

**レポート作成者:** GitHub Copilot Agent  
**最終更新日:** 2026年1月13日

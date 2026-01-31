# アクセシビリティクイックリファレンス

開発時に参照するためのアクセシビリティベストプラクティスガイド。

## 🎯 基本原則

アクセシビリティは以下の4つの原則（POUR）に基づきます：

1. **Perceivable（知覚可能）** - 情報とUIコンポーネントは、ユーザーが知覚できる方法で提示されなければならない
2. **Operable（操作可能）** - UIコンポーネントとナビゲーションは操作可能でなければならない
3. **Understandable（理解可能）** - 情報とUIの操作は理解可能でなければならない
4. **Robust（堅牢）** - コンテンツは、支援技術を含む様々なユーザーエージェントによって解釈できるほど堅牢でなければならない

---

## 📝 コンポーネント別チェックリスト

### ボタン

✅ **すべきこと:**
```tsx
// アイコンのみのボタン
<Button aria-label="閉じる">
  <X className="h-4 w-4" aria-hidden="true" />
</Button>

// テキスト付きボタン
<Button>
  <Save className="h-4 w-4 mr-2" aria-hidden="true" />
  保存
</Button>

// ローディング状態
<Button disabled aria-busy="true">
  <Loader className="h-4 w-4 mr-2 animate-spin" aria-hidden="true" />
  処理中...
</Button>
```

❌ **やってはいけないこと:**
```tsx
// アイコンのみで説明なし
<Button>
  <X className="h-4 w-4" />
</Button>

// disabledの説明なし
<Button disabled>送信</Button>
```

### リンク

✅ **すべきこと:**
```tsx
// 明確なリンクテキスト
<Link href="/about">会社概要</Link>

// 外部リンク
<Link href="https://example.com" target="_blank" rel="noopener noreferrer">
  外部サイト
  <ExternalLink className="ml-1 h-4 w-4" aria-hidden="true" />
  <span className="sr-only">（新しいタブで開きます）</span>
</Link>
```

❌ **やってはいけないこと:**
```tsx
// 曖昧なリンクテキスト
<Link href="/details">こちら</Link>
<Link href="/more">もっと見る</Link>
```

### フォーム

✅ **すべきこと:**
```tsx
<FormField
  control={form.control}
  name="email"
  render={({ field }) => (
    <FormItem>
      <FormLabel>
        メールアドレス
        <span className="text-destructive ml-1" aria-hidden="true">*</span>
      </FormLabel>
      <FormControl>
        <Input 
          type="email"
          placeholder="example@example.com"
          aria-required="true"
          aria-invalid={!!errors.email}
          aria-describedby={errors.email ? "email-error" : undefined}
          {...field}
        />
      </FormControl>
      {errors.email && (
        <FormMessage id="email-error" role="alert">
          {errors.email.message}
        </FormMessage>
      )}
      <FormDescription>
        登録時に使用したメールアドレスを入力してください
      </FormDescription>
    </FormItem>
  )}
/>
```

❌ **やってはいけないこと:**
```tsx
// ラベルなし
<Input type="email" placeholder="メールアドレス" />

// プレースホルダーをラベル代わりに使用
<Input placeholder="名前を入力" />
```

### 画像

✅ **すべきこと:**
```tsx
// 意味のある画像
<img 
  src="/product.jpg" 
  alt="白いコットンTシャツ、フロントビュー"
  width={400}
  height={400}
/>

// 装飾的な画像
<img 
  src="/decoration.svg" 
  alt=""
  aria-hidden="true"
  width={100}
  height={100}
/>

// Next.js Image
<Image
  src="/product.jpg"
  alt="白いコットンTシャツ、フロントビュー"
  width={400}
  height={400}
/>
```

❌ **やってはいけないこと:**
```tsx
// alt属性なし
<img src="/product.jpg" />

// 無意味なalt属性
<img src="/product.jpg" alt="画像" />
<img src="/product.jpg" alt="img_001.jpg" />
```

### アイコン

✅ **すべきこと:**
```tsx
// テキストと併用（装飾）
<Button>
  <Save className="h-4 w-4 mr-2" aria-hidden="true" />
  保存
</Button>

// アイコンのみ（意味を持つ）
<button aria-label="メニューを開く">
  <Menu className="h-6 w-6" aria-hidden="true" />
</button>

// 状態を示すアイコン
<span className="flex items-center">
  <CheckCircle className="h-4 w-4 text-green-600 mr-1" aria-hidden="true" />
  <span>完了</span>
  <span className="sr-only">ステータス:</span>
</span>
```

### ナビゲーション

✅ **すべきこと:**
```tsx
// メインナビゲーション
<nav aria-label="メインナビゲーション">
  <ul>
    <li><Link href="/">ホーム</Link></li>
    <li><Link href="/about">会社概要</Link></li>
    <li><Link href="/contact">お問い合わせ</Link></li>
  </ul>
</nav>

// サブナビゲーション
<nav aria-label="カテゴリーナビゲーション">
  {/* カテゴリーリンク */}
</nav>

// パンくずリスト
<nav aria-label="パンくずリスト">
  <ol className="flex gap-2">
    <li><Link href="/">ホーム</Link></li>
    <li aria-hidden="true">/</li>
    <li><Link href="/products">商品一覧</Link></li>
    <li aria-hidden="true">/</li>
    <li aria-current="page">商品詳細</li>
  </ol>
</nav>
```

### モーダル・ダイアログ

✅ **すべきこと:**
```tsx
<Dialog open={isOpen} onOpenChange={setIsOpen}>
  <DialogTrigger asChild>
    <Button>ダイアログを開く</Button>
  </DialogTrigger>
  <DialogContent>
    <DialogHeader>
      <DialogTitle>確認</DialogTitle>
      <DialogDescription>
        この操作は取り消せません。本当に削除しますか？
      </DialogDescription>
    </DialogHeader>
    <DialogFooter>
      <Button variant="outline" onClick={() => setIsOpen(false)}>
        キャンセル
      </Button>
      <Button variant="destructive" onClick={handleDelete}>
        削除
      </Button>
    </DialogFooter>
  </DialogContent>
</Dialog>

// 重要な警告
<AlertDialog>
  <AlertDialogContent>
    <AlertDialogHeader>
      <AlertDialogTitle>警告</AlertDialogTitle>
      <AlertDialogDescription>
        すべてのデータが失われます。
      </AlertDialogDescription>
    </AlertDialogHeader>
    {/* ... */}
  </AlertDialogContent>
</AlertDialog>
```

### テーブル

✅ **すべきこと:**
```tsx
<Table>
  <TableCaption>売上実績（2024年1月）</TableCaption>
  <TableHeader>
    <TableRow>
      <TableHead>日付</TableHead>
      <TableHead className="text-right">売上</TableHead>
      <TableHead className="text-right">注文数</TableHead>
    </TableRow>
  </TableHeader>
  <TableBody>
    <TableRow>
      <TableCell>2024-01-01</TableCell>
      <TableCell className="text-right">¥50,000</TableCell>
      <TableCell className="text-right">25</TableCell>
    </TableRow>
  </TableBody>
</Table>
```

### ローディング状態

✅ **すべきこと:**
```tsx
// ローディング中
{isLoading && (
  <div role="status" aria-live="polite" aria-busy="true">
    <Loader className="h-8 w-8 animate-spin" aria-hidden="true" />
    <span>読み込み中...</span>
    <span className="sr-only">データを読み込んでいます</span>
  </div>
)}

// スケルトンローディング
<div role="status" aria-label="読み込み中">
  <Skeleton className="h-8 w-full" />
  <Skeleton className="h-8 w-full mt-2" />
  <span className="sr-only">コンテンツを読み込んでいます</span>
</div>
```

### エラー状態

✅ **すべきこと:**
```tsx
// エラーメッセージ
{error && (
  <div role="alert" aria-live="assertive" className="text-destructive">
    <AlertCircle className="h-4 w-4 mr-2" aria-hidden="true" />
    <span>エラーが発生しました: {error.message}</span>
  </div>
)}

// フォームエラー
<FormMessage role="alert" aria-live="polite">
  {errors.email?.message}
</FormMessage>
```

### 動的コンテンツ

✅ **すべきこと:**
```tsx
// 重要な更新
<div aria-live="assertive" aria-atomic="true">
  新しい注文が{newOrderCount}件届きました
</div>

// 穏やかな更新
<div aria-live="polite" aria-atomic="true">
  最終更新: {lastUpdated}
</div>

// 時間経過の表示
<span aria-live="polite" aria-atomic="true">
  {elapsedMinutes}分経過
</span>
```

### クリック可能なカード

✅ **すべきこと:**
```tsx
<Card
  tabIndex={0}
  role="button"
  aria-label={`${item.name}の詳細を表示`}
  onClick={handleClick}
  onKeyDown={(e) => {
    if (e.key === 'Enter' || e.key === ' ') {
      e.preventDefault();
      handleClick();
    }
  }}
  className="cursor-pointer focus:outline-none focus:ring-2 focus:ring-ring"
>
  <CardContent>
    {/* コンテンツ */}
  </CardContent>
</Card>
```

❌ **やってはいけないこと:**
```tsx
// キーボードアクセス不可
<Card onClick={handleClick}>
  <CardContent>{/* コンテンツ */}</CardContent>
</Card>
```

---

## 🎨 色とコントラスト

### コントラスト比の基準

- **通常のテキスト**: 最小 4.5:1
- **大きなテキスト（18pt以上または太字14pt以上）**: 最小 3:1
- **UIコンポーネント**: 最小 3:1

### 色だけに依存しない

❌ **悪い例:**
```tsx
<span className="text-red-600">重要</span>
<span className="text-green-600">完了</span>
```

✅ **良い例:**
```tsx
<span className="text-red-600 font-bold">
  <AlertTriangle className="h-4 w-4 mr-1" aria-hidden="true" />
  重要
</span>
<span className="text-green-600">
  <CheckCircle className="h-4 w-4 mr-1" aria-hidden="true" />
  完了
</span>
```

---

## ⌨️ キーボードナビゲーション

### キーボード操作の基本

- **Tab**: 次の要素へフォーカス移動
- **Shift + Tab**: 前の要素へフォーカス移動
- **Enter**: リンクやボタンを実行
- **Space**: ボタンやチェックボックスを実行
- **Escape**: モーダルやドロップダウンを閉じる
- **矢印キー**: ラジオボタングループ、タブ、メニュー内の移動

### フォーカス可能な要素

デフォルトでフォーカス可能:
- `<a href="...">`
- `<button>`
- `<input>`
- `<select>`
- `<textarea>`
- `<summary>`

カスタム要素をフォーカス可能にする:
```tsx
<div
  tabIndex={0}
  role="button"
  onKeyDown={handleKeyDown}
>
  カスタムボタン
</div>
```

### スキップリンク

```tsx
// layout.tsx
<body>
  <a
    href="#main-content"
    className="sr-only focus:not-sr-only focus:absolute focus:top-4 focus:left-4 focus:z-50 focus:bg-primary focus:text-primary-foreground focus:px-4 focus:py-2 focus:rounded"
  >
    メインコンテンツへスキップ
  </a>
  {/* rest of content */}
  <main id="main-content">
    {/* メインコンテンツ */}
  </main>
</body>
```

---

## 📱 モバイルアクセシビリティ

### タッチターゲットサイズ

最小サイズ: **44x44px**

```tsx
// 良い例
<Button className="min-w-[44px] min-h-[44px] p-2">
  <Icon className="h-5 w-5" />
</Button>
```

### ズーム対応

```tsx
// viewport設定（layout.tsx）
export const metadata = {
  viewport: {
    width: 'device-width',
    initialScale: 1,
    maximumScale: 5, // ユーザーのズームを許可
  }
}
```

---

## 🧪 テスト方法

### 手動テスト

1. **キーボードのみでナビゲーション**
   - マウスを使わずにすべての機能にアクセスできるか確認

2. **スクリーンリーダーテスト**
   - macOS: VoiceOver（Command + F5）
   - Windows: NVDA（無料）またはJAWS

3. **ズームテスト**
   - ブラウザで200%にズーム
   - レイアウトが崩れないか確認

4. **カラーコントラスト**
   - ブラウザの開発者ツールで確認
   - WebAIM Contrast Checker使用

### 自動テスト

```tsx
// jest + jest-axe
import { axe, toHaveNoViolations } from 'jest-axe';
import { render } from '@testing-library/react';

expect.extend(toHaveNoViolations);

test('should not have accessibility violations', async () => {
  const { container } = render(<MyComponent />);
  const results = await axe(container);
  expect(results).toHaveNoViolations();
});
```

### ブラウザ拡張機能

- **axe DevTools**: Chrome/Firefox拡張機能
- **WAVE**: WebAIMのアクセシビリティ評価ツール
- **Lighthouse**: Chromeの開発者ツール

---

## 🚫 よくある間違い

### 1. `<div>` や `<span>` をボタンとして使用

❌ **悪い例:**
```tsx
<div onClick={handleClick}>クリック</div>
```

✅ **良い例:**
```tsx
<button onClick={handleClick}>クリック</button>
// または
<Button onClick={handleClick}>クリック</Button>
```

### 2. プレースホルダーをラベル代わりに使用

❌ **悪い例:**
```tsx
<input type="text" placeholder="名前" />
```

✅ **良い例:**
```tsx
<Label htmlFor="name">名前</Label>
<Input id="name" type="text" placeholder="山田太郎" />
```

### 3. 自動再生する動画や音声

❌ **悪い例:**
```tsx
<video autoPlay>
  <source src="video.mp4" />
</video>
```

✅ **良い例:**
```tsx
<video controls>
  <source src="video.mp4" />
  <track kind="captions" src="captions.vtt" srclang="ja" label="日本語" />
</video>
```

### 4. 時間制限のある操作

ユーザーが時間を延長または無効化できるようにする。

### 5. フォーカスの強制移動

ユーザーの意図しないフォーカス移動は避ける。

---

## 📚 参考リソース

### 公式ドキュメント
- [WCAG 2.1](https://www.w3.org/WAI/WCAG21/quickref/)
- [MDN: Accessibility](https://developer.mozilla.org/ja/docs/Web/Accessibility)
- [Radix UI Accessibility](https://www.radix-ui.com/primitives/docs/overview/accessibility)

### ツール
- [axe DevTools](https://www.deque.com/axe/devtools/)
- [WAVE](https://wave.webaim.org/)
- [WebAIM Contrast Checker](https://webaim.org/resources/contrastchecker/)
- [Lighthouse](https://developers.google.com/web/tools/lighthouse)

### ライブラリ
- [eslint-plugin-jsx-a11y](https://github.com/jsx-eslint/eslint-plugin-jsx-a11y)
- [jest-axe](https://github.com/nickcolley/jest-axe)
- [@axe-core/react](https://github.com/dequelabs/axe-core-npm/tree/develop/packages/react)

---

**このガイドは定期的に更新されます。**  
**最終更新日: 2026年1月13日**

"use client";

import { Button } from "@workspace/ui/components/button";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@workspace/ui/components/dialog";
import { Input } from "@workspace/ui/components/input";
import { Plus, RefreshCw, Search, Settings } from "lucide-react";
import { useState } from "react";

interface HeaderToolbarProps {
  onAddCategory: (name: string) => void;
  onRefresh?: () => void;
  searchValue?: string;
  onSearchChange?: (value: string) => void;
}

/**
 * ヘッダーツールバーコンポーネント
 * カテゴリー追加、検索、設定などの機能を提供
 */
export function HeaderToolbar({
  onAddCategory,
  onRefresh,
  searchValue = "",
  onSearchChange,
}: HeaderToolbarProps) {
  const [newCategoryName, setNewCategoryName] = useState("");
  const [isDialogOpen, setIsDialogOpen] = useState(false);

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    const trimmedName = newCategoryName.trim();

    if (!trimmedName) return;

    onAddCategory(trimmedName);
    setNewCategoryName("");
    setIsDialogOpen(false);
  };

  const handleDialogOpenChange = (open: boolean) => {
    setIsDialogOpen(open);
    if (!open) {
      setNewCategoryName("");
    }
  };

  return (
    <div className='flex items-center justify-between p-4 border-b bg-white'>
      <div className='flex items-center gap-4'>
        <h1 className='text-2xl font-bold'>メニュー管理</h1>

        <Dialog open={isDialogOpen} onOpenChange={handleDialogOpenChange}>
          <DialogTrigger asChild>
            <Button className='gap-2'>
              <Plus className='h-4 w-4' />
              カテゴリー追加
            </Button>
          </DialogTrigger>
          <DialogContent className='sm:max-w-md'>
            <DialogHeader>
              <DialogTitle>新しいカテゴリーを追加</DialogTitle>
            </DialogHeader>
            <form onSubmit={handleSubmit} className='space-y-4'>
              <div>
                <label htmlFor='categoryName' className='text-sm font-medium'>
                  カテゴリー名
                </label>
                <Input
                  id='categoryName'
                  placeholder='例: 麺類、ご飯物、一品料理...'
                  value={newCategoryName}
                  onChange={e => setNewCategoryName(e.target.value)}
                  autoFocus
                  required
                  maxLength={50}
                />
              </div>
              <div className='flex justify-end gap-2'>
                <Button
                  type='button'
                  variant='outline'
                  onClick={() => setIsDialogOpen(false)}
                >
                  キャンセル
                </Button>
                <Button type='submit' disabled={!newCategoryName.trim()}>
                  追加
                </Button>
              </div>
            </form>
          </DialogContent>
        </Dialog>
      </div>

      <div className='flex items-center gap-3'>
        {/* 検索機能 */}
        <div className='relative'>
          <Search className='absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground' />
          <Input
            placeholder='メニューを検索...'
            value={searchValue}
            onChange={e => onSearchChange?.(e.target.value)}
            className='pl-10 w-64'
          />
        </div>

        {/* リフレッシュボタン */}
        {onRefresh && (
          <Button
            variant='outline'
            size='sm'
            onClick={onRefresh}
            className='gap-2'
            title='データを再読み込み'
          >
            <RefreshCw className='h-4 w-4' />
          </Button>
        )}

        {/* 設定ボタン（将来の拡張用） */}
        <Button
          variant='ghost'
          size='sm'
          className='gap-2'
          title='設定'
          disabled
        >
          <Settings className='h-4 w-4' />
        </Button>
      </div>
    </div>
  );
}

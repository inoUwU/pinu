"use client";
import { Button } from "@workspace/ui/components/button";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@workspace/ui/components/dialog";
import { useState } from "react";

type Props = {
  table: number;
};

export default function BillDialog({ table }: Props) {
  const [isDialogOpen, setIsDialogOpen] = useState(false);
  const handleKaikei = () => {
    // Handle the "会計する" action here
  };
  return (
    <Dialog open={isDialogOpen} onOpenChange={setIsDialogOpen}>
      <DialogTrigger>会計</DialogTrigger>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>{table}テーブル</DialogTitle>
          <DialogDescription>
            This action cannot be undone. This will permanently delete your
            account and remove your data from our servers.
          </DialogDescription>
        </DialogHeader>
        <DialogFooter>
          <div>
            <Button variant='outline' onClick={() => setIsDialogOpen(false)}>
              キャンセル
            </Button>
          </div>
          <Button
            variant='outline'
            className='bg-red-500 text-white'
            onClick={handleKaikei}
          >
            会計する
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}

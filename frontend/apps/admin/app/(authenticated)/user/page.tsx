"use client";

import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
} from "@workspace/ui/components/dialog";
import { useState } from "react";
import { createColumns } from "@/app/(authenticated)/user/components/column";
import { DataTable } from "@/app/(authenticated)/user/components/data-table";
import UserForm from "@/app/(authenticated)/user/components/form";
import type { UserFromData } from "@/app/(authenticated)/user/types/table";

const initialData: UserFromData[] = [
  {
    id: "1",
    isAdmin: true,
    name: "Admin User",
  },
];

export function UserTable() {
  const [data, setData] = useState<UserFromData[]>(initialData);
  const [editingUser, setEditingUser] = useState<UserFromData | null>(null);
  const [isDialogOpen, setIsDialogOpen] = useState(false);
  const columns = createColumns();

  const handleCreate = (newRecord: Omit<UserFromData, "id">) => {
    const record = { ...newRecord, id: String(data.length + 1) };
    setData([...data, record]);
    setIsDialogOpen(false);
  };

  const handleUpdate = (updatedUser: UserFromData) => {
    setData(
      data.map(record => (record.id === updatedUser.id ? updatedUser : record)),
    );
    setIsDialogOpen(false);
    setEditingUser(null);
  };

  const handleDelete = (id: string) => {
    setData(data.filter(record => record.id !== id));
  };

  const handleEdit = (record: UserFromData) => {
    setEditingUser(record);
    setIsDialogOpen(true);
  };

  const openCreateDialog = () => {
    setEditingUser(null);
    setIsDialogOpen(true);
  };
  return (
    <div>
      <Dialog open={isDialogOpen} onOpenChange={setIsDialogOpen}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>{editingUser ? "Edit" : "Create New"}</DialogTitle>
            <DialogDescription>
              Please fill out the form below to{" "}
              {editingUser ? "update the data" : "create a new data"}.
            </DialogDescription>
          </DialogHeader>
          <div>
            <UserForm
              onSubmit={editingUser ? handleUpdate : handleCreate}
              initialData={editingUser}
            />
          </div>
        </DialogContent>
      </Dialog>
      <DataTable
        columns={columns}
        data={data}
        onAdd={openCreateDialog}
        onEdit={handleEdit}
        onDelete={handleDelete}
      />
    </div>
  );
}

export default function UserPage() {
  return (
    <div>
      <div className='mb-4'>
        <h2>ユーザー管理</h2>
      </div>
      <div>
        <UserTable />
      </div>
    </div>
  );
}

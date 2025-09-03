"use client";

import useSWRSubscription from "swr/subscription";
import TableCard from "./components/TableCard";
export default function BillPage() {
  return (
    <div>
      <TableCard
        table={{ id: 1, name: "テーブル1", price: 1000, quantity: 1 }}
      />
      <div className='mb-4'>
        <h2>会計</h2>
      </div>
      <div className='flex flex-row'>
        <div className='flex-1 bg-amber-600'>
          {/* 左側のテーブル */}
          ひだり
        </div>
      </div>
    </div>
  );
}

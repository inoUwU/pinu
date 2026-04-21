export type TableOrder = {
  id: number;
  tableNumber: number;
  elapsedMinutes: Date;
  dishes: { name: string; qty: number }[];
  status: "pending" | "completed";
  completedAt?: Date;
};

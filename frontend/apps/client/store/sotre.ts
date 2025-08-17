import { create } from "zustand";
import checkoutPage from "@/app/checkout/page";

interface CheckOutState {
  isCheckOuted: boolean;
  setIsCheckOuted: () => void;
}

export const useCheckOutStore = create<CheckOutState>()(set => ({
  isCheckOuted: false,
  setIsCheckOuted: () => set({ isCheckOuted: true }),
}));

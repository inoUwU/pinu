import { create } from "zustand";

interface CheckOutState {
  isCheckOuted: boolean;
  setIsCheckOuted: () => void;
}

export const useCheckOutStore = create<CheckOutState>()(set => ({
  isCheckOuted: false,
  setIsCheckOuted: () => set({ isCheckOuted: true }),
}));

export type LocalIPResponse = {
  ip: string;
  allIPs: Array<IPInfo>;
};

export type IPInfo = {
  interface: string;
  address: string;
  priority: number;
};

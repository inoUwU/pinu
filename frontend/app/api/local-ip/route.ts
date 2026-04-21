import os from "node:os";
import { NextResponse } from "next/server";
import type { IPInfo, LocalIPResponse } from "./types";

export async function GET() {
  const networkInterfaces = os.networkInterfaces();
  let localIP = "localhost";

  // 優先順位でIPアドレスを検索
  const possibleIPs: IPInfo[] = [];

  Object.keys(networkInterfaces).forEach(interfaceName => {
    const networkInterface = networkInterfaces[interfaceName];
    networkInterface?.forEach(network => {
      // IPv4で内部ネットワークでないものを対象
      if (network.family === "IPv4" && !network.internal) {
        // 192.168.x.x, 10.x.x.x, 172.16-31.x.x の範囲を優先
        if (
          network.address.startsWith("192.168.") ||
          network.address.startsWith("10.") ||
          /^172\.(1[6-9]|2[0-9]|3[0-1])\./.test(network.address)
        ) {
          possibleIPs.push({
            interface: interfaceName,
            address: network.address,
            priority: network.address.startsWith("192.168.")
              ? 1
              : network.address.startsWith("10.")
                ? 2
                : 3,
          });
        }
      }
    });
  });

  // 優先順位でソートして最適なIPを選択
  if (possibleIPs.length > 0) {
    possibleIPs.sort((a, b) => a.priority - b.priority);
    const primaryIP = possibleIPs[0];
    if (primaryIP) {
      localIP = primaryIP.address;
    }
  }

  console.log("検出されたIP一覧:", possibleIPs);
  console.log("選択されたIP:", localIP);

  return NextResponse.json({
    ip: localIP,
    allIPs: possibleIPs, // デバッグ用
  } satisfies LocalIPResponse);
}

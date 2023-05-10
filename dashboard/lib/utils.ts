import { clsx, type ClassValue } from "clsx";
import { twMerge } from "tailwind-merge";
import * as timeago from "timeago.js";

export function shortHash(hash: string | undefined, length = 10) {
  if (!hash) {
    return "-";
  }
  return `${hash.slice(0, length)}...`;
}

export function age(ts: number | undefined | null) {
  if (!ts) {
    return "-";
  }
  return timeago.format(ts * 1000);
}

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

export function formatNumber(n: number | string) {
  const formatter = new Intl.NumberFormat("en-US", {
    maximumFractionDigits: 2,
  });
  return formatter.format(Number(n));
}

// format any POST_CREATED style to PostCreated style
export function formatEventType(type: string | null) {
  if (!type) {
    return "-";
  }
  return type
    .split("_")
    .map((w) => w[0] + w.slice(1).toLowerCase())
    .join("");
}

const ipfsGateway = "https://lens.infura-ipfs.io";
export function getIPFSURL(picture: any) {
  let url = "";
  if (!picture) {
    return url;
  }
  if (picture.__typename === "MediaSet") {
    url = picture.original.url;
  }
  if (picture.__typename === "NftImage") {
    url = picture.uri;
  }

  if (url && url.startsWith("ipfs://")) {
    const cid = url.replace("ipfs://", "");
    return `${ipfsGateway}/ipfs/${cid}`;
  }
  return url;
}

export async function getJSONObj(url: string) {
  const res = await fetch(url);
  const json = await res.json();
  console.log("json", json);
  return json;
}

import { clsx, type ClassValue } from "clsx";
import { twMerge } from "tailwind-merge";
import * as timeago from "timeago.js";

export function shortHash(hash: string, length = 10) {
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

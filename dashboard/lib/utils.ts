export function shortHash(hash: string, length = 10) {
  return `${hash.slice(0, length)}...`;
}

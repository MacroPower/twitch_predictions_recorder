interface Timestamped {
  timestamp: Date | string;
}

export function sort(a: Timestamped, b: Timestamped) {
  if (a.timestamp < b.timestamp) {
    return 1;
  }
  if (a.timestamp > b.timestamp) {
    return -1;
  }
  return 0;
}

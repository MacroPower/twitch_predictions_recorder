export function getColor(colorName: string): string {
  if (colorName === "BLUE") {
    return "rgb(56, 122, 255)";
  }
  if (colorName === "PINK") {
    return "rgb(245, 0, 155)";
  }
  return "white";
}

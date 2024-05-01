function reversePrefix(word: string, ch: string): string {
  const idx = word.indexOf(ch);
  if (idx === -1) {
    return word;
  }

  let chars = word.split("");
  let start = 0;
  let end = idx;
  while (start < end) {
    [chars[start], chars[end]] = [chars[end], chars[start]];
    start++;
    end--;
  }

  return chars.join("");
}

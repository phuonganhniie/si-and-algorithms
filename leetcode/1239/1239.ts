function maxLength(arr: string[]): number {
  let bitmasks: number[] = [];

  // convert each string to a bitmask
  for (let string of arr) {
    let bitmask = 0;
    for (let char of string) {
      let charBit = 1 << (char.charCodeAt(0) - "a".charCodeAt(0));
      if (bitmask & charBit) {
        // check for duplicate characters
        bitmask = 0;
        break;
      }
      bitmask |= charBit;
    }
    if (bitmask) {
      bitmasks.push(bitmask);
    }
  }

  console.log(
    "Bitmasks:",
    bitmasks.map((b) => b.toString(2))
  );

  function backtrack(index: number, currentBitMask: number): number {
    if (index === bitmasks.length) {
      return currentBitMask.toString(2).split("0").join("").length; // count of 1's in bitmask
    }

    let maxLength = 0;

    // try including the current string
    if (!(currentBitMask & bitmasks[index])) {
      // check if disjoint
      maxLength = backtrack(index + 1, currentBitMask | bitmasks[index]);
    }

    // try excluding the current string
    maxLength = Math.max(maxLength, backtrack(index + 1, currentBitMask));

    return maxLength;
  }

  return backtrack(0, 0);
}

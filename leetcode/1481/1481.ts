function findLeastNumOfUniqueInts(arr: number[], k: number): number {
  let freqMap = new Map<number, number>();

  for (let num of arr) {
    if (freqMap.has(num)) {
      freqMap.set(num, freqMap.get(num)! + 1);
    } else {
      freqMap.set(num, 1);
    }
  }

  let frequencies = Array.from(freqMap.values());

  frequencies.sort((a, b) => a - b);

  let leastUniqueInts: number = freqMap.size;
  for (let freq of frequencies) {
    if (k >= freq) {
      k -= freq;
      leastUniqueInts--;
    } else {
      break;
    }
  }

  return leastUniqueInts;
}

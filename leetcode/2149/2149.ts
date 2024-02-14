function rearrangeArray(nums: number[]): number[] {
  const n: number = nums.length;
  const rs: number[] = new Array(n).fill(0);
  let posIndex: number = 0,
    negIndex: number = 1;

  for (let num of nums) {
    if (num > 0) {
      rs[posIndex] = num;
      posIndex += 2;
    }

    if (num < 0) {
      rs[negIndex] = num;
      negIndex += 2;
    }
  }

  return rs;
}

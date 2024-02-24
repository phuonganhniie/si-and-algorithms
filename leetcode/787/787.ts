interface FlightInfo {
  nextCity: number;
  cost: number;
}

function findCheapestPrice(
  n: number,
  flights: number[][],
  src: number,
  dst: number,
  k: number
): number {
  const graph: Map<number, FlightInfo[]> = new Map();
  flights.forEach((flight) => {
    const [from, to, cost] = flight;
    if (!graph.has(from)) graph.set(from, []);
    graph.get(from)!.push({ nextCity: to, cost: cost });
  });

  const queue: {
    currentCity: number;
    costSofar: number;
    increasedStop: number;
  }[] = [{ currentCity: 0, costSofar: 0, increasedStop: -1 }];

  const minCostMap: number[] = new Array(n).fill(Number.MAX_SAFE_INTEGER);
  minCostMap[src] = 0;

  while (queue.length > 0) {
    const { currentCity, costSofar, increasedStop } = queue.shift()!;

    if (increasedStop > k) {
      continue;
    }

    graph.get(currentCity)?.forEach(({ nextCity, cost }) => {
      const nextCost = cost + costSofar;
      if (nextCost < minCostMap[nextCity] && increasedStop + 1 <= k) {
        minCostMap[nextCity] = nextCost;
        queue.push({
          currentCity: nextCity,
          costSofar: nextCost,
          increasedStop: increasedStop + 1,
        });
      }
    });
  }

  return minCostMap[dst] === Number.MAX_SAFE_INTEGER ? -1 : minCostMap[dst];
}

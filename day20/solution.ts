type Point = {
    x: number;
    y: number;
};

type Grid = {
    grid: string[][];
    start: Point;
    end: Point;
};

function parseGrid(input: string): Grid {
    const lines = input.trim().split('\n');
    const grid = lines.map(line => line.split(''));
    let start: Point = { x: 0, y: 0 };
    let end: Point = { x: 0, y: 0 };

    // Find start and end positions
    for (let y = 0; y < grid.length; y++) {
        for (let x = 0; x < grid[y].length; x++) {
            if (grid[y][x] === 'S') {
                start = { x, y };
                grid[y][x] = '.'; // Convert S to . for easier path finding
            } else if (grid[y][x] === 'E') {
                end = { x, y };
                grid[y][x] = '.'; // Convert E to . for easier path finding
            }
        }
    }

    return { grid, start, end };
}

function getNeighbors(point: Point, grid: string[][]): Point[] {
    const directions = [
        { x: 0, y: -1 }, // up
        { x: 0, y: 1 },  // down
        { x: -1, y: 0 }, // left
        { x: 1, y: 0 },  // right
    ];

    return directions
        .map(dir => ({
            x: point.x + dir.x,
            y: point.y + dir.y
        }))
        .filter(p => 
            p.y >= 0 && p.y < grid.length &&
            p.x >= 0 && p.x < grid[0].length &&
            grid[p.y][p.x] === '.'
        );
}

function manhattanDistance(a: Point, b: Point): number {
    return Math.abs(a.x - b.x) + Math.abs(a.y - b.y);
}

function findShortestPath(grid: Grid, start: Point, end: Point): number {
    const queue: [Point, number][] = [[start, 0]];
    const visited = new Set<string>();
    const minDist = new Map<string, number>();
    
    while (queue.length > 0) {
        queue.sort((a, b) => {
            const aDist = a[1] + manhattanDistance(a[0], end);
            const bDist = b[1] + manhattanDistance(b[0], end);
            return aDist - bDist;
        });
        
        const [current, steps] = queue.shift()!;
        const key = `${current.x},${current.y}`;
        
        if (current.x === end.x && current.y === end.y) {
            return steps;
        }
        
        if (visited.has(key)) continue;
        visited.add(key);
        
        const currentDist = minDist.get(key) ?? Infinity;
        if (steps >= currentDist) continue;
        minDist.set(key, steps);
        
        for (const neighbor of getNeighbors(current, grid.grid)) {
            queue.push([neighbor, steps + 1]);
        }
    }
    
    return Infinity;
}

type CheatResult = {
    startPoint: Point;
    endPoint: Point;
    timeSaved: number;
};

function findPathWithCheat(grid: Grid, cheatStart: Point, cheatEnd: Point): number {
    // First find shortest path to cheat start
    const toCheatStart = findShortestPath(grid, grid.start, cheatStart);
    if (toCheatStart === Infinity) return Infinity;
    
    // Then find shortest path from cheat end to goal
    const fromCheatEnd = findShortestPath(grid, cheatEnd, grid.end);
    if (fromCheatEnd === Infinity) return Infinity;
    
    // Total time is path to cheat start + 2 (wall traversal) + path from cheat end
    return toCheatStart + 2 + fromCheatEnd;
}

function isValidCheatEndpoint(point: Point, grid: string[][]): boolean {
    return point.y >= 0 && point.y < grid.length &&
           point.x >= 0 && point.x < grid[0].length &&
           grid[point.y][point.x] === '.';
}

function findAllCheats(grid: Grid): CheatResult[] {
    const cheats: CheatResult[] = [];
    const normalTime = findShortestPath(grid, grid.start, grid.end);
    console.log('Normal path length:', normalTime);
    
    // Try all possible cheat start points (must be on normal path)
    for (let y = 0; y < grid.grid.length; y++) {
        for (let x = 0; x < grid.grid[0].length; x++) {
            if (grid.grid[y][x] !== '.') continue;
            
            const startPoint = { x, y };
            // Skip if we can't reach this point normally
            if (findShortestPath(grid, grid.start, startPoint) === Infinity) continue;
            
            // Try all possible end points that are exactly 2 steps away through walls
            for (let dy = -2; dy <= 2; dy++) {
                for (let dx = -2; dx <= 2; dx++) {
                    if (Math.abs(dx) + Math.abs(dy) !== 2) continue; // Must be exactly 2 steps away
                    
                    const endPoint = { x: x + dx, y: y + dy };
                    if (!isValidCheatEndpoint(endPoint, grid.grid)) continue;
                    
                    const timeWithCheat = findPathWithCheat(grid, startPoint, endPoint);
                    const timeSaved = normalTime - timeWithCheat;
                    
                    if (timeSaved > 0) {
                        cheats.push({
                            startPoint,
                            endPoint,
                            timeSaved
                        });
                    }
                }
            }
        }
    }
    
    return cheats;
}

// Read input file
const input = await Deno.readTextFile('./input.txt');
const grid = parseGrid(input);

// Part 1: Find shortest normal path
const shortestPath = findShortestPath(grid, grid.start, grid.end);
console.log('Shortest normal path:', shortestPath);

// Part 2: Find all possible cheats
const cheats = findAllCheats(grid);

// Group cheats by time saved
const cheatsByTime = new Map<number, number>();
for (const cheat of cheats) {
    cheatsByTime.set(cheat.timeSaved, (cheatsByTime.get(cheat.timeSaved) || 0) + 1);
}

console.log('\nCheats by time saved:');
for (const [time, count] of [...cheatsByTime.entries()].sort((a, b) => a[0] - b[0])) {
    console.log(`${count} cheat${count > 1 ? 's' : ''} that save${count === 1 ? 's' : ''} ${time} picoseconds`);
}

const cheatsOver100 = cheats.filter(c => c.timeSaved >= 100);
console.log('\nNumber of cheats that save at least 100 picoseconds:', cheatsOver100.length);

// Debug output for cheats that save a lot of time
const topCheats = [...cheats].sort((a, b) => b.timeSaved - a.timeSaved).slice(0, 5);
console.log('\nTop 5 time-saving cheats:');
for (const cheat of topCheats) {
    console.log(`Saves ${cheat.timeSaved} picoseconds:`,
        `Start(${cheat.startPoint.x},${cheat.startPoint.y})`,
        `-> End(${cheat.endPoint.x},${cheat.endPoint.y})`);
}

function tcgen() {
    const nodes = Math.ceil(Math.random() * 10)
    const edges = Math.ceil(Math.random() * 100)

    if (edges > (nodes * (nodes - 1) / 2)) {
        throw ""
    }

    let output = ''

    const startPoint = Math.ceil(Math.random() * nodes)

    output += `${nodes} ${edges} ${startPoint}\n`
    output += generateEdge(nodes, edges) + "\n"

    return output
}

function generateEdge(nodes, edges_count) {
    const edges = []
    // const p1 Math.random() * nodes
    while (edges.length !== edges_count) {
        const p1 = Math.ceil(Math.random() * nodes)
        const p2 = Math.ceil(Math.random() * nodes)

        if (edges.includes(`${p1} ${p2}`) || edges.includes(`${p2} ${p1}`)) { continue }

        edges.push(`${p1} ${p2}`)
    }

    return edges.join("\n")
}

let tc

while (true) {
    try {
        tc = tcgen()
        break
    } catch { }
}

console.log(tc)
console.log("---------")

class Graph {
    constructor() {
        this.nodeDict = {};
        this.visitedNodes = new Set();
    }

    appendEdge(p1, p2) {
        if (!this.nodeDict[p1]) {
            this.nodeDict[p1] = [];
        }
        if (!this.nodeDict[p2]) {
            this.nodeDict[p2] = [];
        }

        this.nodeDict[p1].push(p2);
        this.nodeDict[p2].push(p1);
    }

    dfs(start) {
        const stack = [start];
        this.visitedNodes = new Set();

        while (stack.length > 0) {
            const node = stack.pop();
            if (this.visitedNodes.has(node)) {
                continue;
            }

            this.visitedNodes.add(node);
            process.stdout.write(node + ' ');

            const nextNodes = (this.nodeDict[node] || []).sort((a, b) => b - a);
            for (const nextNode of nextNodes) {
                stack.push(nextNode);
            }
        }
        console.log();
    }

    bfs(start) {
        const queue = [start];
        this.visitedNodes = new Set();

        while (queue.length > 0) {
            const node = queue.shift();
            if (this.visitedNodes.has(node)) {
                continue;
            }

            this.visitedNodes.add(node);
            process.stdout.write(node + ' ');

            const nextNodes = (this.nodeDict[node] || []).sort((a, b) => a - b);
            for (const nextNode of nextNodes) {
                queue.push(nextNode);
            }
        }
        console.log();
    }
}

// 입력 처리 및 실행 부분
const fs = require('fs');
const input = tc.split("\n")
const [nodes, edges, startPoint] = input[0].split(' ').map(Number);
const graph = new Graph();

for (let i = 1; i <= edges; i++) {
    const [p1, p2] = input[i].split(' ').map(Number);
    graph.appendEdge(p1, p2);
}

graph.dfs(startPoint);
graph.bfs(startPoint);


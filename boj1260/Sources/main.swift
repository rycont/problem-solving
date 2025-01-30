class Node {
    var id: Int
    var connections: [Int]

    init(id: Int, connections: [Int] = []) {
        self.id = id
        self.connections = connections
    }
}

let firstLine = readLine()!.split(separator: " ").map { Int($0)! }

let nodes = firstLine[0]
let edges = firstLine[1]
let startPoint = firstLine[2]

let graph = Graph()

for _ in 0..<edges {
    let edgeInput = readLine()!.split(separator: " ").map { Int($0)! }
    let p1 = edgeInput[0]
    let p2 = edgeInput[1]

    graph.appendEdge(p1, p2)
}

graph.dfs(start: startPoint)
print()
graph.bfs(start: startPoint)
print()

class Graph {
    var visitedNodes = Set<Int>()
    var nodeDict = [Int: Node]()

    func appendEdge(_ p1: Int, _ p2: Int) {
        if nodeDict[p1] == nil {
            nodeDict[p1] = Node(
                id: p1, connections: []
            )
        }

        if nodeDict[p2] == nil {
            nodeDict[p2] = Node(
                id: p2, connections: []
            )
        }

        nodeDict[p1]!.connections.append(p2)
        nodeDict[p2]!.connections.append(p1)
    }

    func dfs(start: Int) {
        var stack = [Int]()
        stack.append(start)

        while !stack.isEmpty {
            let nextNodeId = stack.removeLast()

            if !nodeDict.keys.contains(nextNodeId) {
                visitedNodes.insert(nextNodeId)
                print(nextNodeId, terminator: " ")
                continue
            }

            let node = nodeDict[nextNodeId]!

            if visitedNodes.contains(node.id) {
                continue
            }

            visitedNodes.insert(node.id)
            print(node.id, terminator: " ")

            for next in node.connections.sorted(by: >) {
                stack.append(next)
            }
        }
    }

    func bfs(start: Int) {
        clearVisit()

        var queue = [Int]()
        queue.append(startPoint)

        while !queue.isEmpty {
            let nextNodeId = queue.removeFirst()

            if !nodeDict.keys.contains(nextNodeId) {
                visitedNodes.insert(nextNodeId)
                print(nextNodeId, terminator: " ")
                continue
            }

            let node = nodeDict[nextNodeId]!

            if visitedNodes.contains(nextNodeId) {
                continue
            }

            visitedNodes.insert(nextNodeId)
            print(node.id, terminator: " ")

            for next in node.connections.sorted() {
                queue.append(next)
            }
        }
    }

    func clearVisit() {
        visitedNodes.removeAll()
    }
}

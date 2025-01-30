class Node:
    def __init__(self, id, connections=None):
        if connections is None:
            connections = []
        self.id = id
        self.connections = connections


class Graph:
    def __init__(self):
        self.visited_nodes = set()
        self.node_dict = {}

    def append_edge(self, p1, p2):
        if p1 not in self.node_dict:
            self.node_dict[p1] = Node(p1)
        if p2 not in self.node_dict:
            self.node_dict[p2] = Node(p2)

        self.node_dict[p1].connections.append(p2)
        self.node_dict[p2].connections.append(p1)

    def dfs(self, start):
        stack = [start]

        while stack:
            next_node_id = stack.pop()

            if next_node_id not in self.node_dict:
                print(next_node_id, end=" ")
                continue

            node = self.node_dict[next_node_id]

            if node.id in self.visited_nodes:
                continue

            self.visited_nodes.add(node.id)
            print(node.id, end=" ")

            for next_node in sorted(node.connections, reverse=True):
                stack.append(next_node)

    def bfs(self, start):
        self.clear_visit()

        queue = [start]

        while queue:
            next_node_id = queue.pop(0)

            if next_node_id not in self.node_dict:
                print(next_node_id, end=" ")
                continue

            node = self.node_dict[next_node_id]

            if node.id in self.visited_nodes:
                continue

            self.visited_nodes.add(node.id)
            print(node.id, end=" ")

            for next_node in sorted(node.connections):
                queue.append(next_node)

    def clear_visit(self):
        self.visited_nodes.clear()


# 입력 처리
first_line = list(map(int, input().split()))
nodes = first_line[0]
edges = first_line[1]
start_point = first_line[2]

graph = Graph()

for _ in range(edges):
    edge_input = list(map(int, input().split()))
    p1 = edge_input[0]
    p2 = edge_input[1]

    graph.append_edge(p1, p2)

graph.dfs(start=start_point)
print()
graph.bfs(start=start_point)
print()

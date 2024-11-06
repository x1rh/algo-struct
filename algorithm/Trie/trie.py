class Node:
    def __init__(self, value=None):
        self.value = value
        self.children = dict()


class Trie:
    def __init__(self):
        self.root = None

    def empty(self):
        return not self.root

    def insert(self, key):
        if self.empty():
            self.root = Node()
        node = self.root
        for ch in key:
            if ch not in node.children:
                node.children[ch] = Node()
            node = node.children[ch]
        node.value = key

    def find(self, key):
        if self.empty():
            return False
        node = self.root
        for ch in key:
            if ch not in node.children:
                return False
            else:
                node = node.children[ch]
        return True

    def dfs(self, root, ans):
        if self.empty():
            return None
        if root.value:
            ans.append(root.value)
        for k, v in root.children.items():
            self.dfs(v, ans)

    def match_prefix(self, prefix):
        if self.empty() or not self.find(prefix):
            return None

        res = list()
        node = self.root
        for ch in prefix:
            node = node.children[ch]
        self.dfs(node, res)
        return res


if __name__ == '__main__':
    trie = Trie()
    trie.insert('dog')
    trie.insert('cat')
    trie.insert('california')

    print(trie.find('adog'))
    print(trie.find('acat'))
    print(trie.find('dog'))
    print(trie.find('california'))

    print('prefix c:', trie.match_prefix('c'))
    print('prefix ca:', trie.match_prefix('ca'))

    content = list()
    trie.dfs(trie.root, content)
    print('all content in trie tree:', content)
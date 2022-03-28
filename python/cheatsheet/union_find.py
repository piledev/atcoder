class UnionFind:
    # UnionFind is a data structure good at group management.
    # Merging groups and checking group to which x belongs are especially fast.
    # n: Number of peaple to divide into groups.
    def __init__(self,n):
        self.n=n
        self.parent_size=[-1]*n
    # Records a and b as members of the same group.
    def merge(self, a, b):
        x, y=self.root(a),self.root(b)
        if x==y: return
        if abs(self.parent_size[x]) < abs(self.parent_size[y]): x, y = y, x
        self.parent_size[x] += self.parent_size[y]
        self.parent_size[y] = x
        return
    # See if a and b belongs to the same group.
    # Returns True for the same group.
    # Returns False for the different groups.
    def same(self, a, b):
        return self.root(a) == self.root(b)
    # Returns the root of the group to which `a` belongs.
    def root(self, a):
        if self.parent_size[a]<0: return a
        self.parent_size[a]=self.root(self.parent_size[a])
        return self.parent_size[a]
    # Returns the number of the group to which `a` belongs.
    def size(self, a):
        return abs(self.parent_size[self.root(a)])
    # Returns all groups in the form of a two-demensinal array.
    def groups(self):
        result=[[] for _ in range(self.n)]
        for i in range(self.n):
            result[self.root(i)].append(i)
        return [r for r in result if r != []]


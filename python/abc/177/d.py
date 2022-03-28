class UnionFind:
    # UnionFind is 
    # n: Number of peaple to divide into groups.
    def __init__(self,n):
        self.n=n
        self.parent_size=[-1]*n
    # Records a and b as members of the same group.
    def merge(self, a, b):
        x, y=self.leader(a), self.leader(b)
        if x==y: return
        if abs(self.parent_size[x]) < abs(self.parent_size[y]): x, y = y, x
        self.parent_size[x] += self.parent_size[y]
        self.parent_size[y] = x
        return
    # See if a and b belongs to the same group.
    # Returns True for the same group.
    # Returns False for the different groups.
    def same(self, a, b):
        return self.leader(a) == self.leader(b)
    # Returns the leader of the group to which `a` belongs.
    def leader(self, a):
        if self.parent_size[a]<0: return a
        self.parent_size[a]=self.leader(self.parent_size[a])
        return self.parent_size[a]
    # Returns the number of the group to which `a` belongs.
    def size(self, a):
        return abs(self.parent_size[self.leader(a)])
    # Returns all groups in the form of a two-demensinal array.
    def groups(self):
        result=[[] for _ in range(self.n)]
        for i in range(self.n):
            result[self.leader(i)].append(i)
        return [r for r in result if r != []]

n,m=map(int,input().split())
uf = UnionFind(n)
for i in range(m):
    a,b=map(int,input().split())
    a-=1
    b-=1
    uf.merge(a, b)

friends_groups=uf.groups()
friends_sizes=[]
for group in friends_groups:
    friends_sizes.append(len(group))
print(max(friends_sizes))


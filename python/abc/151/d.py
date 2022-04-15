h,w=map(int,input().split())
maze=list()
for i in range(h):
    maze.append(list(input()))

from collections import deque
# bfs() returns max value of min path: 
def bfs(start_row,start_col):
    if maze[start_row][start_col]=="#":
        return 0

    maze_count=[[-1]*w for i in range(h)]
    maze_count[start_row][start_col]=0
    que=deque()
    que.append((start_row,start_col))
    while len(que)>0:
       cur_row,cur_col=que.popleft() 
       # up 
       if 0<=cur_row-1<h and 0<=cur_col<w:
           if maze[cur_row-1][cur_col]=='.':
               if maze_count[cur_row-1][cur_col]==-1:
                   maze_count[cur_row-1][cur_col]=maze_count[cur_row][cur_col]+1
                   que.append((cur_row-1,cur_col))
       # down
       if 0<=cur_row+1<h and 0<=cur_col<w:
           if maze[cur_row+1][cur_col]=='.':
               if maze_count[cur_row+1][cur_col]==-1:
                   maze_count[cur_row+1][cur_col]=maze_count[cur_row][cur_col]+1
                   que.append((cur_row+1,cur_col))
       # left 
       if 0<=cur_row<h and 0<=cur_col-1<w:
           if maze[cur_row][cur_col-1]=='.':
               if maze_count[cur_row][cur_col-1]==-1:
                   maze_count[cur_row][cur_col-1]=maze_count[cur_row][cur_col]+1
                   que.append((cur_row,cur_col-1))
       # right
       if 0<=cur_row<h and 0<=cur_col+1<w:
           if maze[cur_row][cur_col+1]=='.':
               if maze_count[cur_row][cur_col+1]==-1:
                   maze_count[cur_row][cur_col+1]=maze_count[cur_row][cur_col]+1
                   que.append((cur_row,cur_col+1))

    ret=0
    for i in range(h):
        for j in range(w):
            ret=max(ret,maze_count[i][j])
    return ret 
    
ans=0
for i in range(h):
    for j in range(w):
        ans=max(ans,bfs(i,j))

print(ans)

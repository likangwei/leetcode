# 洪强宁公司《爱因互动》 出的题目

# 题目: N*M的网格，将水倒入最上角网格，问一共流过了几个网格

grid = [[3, 5, 1], [2, 1, 5], [4,2, 1]]
n, m = len(grid), len(grid[0])
lst = [(0,0)]
rst = set()
while lst:
    element = lst.pop()
    rst.add(element)
    i, j = element
    left, right, top, bottom = (i, j-1), (i, j+1), (i-1, j), (i+1, j)
    for t in [left, right, top, bottom]:
        if t not in rst and 0 <= t[0] <= (n-1) and 0 <= t[1] <= (m-1):
            if grid[i][j] >= grid[t[0]][t[1]]:
                lst.append(t)
print len(rst)

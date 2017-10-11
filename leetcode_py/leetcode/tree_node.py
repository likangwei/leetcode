__author__ = 'lxzMac'

class TreeNode:
    def __init__(self, x):
        self.val = x
        self.par = None
        self.left = None
        self.right = None
        self.next = None

class TreeLinkNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None
        self.next = None

def getTreeNodeRoot(str):

    test = """1,2,2,#,3,3"""
    test = str[1:-1]
    strs = test.split(",")

    root = TreeNode(None)
    list = []
    list.append(root)

    for i in range(len(strs)):
        str = strs[i].strip().replace("{", "").replace("}", "")
        if str == "#":
            strs[i] = None

    cur_level_roots = [root]
    upstair_roots = []
    cur_str_index = 0


    while cur_level_roots:

        cur_level_roots_length = len(cur_level_roots)

        for idx in range(cur_level_roots_length):
            cur_root = cur_level_roots[idx]
            upstair_root_index = idx/2
            isRight = idx%2
            if cur_str_index < len(strs):
                cur_str = strs[cur_str_index]
            else:
                cur_str = None
            cur_str_index = cur_str_index + 1
            if cur_str == None:
                if isRight:
                    upstair_roots[upstair_root_index].right = None
                else:
                    upstair_roots[upstair_root_index].left = None
            else:
                cur_root.val = cur_str
                cur_root.left = TreeNode(None)
                cur_root.right = TreeNode(None)
                cur_level_roots.append(cur_root.left)
                cur_level_roots.append(cur_root.right)

        upstair_roots = cur_level_roots[0:cur_level_roots_length]
        for up_root in upstair_roots:
            if up_root.val == None:
                upstair_roots.remove(up_root)
        del cur_level_roots[0:cur_level_roots_length]


    return root
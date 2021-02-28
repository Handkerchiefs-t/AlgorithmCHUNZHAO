## h学习笔记

- Trie树

  - Trie可以方便地对字符串进行前缀匹配

  - Trie需要实现的主要方法有三个：

    - Insert：向Trie树中添加一个单词
    - Search：搜索一个单词是否存在于Trie树中
    - StartsWith：判断Trie中是否有指定前缀的单词

  - 力扣208题，要我们的实现Trie树：

    ```go
    type Trie struct {
        isEnd bool
        nextList [26]*Trie
    }
    
    
    /** Initialize your data structure here. */
    func Constructor() Trie {
        return Trie{}
    }
    
    
    /** Inserts a word into the trie. */
    func (this *Trie) Insert(word string)  {
        root := this
    
        for _, char := range word {
            if root.nextList[char - 'a'] == nil {
                root.nextList[char - 'a'] = &Trie{}
            }
    
            root = root.nextList[char - 'a']
        }
    
        root.isEnd = true
    }
    
    
    /** Returns if the word is in the trie. */
    func (this *Trie) Search(word string) bool {
        root := this
    
        for _, char := range word {
            if root.nextList[char - 'a'] == nil {
                return false
            }
    
            root = root.nextList[char - 'a']
        }
    
        return root.isEnd
    }
    
    
    /** Returns if there is any word in the trie that starts with the given prefix. */
    func (this *Trie) StartsWith(prefix string) bool {
        root := this
    
        for _, char := range prefix {
            if root.nextList[char - 'a'] == nil {
                return false
            }
    
            root = root.nextList[char - 'a']
        }
    
        return true
    }
    
    
    /**
     * Your Trie object will be instantiated and called as such:
     * obj := Constructor();
     * obj.Insert(word);
     * param_2 := obj.Search(word);
     * param_3 := obj.StartsWith(prefix);
     */
    ```

  - Trie树可以加快字符串的查找、也可以用于实现搜索提示功能

- 并查集

  - 并查集的使用方式很像集合，但是似乎性能不如直接使用语言提供的集合

  - 和集合不同的是，并查集中一个集合会有一个元素作为这个集合的代表

  - 并查集有三个方法：

    - init：初始化
    - find：找到一个元素所属的集合代表
    - union：将两个集合合并

  - 借助力扣547，可以看到并查集的典型实现：

    ```go
    func findCircleNum(isConnected [][]int) int {
        n := len(isConnected)
    
        // 初始化
        parent := make([]int, n)
        // parent[i] = i 是并查集最常见的一种初始化方法
        for i := range parent {
            parent[i] = i
        }
    
        // 寻找一个元素的代表
        // 是代表的条件为：i == parent[i]
        _find := func (i int) int {
            root := i
    
            for i != parent[i] {
                i = parent[i]
            }
    
            // 这一步是将路径压缩
            // 即将一个集合里的所有元素的 parent 都指向代表
            for root != parent[root] {
                tmp := root
                root = parent[root]
                parent[tmp] = i
            }
    
            return i
        }
    
        // 将两个元素所在的集合合并
        // 即：将两个集合的代表合并成一个
        _union := func (i, j int) {
            p1 := _find(i)
            p2 := _find(j)
            
            parent[p1] = p2
        }
    
        for i := 0; i < n; i++ {
            for j := 0; j < n; j++ {
                if isConnected[i][j] == 1 {
                    _union(i, j)
                }
            }
        }
    
        // fmt.Println(parent)
    
        count := 0
        for i, v := range parent {
            if i == v {
                count++
            }
        }
    
        return count
    }
    ```

- 高级搜索

  - 剪枝：在解决问题的时候，我们可以使用剪枝来减少计算
    - 如果你要用决策树来解决问题，在决策过程中会出现重复问题（如计算斐波那契数列），使用记忆化的方法剪枝可以避免重复计算
    - 同样是决策树，到达某个状态可能有很多种路径来实现。不同的路径付出的代价是不同的，使用最优解代替次优解来剪枝，可以提高搜索效率
    - 实际上，上面两种方式的剪枝和动态规划有异曲同工之妙
    - 如果你要暴力回溯求解某一类问题，这类问题必须要达成一定的条件才行，那你可以在每次前进的时候检查是否达成这个条件，如果失败，直接放弃这之后的尝试
  - 双向BFS：单向的广度优先遍历需要解决的问题可能会越来越大，这导致新遍历一层要会费很多的时间。如果使用双向BFS，则可以更快地探索下一层
  - 启发式搜索
    - 暴力搜索很耗时，它会枚举全部解
    - 如果我们可以判断某个点（某个状态）更接近终点，优先探索这个点，可能会有更好的结果
    - 使用这种“判断更接近终点”的方式搜索，就是启发式搜索。要实现启发式搜索，需要使用优先队列并设计一个估价函数
    - 估价函数：很多问题很难找到一个最好的估价函数，所以个这个估价函数是启发式搜索的灵魂
    - 优先队列：保证由估价函数计算出的“更优解”优先搜索

- AVL和红黑树

  - 二叉搜索树：左子树的所有节点都比当前节点的值要小，右边的要大
  - 二叉搜索树的问题：会存在高度差很大的情况，极端情况下会退化成一个链表。为了解决这种问题，有了平衡二叉搜索树
  - AVL 和 红黑树 是比较常用的平衡二叉搜索树
  - AVL：
    - 平衡因子：子树的高度差在因子规定的范围内
    - 旋转操作：左旋、右旋、左右旋、右左旋
    - 不足：要额外信息、由于严格平衡会导致频繁变更
  - 红黑树：
    - 近似平衡二叉树：任何一个节点左右子树的高度差小于两倍
    - 根节点是黑色、叶子节点也是黑色
    - 不能有两个相邻的红色节点
    - 任何一个节点到叶子节点有相同的黑色节点
  - 两者比较：
    - AVL的查询更快
    - 红黑更快添加和删除
    - AVL要存额外信息，红黑只要一个bit表示红和黑
    - 场景：多查询用AVL，其他情况用红黑树比较多

- 位运算

  - 判断奇偶：
    - x & 1 == 0：偶数
    - x & 1 == 1：奇数
  - 除以2：x = x >> 1
  - 将最低位的1置零：x = x & (x-1)
  - 得到最低位的1的幂值：x & (-x) ， 注意 -x 是负数，不是取反
  - 获取第n位的值：(x >> n) & 1
  - 获取第n位的幂值：x & (1 << n)
  - 将指定位置设置为1：x | (1 << n)
  - 将指定位置设置为0：x & (~(1 << n))
  - 将最高位到n清零：x & ((1<<n)-1)
  - 最右边的n位清零：x & (~0 << n)






















































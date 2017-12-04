
package main
import "fmt"
/*

https://leetcode.com/problems/add-and-search-word-data-structure-design/description/

description:

专家答案:
type WordDictionary struct {
    root *TrieNode
}

type TrieNode struct {
    children [26]*TrieNode
    isWord bool
}

func Constructor() WordDictionary {
    var dict WordDictionary
    dict.root = new(TrieNode)
    return dict
}


func (this *WordDictionary) AddWord(word string)  {
    node := this.root
    for i := 0; i < len(word); i++ {
        ci := int(word[i] - 'a')
        if node.children[ci] == nil {
            node.children[ci] = new(TrieNode)
        }
        
        node = node.children[ci]
    }
    
    node.isWord = true
}

func doSearch(node *TrieNode, word string) bool {
    if len(word) == 0 {
        return node.isWord
    }
    
    for i := 0; i < len(word); i++ {
        c := word[i]
        if c == '.' {
            for j := 0; j < 26; j++ {
                if node.children[j] == nil {continue}
                if doSearch(node.children[j], word[i + 1:]) {return true}
            }
            
            node = nil
        } else {
            node = node.children[int(c - 'a')]
        }
        
        if node == nil {return false}
    }
    
    return node.isWord
}

func (this *WordDictionary) Search(word string) bool {
    return doSearch(this.root, word)
}


# 速度类

Q: 此次的时间消耗是？ vs 专家？
A: 176ms, 专家是 152ms,但是我提交专家答案是209ms

Q: 此题目的算法复杂度是？
A:

# 可读性
Q: 此次的算法精炼度是怎么样的？
A:

Q: 此次的命名、可读性怎么样？
A:

# 思路类
Q: 此次是否有你递归，专家没递归的情况？
A: 

Q: 为什么你没有想到专家思路？
A:

# 后续改进
Q: 此次做的差的地方?
A: 

Q: 此次发挥好的地方?
A:


Q: 之前的总结是否生效？这次没用上的原因是？以后如何保证能用上？
A:

Q: 后续如何改进?
A:

Q: 此问题是否可以有实际应用场景？
A:

Q: 此题目对应的TAG是trie, design, backtracking, 此TAG解决了哪类问题？
A: 


3. TODO:



*/

//Expert answer

type Node struct{
	C byte
	Valid bool
	Childs []*Node
}

func (this *Node) Str() string{
	childs := []string{}
	for _, c := range this.Childs{
		childs = append(childs, c.Str())
	}
	return fmt.Sprintf("{%c_%v: %v}", this.C, this.Valid, childs)
}

type WordDictionary struct {
	head *Node
}


/** Initialize your data structure here. */
func Constructor() WordDictionary {
	var lst []*Node
	head := Node{byte(' '), false, lst}
	return WordDictionary{head: &head}
}


/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string)  {
	cur := this.head
	for i:=0; i<len(word); i++{
		var f *Node
		for _, child := range cur.Childs{
			if child.C == word[i]{
				f = child
				break
			}
		}
		if f == nil{
			f = &Node{word[i], false, []*Node{}}
			cur.Childs = append(cur.Childs, f)
		}
		cur = f
	}
	cur.Valid = true
}


/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	var curStair = []*Node{this.head}
	for i:=0; i<len(word); i++{
		NextStair := []*Node{}
		b := word[i]
		for _, node := range curStair{
			if b == byte('.'){
				NextStair = append(NextStair, node.Childs...)
			}else{
				for _, child := range node.Childs{
					if child.C == b{
						NextStair = append(NextStair, child)
						break
					}
				}
			}
		}
		curStair = NextStair
	}
	for _, node := range curStair{
		if node.Valid{
			return true
		}
	}
	return false
}

func (this *WordDictionary) Str()string{
	return this.head.Str()
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */

func main() {
	to_test := []string{"a", "b", "abcdefg"}
	wdic := Constructor()
	for i:=0; i < len(to_test); i++{
		wdic.AddWord(to_test[i])
	}
	fmt.Println(wdic.Str())

	to_test = []string{"a", "c", "abcef", "abcdefg",}
	for i:=0; i < len(to_test); i++{
		fmt.Println(to_test[i], wdic.Search(to_test[i]))
	}
	
}
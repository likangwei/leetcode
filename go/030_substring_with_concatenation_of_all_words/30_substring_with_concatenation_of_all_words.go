/*
https://leetcode.com/problems/implement-strstr/description/
高手答案 xms < xms

*/

package main
import "fmt"
/*
https://leetcode.com/problems/substring-with-concatenation-of-all-words/description/

我的结果：
  1. 第一次耗时非常长， 用内存也多
  2. 第二次优化完，accept



高手答案
func skip(s string, start, k int, word string, occurence map[string]int) int {
	var i int
	for i = start; i <= len(s)-k; i += k {
		w := s[i : i+k]
		occurence[w]--
		if w == word {
			break
		}
	}
	return i + k
}

func clearOccurence(occurence map[string]int) {
	for word, _ := range occurence {
		occurence[word] = 0
	}
}

func find(s string, offset, k, total int, tbl map[string]int) []int {
	start := -1
	res := make([]int, 0)
	occurence := make(map[string]int)
	for i := 0; i <= len(s)-k; i += k {
		word := s[i : i+k]
		if _, ok := tbl[word]; !ok {
			start = -1
			continue
		}
		if start == -1 {
			start = i
			clearOccurence(occurence)
		}
		occurence[word] += 1
		if occurence[word] > tbl[word] {
			start = skip(s, start, k, word, occurence)
		}
		if i+k-start == total {
			res = append(res, start+offset)
		}
	}
	return res
}

func findSubstring(s string, words []string) []int {
	if len(s) == 0 || len(words) == 0 {
		return []int{}
	}
	k := len(words[0])
	if len(s) < k {
		return []int{}
	}
	tbl := make(map[string]int)
	for _, word := range words {
		tbl[word] += 1
	}
	total := len(words) * k
	ans := make([]int, 0)
	for i := 0; i < k; i++ {
		res := find(s[i:], i, k, total, tbl)
		ans = append(ans, res...)
	}
	return ans
}
当前水平改进：
  1. 为什么第一次耗时长，内存多？ 因为我喜欢遍历所有情况，将所有可能出现的情景都复现，有好多次这种操作了

高手答案没看懂。。
*/

type Map struct {
	Childs *map[string]*Map
}


func buildCombinationMap(words []string) *Map{
	if len(words) == 0{
		return nil
	}
	childs := make(map[string]*Map)
	rst := Map{&childs}
	for idx, s := range words{
		left_words := []string{}
		left_words = append(left_words, words[0:idx]...)
		left_words = append(left_words, words[idx+1:len(words)]...)
		childs[s] = buildCombinationMap(left_words)
	}
	return &rst
}

func getStr(m *Map) string {
	if m == nil{
		return "nil"
	}
	rst := make(map[string]string)
	for k, v := range *m.Childs{
		rst[k] = getStr(v)
	}
	
	return fmt.Sprint(rst)
}

func findSubstring(s string, words []string) []int {
	rst := []int{}
	if len(words) == 0 || len(words[0]) == 0{
		return []int{0}
	}
	if len(s)==0{
		return []int{}
	}
	len_total := len(words[0]) * len(words)
	skip := len(words[0])
	if len_total > len(s){
		return []int{}
	}
	m := buildCombinationMap(words)
	fmt.Println(getStr(m))
	for i:=0; i <= len(s) - len_total; i++{
		curMap := m
		found := true
		for j:=i; j<i+len_total; j+=skip{
			curStr := s[j:j+skip]
			childs := *curMap.Childs
			v, ok := childs[curStr]
			if ok{
				curMap = v
			}else{
				found = false
				break
			}
		}
		if found{
			rst = append(rst, i)
		}
	}
	return rst
}


func findSubstring2(s string, words []string) []int {
	rst := []int{}
	if len(words) == 0 || len(words[0]) == 0{
		return []int{0}
	}
	if len(s)==0{
		return []int{}
	}
	len_total := len(words[0]) * len(words)
	skip := len(words[0])
	if len_total > len(s){
		return []int{}
	}
	m := make(map[string]int)
	for _, s := range words{
		v, ok := m[s]
		if ok{
			m[s] = v + 1
		}else{
			m[s] = 1
		}
	}
	for i:=0; i <= len(s) - len_total; i++{
		curMap := make(map[string]int)
		found := true
		for j:=i; j<i+len_total; j+=skip{
			curStr := s[j:j+skip]
			v, ok := m[curStr]
			if ok{
				c, ok := curMap[curStr]
				if ok{
					c = c + 1
				}else{
					c = 1
				}
				curMap[curStr] = c
				if c > v{
					found = false
					break
				}

			}else{
				found = false
				break
			}
		}
		if found{
			rst = append(rst, i)
		}
	}
	return rst
}

func skip(s string, start, k int, word string, occurence map[string]int) int {
	fmt.Println("skip", s, start, k, word, occurence)
	var i int
	for i = start; i <= len(s)-k; i += k {
		w := s[i : i+k]
		occurence[w]--
		if w == word {
			break
		}
	}
	return i + k
}

func clearOccurence(occurence map[string]int) {
	for word, _ := range occurence {
		occurence[word] = 0
	}
}

func find(s string, offset, k, total int, tbl map[string]int) []int {
	fmt.Println(s, offset, k, total, tbl)
	start := -1
	res := make([]int, 0)
	occurence := make(map[string]int)
	for i := 0; i <= len(s)-k; i += k {
		word := s[i : i+k]
		fmt.Println(word)
		if _, ok := tbl[word]; !ok {
			start = -1
			continue
		}
		if start == -1 {
			start = i
			clearOccurence(occurence)
		}
		occurence[word] += 1
		if occurence[word] > tbl[word] {
			start = skip(s, start, k, word, occurence)
		}
		if i+k-start == total {
			res = append(res, start+offset)
		}
	}
	return res
}

func findSubstring3(s string, words []string) []int {
	if len(s) == 0 || len(words) == 0 {
		return []int{}
	}
	k := len(words[0])
	if len(s) < k {
		return []int{}
	}
	tbl := make(map[string]int)
	for _, word := range words {
		tbl[word] += 1
	}
	total := len(words) * k
	ans := make([]int, 0)
	for i := 0; i < k; i++ {
		res := find(s[i:], i, k, total, tbl)
		ans = append(ans, res...)
	}
	return ans
}

func main() {
	to_test := []string{
		"barfoothefoobarman",
		"pjzkrkevzztxductzzxmxsvwjkxpvukmfjywwetvfnujhweiybwvvsrfequzkhossmootkmyxgjgfordrpapjuunmqnxxdrqrfgkrsjqbszgiqlcfnrpjlcwdrvbumtotzylshdvccdmsqoadfrpsvnwpizlwszrtyclhgilklydbmfhuywotjmktnwrfvizvnmfvvqfiokkdprznnnjycttprkxpuykhmpchiksyucbmtabiqkisgbhxngmhezrrqvayfsxauampdpxtafniiwfvdufhtwajrbkxtjzqjnfocdhekumttuqwovfjrgulhekcpjszyynadxhnttgmnxkduqmmyhzfnjhducesctufqbumxbamalqudeibljgbspeotkgvddcwgxidaiqcvgwykhbysjzlzfbupkqunuqtraxrlptivshhbihtsigtpipguhbhctcvubnhqipncyxfjebdnjyetnlnvmuxhzsdahkrscewabejifmxombiamxvauuitoltyymsarqcuuoezcbqpdaprxmsrickwpgwpsoplhugbikbkotzrtqkscekkgwjycfnvwfgdzogjzjvpcvixnsqsxacfwndzvrwrycwxrcismdhqapoojegggkocyrdtkzmiekhxoppctytvphjynrhtcvxcobxbcjjivtfjiwmduhzjokkbctweqtigwfhzorjlkpuuliaipbtfldinyetoybvugevwvhhhweejogrghllsouipabfafcxnhukcbtmxzshoyyufjhzadhrelweszbfgwpkzlwxkogyogutscvuhcllphshivnoteztpxsaoaacgxyaztuixhunrowzljqfqrahosheukhahhbiaxqzfmmwcjxountkevsvpbzjnilwpoermxrtlfroqoclexxisrdhvfsindffslyekrzwzqkpeocilatftymodgztjgybtyheqgcpwogdcjlnlesefgvimwbxcbzvaibspdjnrpqtyeilkcspknyylbwndvkffmzuriilxagyerjptbgeqgebiaqnvdubrtxibhvakcyotkfonmseszhczapxdlauexehhaireihxsplgdgmxfvaevrbadbwjbdrkfbbjjkgcztkcbwagtcnrtqryuqixtzhaakjlurnumzyovawrcjiwabuwretmdamfkxrgqgcdgbrdbnugzecbgyxxdqmisaqcyjkqrntxqmdrczxbebemcblftxplafnyoxqimkhcykwamvdsxjezkpgdpvopddptdfbprjustquhlazkjfluxrzopqdstulybnqvyknrchbphcarknnhhovweaqawdyxsqsqahkepluypwrzjegqtdoxfgzdkydeoxvrfhxusrujnmjzqrrlxglcmkiykldbiasnhrjbjekystzilrwkzhontwmehrfsrzfaqrbbxncphbzuuxeteshyrveamjsfiaharkcqxefghgceeixkdgkuboupxnwhnfigpkwnqdvzlydpidcljmflbccarbiegsmweklwngvygbqpescpeichmfidgsjmkvkofvkuehsmkkbocgejoiqcnafvuokelwuqsgkyoekaroptuvekfvmtxtqshcwsztkrzwrpabqrrhnlerxjojemcxel",
	}

	param2 := [][]string{
		[]string{"foo", "bar"},
		[]string{"dhvf","sind","ffsl","yekr","zwzq","kpeo","cila","tfty","modg","ztjg","ybty","heqg","cpwo","gdcj","lnle","sefg","vimw","bxcb"},
	}

	for i:=0; i < len(to_test); i+=1{
		rst := findSubstring3(to_test[i], param2[i])
		fmt.Println(to_test[i], param2[i], rst)
	}
}

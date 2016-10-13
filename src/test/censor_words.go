/*
双数组实现的trie树, 查询时间只与输入的字符串相关
假设树高为len(a) = 关键字最大长度，最大为10
message 的长度为len(b)
最坏情况时间复杂度=10 * len(b)
*/

package test

import (
	"bufio"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

type node struct {
	code               rune
	depth, left, right int // 用于获取子节点
}

type Darts struct {
	Base     []int // base数组
	Check    []int // check数组
	KeyCount int
}

type dartsBuild struct {
	darts        Darts
	size         int
	key          [][]rune
	nextCheckPos int
	used         []bool
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Build(key [][]rune) Darts {
	var d = new(dartsBuild)
	d.key = key
	d.resize(512)
	d.darts.Base[0] = 1
	d.darts.KeyCount = 0
	d.nextCheckPos = 0

	var rootNode node
	rootNode.depth = 0
	rootNode.left = 0
	rootNode.right = len(key)
	siblings := d.fetch(rootNode)
	d.insert(siblings)
	return d.darts
}

func (d *dartsBuild) resize(newSize int) {
	if newSize > cap(d.darts.Base) {
		d.darts.Base = append(d.darts.Base, make([]int, (newSize-len(d.darts.Base)))...)
		d.darts.Check = append(d.darts.Check, make([]int, (newSize-len(d.darts.Check)))...)
		d.used = append(d.used, make([]bool, (newSize-len(d.used)))...)
	} else {
		d.darts.Base = d.darts.Base[:newSize]
		d.darts.Check = d.darts.Check[:newSize]
		d.used = d.used[:newSize]
	}
}

func (d *dartsBuild) fetch(parent node) []node {
	var siblings = make([]node, 0, 2)
	var prev rune = 0

	for i := parent.left; i < parent.right; i++ {
		if len(d.key[i]) < parent.depth {
			continue
		}

		tmp := d.key[i]

		var cur rune = 0
		// 如果parenet不是叶节点，则children的code值为编码+1，否则为0
		if len(d.key[i]) != parent.depth {
			cur = tmp[parent.depth] + 1
		}

		if prev > cur {
			fmt.Println(prev, cur, i, parent.depth, d.key[i])
			fmt.Println(d.key[i])
			panic("fetch error 1")
			return siblings[0:0]
		}

		if cur != prev || len(siblings) == 0 {
			var tmpNode node
			tmpNode.depth = parent.depth + 1
			tmpNode.code = cur
			tmpNode.left = i
			//更新上个节点的检索范围
			if len(siblings) != 0 {
				siblings[len(siblings)-1].right = i
			}

			siblings = append(siblings, tmpNode)
		}

		prev = cur
	}

	// 更新最后一个child的最大检索值
	if len(siblings) != 0 {
		siblings[len(siblings)-1].right = parent.right
	}

	return siblings
}

//核心算法，递归生成check和base值
//满足: base[s] + siblings[0].code + 1 = t 和check[t] = base[s]
//每次insert时，children的check位置都会固定下来并得到一个begin值，parent的base值等于第一个child的插入位置=beign值
func (d *dartsBuild) insert(siblings []node) int {
	var begin int = 0

	var pos int = max(int(siblings[0].code)+1, d.nextCheckPos) - 1
	var nonZeroNum int = 0
	first := false

	// 如果空间不足，补足空间
	if len(d.darts.Base) <= pos {
		d.resize(pos + 1)
	}

	for {
	next:
		pos++

		// 如果空间不足，补足空间
		if len(d.darts.Base) <= pos {
			d.resize(pos + 1)
		}

		if d.darts.Check[pos] > 0 {
			nonZeroNum++
			continue
		} else if !first {
			d.nextCheckPos = pos
			first = true
		}

		begin = pos - int(siblings[0].code)
		if len(d.darts.Base) <= (begin + int(siblings[len(siblings)-1].code)) {
			d.resize(begin + int(siblings[len(siblings)-1].code) + 400)
		}

		if d.used[begin] {
			continue
		}

		for i := 1; i < len(siblings); i++ {
			if begin+int(siblings[i].code) >= len(d.darts.Base) {
				fmt.Println(len(d.darts.Base), begin+int(siblings[i].code), begin+int(siblings[len(siblings)-1].code))
			}
			if 0 != d.darts.Check[begin+int(siblings[i].code)] {
				goto next
			}
		}
		break
	}

	// 太多位置被占用，更新检查位置
	if float32(nonZeroNum)/float32(pos-d.nextCheckPos+1) >= 0.95 {
		d.nextCheckPos = pos
	}
	d.used[begin] = true
	d.size = max(d.size, begin+int(siblings[len(siblings)-1].code)+1)

	for i := 0; i < len(siblings); i++ {
		d.darts.Check[begin+int(siblings[i].code)] = begin
	}

	// 递归处理下一层
	for i := 0; i < len(siblings); i++ {
		newSiblings := d.fetch(siblings[i])

		// 如果是叶节点，多一个节点end，其base为负
		if len(newSiblings) == 0 {
			d.darts.Base[begin+int(siblings[i].code)] = -d.darts.KeyCount - 1
			d.darts.KeyCount++
		} else {
			// 非叶节点，递归调用
			h := d.insert(newSiblings)
			d.darts.Base[begin+int(siblings[i].code)] = h
		}
	}
	return begin
}

// 精确查找某个字符串是否存在
func (d Darts) ExactMatch(key []rune) (bool, int) {
	b := d.Base[0]
	var p int // 状态p
	for i := 0; i < len(key); i++ {
		p = b + int(key[i]) + 1 // 新的状态
		if b == d.Check[p] {
			b = d.Base[p]
		} else {
			return false, i
		}
	}

	p = b // 是否存在这个单词：必须要以key的最后一个rune结尾
	n := d.Base[p]
	if b == d.Check[p] && n < 0 {
		return true, -1
	}
	return false, -1
}

// 查找字符串里面是否有违禁子串
func (d Darts) Search(key string) bool {
	key = strings.ToLower(key)
	allBytes := []rune(key)
	keyLength := len(key)
	for i := 0; i < keyLength; i++ {
		for j := i; j < keyLength; j++ {
			isFound, failIdx := d.ExactMatch(allBytes[i : j+1])
			if isFound {
				return true
			}
			if failIdx == 0 {
				break
			}
		}
	}
	return false
}

type dartsKey struct {
	key   []rune /*Key_type*/
	value int
}
type dartsKeySlice []dartsKey

func (r dartsKeySlice) Len() int {
	return len(r)
}
func (r dartsKeySlice) Less(i, j int) bool {
	var l int
	if len(r[i].key) < len(r[j].key) {
		l = len(r[i].key)
	} else {
		l = len(r[j].key)
	}

	for m := 0; m < l; m++ {
		if r[i].key[m] < r[j].key[m] {
			return true
		} else if r[i].key[m] == r[j].key[m] {
			continue
		} else {
			return false
		}
	}
	if len(r[i].key) < len(r[j].key) {
		return true
	}
	return false
}
func (r dartsKeySlice) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func Import(inFile, outFile string) (Darts, error) {
	// 输入文件
	unifile, erri := os.Open(inFile)
	if erri != nil {
		return Darts{}, erri
	}
	defer unifile.Close()

	// 存储到本地文件一份，后续支持直接load
	ofile, erro := os.Create(outFile)
	if erro != nil {
		return Darts{}, erro
	}
	defer ofile.Close()

	// 读取文件，对所有key进行排序
	dartsKeys := make(dartsKeySlice, 0, 100000)
	uniLineReader := bufio.NewReaderSize(unifile, 400)
	line, _, bufErr := uniLineReader.ReadLine()
	for nil == bufErr {
		// 去除左右空白
		key := []rune(strings.TrimSpace(string(line)))
		value := 1
		dartsKeys = append(dartsKeys, dartsKey{key, value})
		line, _, bufErr = uniLineReader.ReadLine()
	}
	sort.Sort(dartsKeys)
	keys := make([][]rune, len(dartsKeys))
	values := make([]int, len(dartsKeys))
	for i := 0; i < len(dartsKeys); i++ {
		keys[i] = dartsKeys[i].key
		values[i] = dartsKeys[i].value
	}

	// 开始构建darts
	fmt.Printf("input dict length: %v\n", len(dartsKeys))
	round := len(keys)
	var d Darts
	d = Build(keys[:round])

	// 确保所有key都被加入
	fmt.Printf("build out length %v\n", len(d.Base))
	t := time.Now()
	for i := 0; i < round; i++ {
		isFound, _ := d.ExactMatch(keys[i])
		if isFound != true {
			fmt.Printf("missing key %s\n", string(keys[i]))
			//err := fmt.Errorf("missing key %s", string(keys[i]))
			continue
		}
	}
	fmt.Println(time.Since(t))

	// 写入本地文件
	enc := gob.NewEncoder(ofile)
	enc.Encode(d)

	// 返回
	return d, nil
}

func ImportFromJson(jsonStr string) (Darts, error) {
	var decodeStr []string
	err := json.Unmarshal([]byte(jsonStr), &decodeStr)
	if err != nil {
		panic(fmt.Sprintf("load censorwords err: %s", err.Error()))
	}

	dartsKeys := make(dartsKeySlice, 0, len(decodeStr))
	for _, str := range decodeStr {
		key := []rune(str)
		value := 1 //权重，可不用
		dartsKeys = append(dartsKeys, dartsKey{key, value})

	}
	sort.Sort(dartsKeys)
	keys := make([][]rune, len(dartsKeys))
	values := make([]int, len(dartsKeys))
	for i := 0; i < len(dartsKeys); i++ {
		keys[i] = dartsKeys[i].key
		values[i] = dartsKeys[i].value
	}
	// 开始构建darts
	fmt.Printf("input dict length: %v\n", len(dartsKeys))
	round := len(keys)
	var d Darts
	d = Build(keys[:round])

	// 确保所有key都被加入
	fmt.Printf("build out length %v\n", len(d.Base))
	t := time.Now()
	for i := 0; i < round; i++ {
		isFound, _ := d.ExactMatch(keys[i])
		if isFound != true {
			fmt.Printf("missing key %s\n", string(keys[i]))
			//err := fmt.Errorf("missing key %s", string(keys[i]))
			continue
		}
	}
	fmt.Println(time.Since(t))
	return d, nil
}

/*
从已有文件直接加载到内存
*/
func Load(filename string) (Darts, error) {
	var dict Darts
	file, err := os.Open(filename)
	if err != nil {
		return Darts{}, err
	}
	defer file.Close()

	dec := gob.NewDecoder(file)
	dec.Decode(&dict)
	return dict, nil
}

func Dumps(d Darts) {
	for i, b := range d.Base {
		if b == 0 {
			continue
		}
		fmt.Printf("i=%d b=%d c=%d\n", i, b, d.Check[i])
	}
}

func Init() (Darts, error) {
	jsonStr := GetJsonCensorWords()
	d, err := ImportFromJson(jsonStr)
	return d, err
}

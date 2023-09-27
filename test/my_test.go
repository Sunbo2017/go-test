package test

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"testing"
	"time"
)

// map append问题
func TransformResultFormat(columns []string, values [][]string) {

	var result []map[string]interface{}
	var TransformResult = make(map[string]interface{})

	fmt.Println("yyyy", len(values))
	for i := 0; i < len(values); i++ {
		for j, v := range columns {
			TransformResult[v] = values[i][j]
		}
		fmt.Println("xxx", TransformResult)
		result = append(result, TransformResult)
	}
	//result = append(result, TransformResult)

	fmt.Println(result)
}

func TestMapTransfor(t *testing.T) {
	c := []string{"name", "grade", "uuid"}
	v := [][]string{
		{"xiaoming", "5", "59525F6C427F339F88B5C81FE9DC3671"},
		{"xiaoli", "51", "D08594DF1F983809A890CF024E64B06B"},
	}
	TransformResultFormat(c, v)
}

// 字节面试题：台阶积水问题，类似滑动窗口思想
// 忽略第0个台阶，从第1个台阶开始算，后边台阶如果矮于第一个台阶，差值即为积水量，
// 直到高于第一个台阶的新台阶出现，第一次循环结束，
// 然后从这个新台阶开始第二次循环
// 该思想存在局限性，不能当作正确答案
func TestStepsWater(t *testing.T) {
	// steps := []int{1, 0, 2, 1, 2, 3, 0, 1, 2, 4}
	steps := []int{0, 0, 2, 1, 2, 3, 0, 1, 3, 2}
	water := 0
	i, j := 0, 1
	for i < len(steps) {
		tempWater := water
		for i+j < len(steps) && steps[i+j] <= steps[i] {
			water += steps[i] - steps[i+j]
			j++
		}
		// 如果一直到最后一个台阶也不高于起始台阶，则无法积水
		if i+j == len(steps)-1 && steps[i+j] <= steps[i] {
			water = tempWater
		}
		// 直接跳到新台阶进行下一次循环
		i += j
		j = 1
	}
	fmt.Println(water)
}

type step struct {
	val      int //当前台阶高度
	leftMax  int //左侧最高台阶高度
	rightMax int //右侧最高台阶高度
}

func makeSteps(steps []int) []step {
	list := make([]step, len(steps))
	for i, v := range steps {
		s := step{val: v}
		list[i] = s
	}
	for i := 1; i < len(steps)-1; i++ {
		list[i].leftMax = Max(list[i-1].leftMax, list[i-1].val)
	}
	for i := len(steps) - 2; i > 0; i-- {
		list[i].rightMax = Max(list[i+1].rightMax, list[i+1].val)
	}
	return list
}

func countStepWater(steps []int) int {
	water := 0
	stepList := makeSteps(steps)
	for i := 1; i < len(steps)-1; i++ {
		left := stepList[i].leftMax - stepList[i].val
		right := stepList[i].rightMax - stepList[i].val
		if left > 0 && right > 0 {
			increment := Min(left, right)
			water += increment
		}
	}
	return water
}

func TestStepWater1(t *testing.T) {
	steps := []int{0, 0, 2, 1, 2, 3, 0, 1, 3, 2}
	water := countStepWater(steps)
	t.Logf("water:%v", water)
}

func TestSort(t *testing.T) {
	ints := []int{0, 5, 2, 1, 3, 4, 6, 9, 8, 7}
	sort.Ints(ints)
	fmt.Println(ints)
}

var (
	letterMap = []string{
		" ",    //0
		"",     //1
		"abc",  //2
		"def",  //3
		"ghi",  //4
		"jkl",  //5
		"mno",  //6
		"pqrs", //7
		"tuv",  //8
		"wxyz", //9
	}
	res = []string{}
)

// LeetCode-cookbook-17：根据手机数字按钮返回对应数字可生成的所有字母组合
func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	res = []string{}
	findCombination(&digits, 0, "")
	return res
}
func findCombination(digits *string, index int, s string) {
	if index == len(*digits) {
		res = append(res, s)
		return
	}
	num := (*digits)[index]
	letter := letterMap[num-'0']
	for i := 0; i < len(letter); i++ {
		findCombination(digits, index+1, s+string(letter[i]))
	}
	return
}

func TestCombinations(t *testing.T) {
	letterCombinations("456")
	fmt.Println(res)
	byten := byte('4')
	r := byten - '1'
	fmt.Println(r)
}

// var res [][]string

// 已知每一个字母可以用 1~26 表示，给定一个数字组成的字符串，问它可以表示多少种字母组合。
// 123456215
func numString(str string, lStr []string) int {
	for i := 0; i < len(str); {

	}

	return 0
}

// 输入一个数组和一个目标值 T，判断数组中是否存在两个数的和为 T
// 最简单方法可以直接双层循环判断和，O(n^2)
// 可以使用map的k，v分别记录元素值和差值，O（n）
func judge2Sum(arr []int, t int) (v1, v2 int) {
	resMap := map[int]int{}
	for _, v := range arr {
		if _, ok := resMap[v]; !ok {
			resMap[t-v] = v
		} else {
			return v, resMap[v]
		}
	}
	return 0, 0
}

// 升级：输入一个数组和一个目标值 T，判断数组中是否存在某些数的和为 T。

func Test2Sum(t *testing.T) {
	array := []int{1, 2, 3, 4, 6, 5, 8, 9}
	// 6, 4
	v1, v2 := judge2Sum(array, 10)
	fmt.Println(v1)
	fmt.Println(v2)
}

/**
 * 输入一个数组和一个目标值 T，判断数组中是否存在两个数的和为 T，返回两数下标
 * @param numbers int整型一维数组
 * @param target int整型
 * @return int整型一维数组
 */
func twoSumIndex(numbers []int, target int) []int {
	// key:数组元素,val:元素下标
	nmap := make(map[int]int, 0)
	for i, v := range numbers {
		if val, ok := nmap[target-v]; ok {
			return []int{val + 1, i + 1}
		} else {
			nmap[v] = i
		}
	}
	return []int{}
}

func TestTwoSumIndex(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 6
	res := twoSumIndex(nums, target)
	t.Log(res)
}

// 一个细胞的寿命是5min 他会在2min和4min 分别分裂出一个新细胞，请问n min后 ，有多少细胞
var sumn = 1

func sum(n int) int {
	for i := 1; i <= n; i++ {
		if i%2 == 0 || i%4 == 0 {
			// sumn += sum(n-i)
			sumn *= 2
		}
		if i%5 == 0 {
			div := i / 5
			sumn = sumn - 2*div
		}
	}
	return sumn
}

func TestSum(t *testing.T) {
	n := 15
	fmt.Println(sum(n))
}

// 判断n是否是质数
// 按素数的定义，也就是只有 1 与本身可以整除，所以可以用 2→ i-1 去除 i，如果都除不尽，i 就是素数。
// 观点对，但却笨拙。当 i>2 时，有哪一个数可以被 i-1 除尽的？没有！
// 为什么？如果 i 不是质数，那么 i=a×b，此时 a 与 b 既不是 i 又不是 1；
// 正因为 a>1，a 至少为 2，因此 b 最多也是 i/2 而已，去除 i 的数用不着是 2→ i-1，而用 2→ i/2 就可以了。
// 不但如此，因为 i=a×b，a 与 b 不能大于 sqrt(i)，为什么呢？
// 如果 a>sqrt(i)，b>sqrt(i)，于是 a×b > sqrt(i)*sqrt(i) = i，因此就都不能整除i了。
// 如果i不是质数，它的因子最大就是 sqrt(i)；换言之，用 2→ sqrt(i)去检验就行了
func judgePrime(n int) bool {
	if n == 1 || n == 0 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// 输出一亿内的所有素数
// 十个协程并发，用时38秒左右
func findPrimeAll(n int) {
	total := 0
	numChan := make(chan int, n)
	resChan := make(chan string)
	exitChan := make(chan int, 10)
	for i := 2; i <= n; i++ {
		numChan <- i
	}
	close(numChan)

	for i := 0; i < 10; i++ {
		go findPrime(numChan, exitChan, resChan, i)
	}

	// 等待结束信号
	go func() {
		count := 0
		for v := range exitChan {
			count++
			fmt.Printf("channel:%v is finish,count=%v;\n", v, count)
			if count == 10 {
				fmt.Println("close resChan")
				close(resChan)
				close(exitChan)
			}
		}
	}()

	// for v := range resChan {
	// 	fmt.Printf("%v,",v)
	// 	total++
	// }
	for {
		_, ok := <-resChan
		if !ok {
			break
		} else {
			//fmt.Println(v)
			total++
		}
	}

	fmt.Printf("total:%v \n", total)
}

func findPrime(in, exit chan int, out chan string, num int) {
	// fmt.Printf("channel id:%v /n", num)
	for v := range in {
		if res := judgePrime(v); res {
			// fmt.Println(v)
			out <- fmt.Sprintf("goroutine-%v:%v", num, v)
		}
	}
	exit <- num
}

func TestFindPrime(t *testing.T) {
	start := time.Now().Unix()

	findPrimeAll(100000000)

	// _CalcPrimes()
	// fmt.Println(_Primes)
	// fmt.Println(100000000, "以内的素数个数为", _N)

	//findPrimeBySieve(100000000)
	fmt.Println("finish...")
	end := time.Now().Unix()

	cost := end - start
	t.Logf("cost:%v", cost)
}

var _Primes []uint64 = []uint64{
	2, 3, 5, 7, 11, 13, 17, 19, 23, 29,
	31, 37, 41, 43, 47, 53, 59, 61, 67, 71,
	73, 79, 83, 89, 97,
}

var _N int

// 设n>1为整数，m为整数，且n≤m<n^2，如果小于n的所有素数都不是m的因子，则m为素数。
func _CalcPrimes() {
	N := len(_Primes)
	i := 0

	for n := uint64(101); n < 10000; n += 2 {
		for i = 1; i < N; i++ { // i从1开始，因为2必然不整除n
			if n%_Primes[i] == 0 {
				break
			}
		}
		if i == N {
			_Primes = append(_Primes, n)
		}
	}

	N = len(_Primes)

	for n := uint64(10001); n < 100000000; n += 2 {
		for i = 1; i < N; i++ {
			if n%_Primes[i] == 0 {
				break
			}
		}
		if i == N {
			_Primes = append(_Primes, n)
		}
	}

	N = len(_Primes)
	_N = N
}

// 生成n个数的channel
func generate(ch chan int, n int) {
	for i := 2; i <= n; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func filter(in, out chan int, prime int) {
	for {
		i := <-in // Receive value of new variable 'i' from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to channel 'out'.
		}
	}
}

// 求素数：用小于n的所有素数去除n,如果都不能整除，则n为素数
// 一个素数不能整除的那个比它自身大的最小的那个数就是素数
// The prime sieve: Daisy-chain filter processes together.
// 网上都是这套代码，实际效率贼差，不可取
func findPrimeBySieve1(n int) {
	ch := make(chan int) // Create a new channel.
	go generate(ch, n)   // Start generate() as a goroutine.
	for {
		prime := <-ch
		// fmt.Printf("prime:%v,", prime)
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
}

// 筛法求素数：依次去掉已知素数的所有倍数
// 筛法确实强，单协程执行完只需1秒多时间
func findPrimeBySieve(n int) {
	isPrime := make([]bool, n)
	for i := 0; i < n; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i < n; i++ {
		if isPrime[i] {
			for j := i * i; j < n; j += i {
				isPrime[j] = false
			}
		}
	}

	count := 0
	for i := 2; i < n; i++ {
		if isPrime[i] {
			// fmt.Printf("%v,", i)
			count++
		}
	}

	fmt.Println("total:", count)
}

//金山云面试题：两个协程交替打印奇数偶数，必须保证按序输出
func TestNum(t *testing.T) {
	fmt.Println("Hello, World!")
	a, b := make(chan int), make(chan int)
	go printNum2(a, b)
	go printNum1(a, b)
	//先起协程，后向channel发数据，否则死锁
	b <- 1
	time.Sleep(2 * time.Second)
}

func printNum1(intChan1, intChan2 chan int) {
	//注意此处i必须为偶数，否则不会打印，也不会向channel发送信号
	for i := 2; i <= 100; i += 2 {
		if _, ok := <-intChan1; ok {
			if i%2 == 0 {
				fmt.Printf("chan2:%v\n", i)
				intChan2 <- i
			}
		}
	}
}

func printNum2(intChan1, intChan2 chan int) {
	//注意此处i必须为奇数，否则不会打印，也不会向channel发送信号
	for i := 1; i <= 100; i += 2 {
		if _, ok := <-intChan2; ok {
			if i%2 != 0 {
				fmt.Printf("chan1:%v\n", i)
				intChan1 <- i
			}
		}
	}
}

//血的教训，可以默写下来才叫掌握了算法思想
func reverseListNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	var pre *ListNode = nil
	for cur != nil {
		// pre, cur, cur.Next = cur, cur.Next, pre

		nxt := cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}

//金山云二面：3个协程交替打印1，2，3，按序输出
var wwg sync.WaitGroup

func TestPrint3(t *testing.T) {
	fmt.Println("Hello, World!")

	chan1, chan2, chan3 := make(chan int), make(chan int), make(chan int)
	//wg := sync.WaitGroup()
	wwg.Add(6)
	go printNum11(chan1, chan2)
	go printNum12(chan2, chan3)
	go printNum13(chan3, chan1)
	chan1 <- 1

	wwg.Wait()
	close(chan1)
	close(chan2)
	close(chan3)
	fmt.Println("FINISH")
}

func printNum11(chan1, chan2 chan int) {
	for i := 0; i < 2; i++ {
		if _, ok := <-chan1; ok {
			fmt.Println(1)
			wwg.Done()
			chan2 <- 1
		}
	}
}

func printNum12(chan1, chan2 chan int) {
	for i := 0; i < 2; i++ {
		if _, ok := <-chan1; ok {
			fmt.Println(2)
			wwg.Done()
			chan2 <- 1
		}
	}
}

func printNum13(chan1, chan2 chan int) {
	for i := 0; i < 2; i++ {
		if _, ok := <-chan1; ok {
			fmt.Println(3)

			// if i == 1 {
			// 	wg.Done()
			// 	break
			// }
			fmt.Println("F---")
			//必须先done，确保其它协程结束后不会再向另一个协程发出信号
			wwg.Done()
			chan2 <- 1
		}
	}
}

//判断数独是否有效
func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		m1 := make(map[byte]bool)
		m2 := make(map[byte]bool)
		m3 := make(map[byte]bool)
		fmt.Printf("i: %d, num[i]: %v\n", i, board[i])

		// 判断每一行是否重复
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				// fmt.Printf("j1: %d, num[j]: %v\n", j, board[i][j])
				if m1[board[i][j]] {
					return false
				}
				m1[board[i][j]] = true
			}

			// 判断每一列是否重复
			if board[j][i] != '.' {
				// fmt.Printf("j2: %d, num[j]: %v\n", j, board[j][i])
				if m2[board[j][i]] {
					return false
				}
				m2[board[j][i]] = true
			}

			// 判断9宫格内的数据是否重复
			row := (i%3)*3 + j%3
			col := (i/3)*3 + j/3
			if board[row][col] != '.' {
				fmt.Printf("board[%d][%d] = %v;", row, col, board[row][col]-'0')
				// fmt.Printf("j3: %d, num[j]: %v\n", j, board[row][col])
				if m3[board[row][col]] {
					return false
				}
				m3[board[row][col]] = true
			}
		}
		fmt.Println("next row...")
	}
	return true
}

func TestValidSudoku(t *testing.T) {
	//'0'的ascii码值为48，所以'5'的码值为48+5=53
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	sudoku := isValidSudoku(board)
	fmt.Println(sudoku)
}

func producer(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Printf("produce:%v\n", i)
		time.Sleep(2 * time.Second)
	}
}

func consumer(id int, ch chan int, ch1 chan int) {
	for {
		select {
		case v1 := <-ch1:
			fmt.Printf("stop:%v\n", v1)
			break
		case v2 := <-ch:
			fmt.Printf("consumer-%v,data:%v\n", id, v2)
		}
	}
}

func TestProduce(t *testing.T) {
	ch := make(chan int, 10)
	ch1 := make(chan int)
	ch2 := make(chan int)

	go producer(ch)
	go consumer(1, ch, ch1)
	go consumer(2, ch, ch2)

	time.Sleep(10 * time.Second)

	ch1 <- 1
	ch2 <- 2
}

func printLetterNumber() {
	var wg sync.WaitGroup
	wg.Add(2)

	letterCh := make(chan bool)
	numberCh := make(chan bool)

	// 打印字母的协程
	go func() {
		for i := 0; i < 26; i++ {
			if _, ok := <-letterCh; ok {
				fmt.Printf("%c \n", 'A'+i)
				numberCh <- true
			}
		}
		close(numberCh)
		wg.Done()
	}()

	// 打印数字的协程
	go func() {
		for i := 1; i <= 26; i++ {
			if _, ok := <-numberCh; ok {
				fmt.Printf("%d \n", i)
				//不加此判断会出现死锁，因为最后已经没有协程再接收letterCh的数据
				if i != 26 {
					letterCh <- true
				}
			}
		}
		close(letterCh)
		wg.Done()
	}()

	letterCh <- true
	wg.Wait()
	fmt.Println("finish----")
}

func TestPrintLN(t *testing.T) {
	printLetterNumber()
}

//s1=456783
//s2=78654977
func addBigNumber(s1, s2 string) string {
	l1, l2 := len(s1), len(s2)
	l := 0
	if l1 > l2 {
		l = l1
	} else {
		l = l2
	}
	list1, list2 := []uint8{}, []uint8{}
	for i := 0; i < l; i++ {
		list1 = append(list1, s1[i])
		list2 = append(list2, s2[i])
	}

	return ""
}

//一群朋友组队玩游戏，至少有5组人，一组至少2人，要求：
//1.每2个人组一队或者3个人组一队，每个人只能加到一个队伍里，不能落单
//2.2人队和3人队各自的队伍数均不得少于1，队伍中的人不能来自相同组
//3.随机组队，重复执行程序得到的结果不一样，总队伍数也不能一样
//4.必须有注释
//注：要同时满足条件1-4

/*
举例：
GroupList=[#小组列表
['少华','少平','少军','少安','少康'],
['福军','福堂','福民','福平','福心']
['小明','小红','小花','小丽','小强'],
['大壮','大力','大1','大2','大3'],
['阿花','阿朵','阿蓝','阿紫','阿红'],
['A','B','C','D','E'],
['一','二','三','四','五'],
]
*/

var groupList = [][]string{
	{"少华", "少平", "少军", "少安", "少康"},
	{"福军", "福堂", "福民", "福平", "福心"},
	{"小明", "小红", "小花", "小丽", "小强"},
	{"大壮", "大力", "大1", "大2", "大3"},
	{"阿花", "阿朵", "阿蓝", "阿紫", "阿红"},
	{"A", "B", "C", "D", "E"},
	{"一", "二", "三", "四", "五"},
}

//横坐标随机数最大值
var x = len(groupList)

type person struct {
	name string //姓名
	x    int    //横坐标
	y    int    //纵坐标
}

func (p person) setName() {
	p.name = groupList[p.x][p.y]
}

//从groupList中删除已加入新组的person
func (p person) deleteFromGroup() {
	list := groupList[p.x]
	lenY := len(list)

	//如果该组就剩这个小朋友一个人了，需要清空这个组
	//if lenY == 1 {
	//	//如果正好是原分组的最后一个组,直接移除这个组
	//	if p.x == x {
	//		groupList = groupList[:x-1]
	//	} else {
	//		//如果不是最后一个组，用最后一个组替换该分组
	//		lastGroup := groupList[x-1]
	//		groupList[p.x] = lastGroup
	//		//移除最后一个组
	//		groupList = groupList[:x-1]
	//	}
	//}

	if p.y != lenY-1 {
		//如果该小朋友不是原分组中最后一个人，记录最后一个人
		lastP := groupList[p.x][lenY-1]
		//把最后一个人放到这个小朋友的位置
		groupList[p.x][p.y] = lastP
		//截取新的小组，长度减一
		newList := groupList[p.x][:lenY-1]
		//更新到groupList中
		groupList[p.x] = newList
	} else {
		//如果该小朋友正好是原分组中最后一个人，直接将长度减一
		newList := groupList[p.x][:lenY-1]
		//更新到groupList中
		groupList[p.x] = newList
	}
}

//获取新的所有分组
func getGroupList1() [][]person {
	rand.Seed(time.Now().UnixNano())

	//计算总人数
	total := 0
	for _, v := range groupList {
		total += len(v)
	}
	//最多可分几个组
	high := total / 2
	if total%2 > 0 {
		high += 1
	}
	//最少要分几组
	low := total / 3
	if total%3 > 0 {
		low += 1
	}
	//随机产生总小组数
	sum := low + getRandN(high-low)
	fmt.Printf("total group==%v\n", sum)

	result := make([][]person, sum, sum)

	for total > 0 {
		//每轮循环会将所有小组中均随机放入一个小朋友
		for i := 0; i < sum; i++ {
			if total <= 0 {
				break
			}
			var personList []person
			if len(result[i]) == 0 {
				personList = make([]person, 0)
			} else {
				personList = result[i]
			}
			xIndex := getRandN(x)
			//检查随机获取的横坐标是否已经存在于该小组中
			for checkX(xIndex, personList) {
				xIndex = getRandN(x)
			}
			//检查横坐标对应分组是否还有成员，若无成员返回第一个有成员的横坐标，避免随机数一直落在空的分组导致死循环
			xIndex = checkGroupList(xIndex, personList)
			//fmt.Printf("x==%v \n", xIndex)
			yList := groupList[xIndex]
			//fmt.Printf("ylist==%v \n", yList)
			//fmt.Printf("glist==%v \n", groupList)

			//如果根据横坐标取到的小组为空，则跳过本次循环
			if len(yList) == 0 {
				continue
			}
			p := person{
				x: xIndex,
				y: getRandN(len(yList)),
			}
			p.name = groupList[p.x][p.y]

			personList = append(personList, p)

			result[i] = personList

			//从groupList中删除已加入新组的person
			p.deleteFromGroup()
			total--
		}
	}

	return result
}

//获取新的所有分组
func getGroupList() [][]person {
	rand.Seed(time.Now().UnixNano())

	var result [][]person
	//计算总人数
	total := 0
	for _, v := range groupList {
		total += len(v)
	}
	//容量，剩余未分组人数
	capacity := total

	//确保二人组和三人组至少有一个
	result = append(result, makeGroup1(), makeGroup2())
	capacity -= 5

	for capacity > 0 {
		//该小组有几个人
		c := getRandN(2) + 2
		personList := make([]person, c)
		for i := 0; i < c; i++ {
			xIndex := getRandN(x)
			//检查随机获取的横坐标是否已经存在于该小组中
			for checkX(xIndex, personList) {
				xIndex = getRandN(x)
			}
			//检查横坐标对应分组是否还有成员，若无成员返回第一个有成员的横坐标，避免随机数一直落在空的分组导致死循环
			xIndex = checkGroupList(xIndex, personList)

			fmt.Printf("x==%v \n", xIndex)
			yList := groupList[xIndex]
			fmt.Printf("ylist==%v \n", yList)
			fmt.Printf("glist==%v \n", groupList)
			p := person{
				x: xIndex,
				y: getRandN(len(yList)),
			}
			p.setName()
			personList = append(personList, p)

			//从groupList中删除已加入新组的person
			p.deleteFromGroup()
		}
		result = append(result, personList)

		fmt.Println(result)
		capacity -= c

		//如果最后只剩下一个小朋友未分组
		if capacity == 1 {
			//获取最后一个人的横坐标
			lastX := 0
			for i, v := range groupList {
				if len(v) > 0 {
					lastX = i
				}
			}
			p := person{x: lastX, y: 0}
			p.setName()
			for i, v := range result {
				if len(v) < 3 {
					if !checkX(lastX, v) {
						result[i] = append(result[i], p)
						return result
					}
				}
			}
		}
	}

	return result
}

func TestGroupList(t *testing.T) {
	res := getGroupList1()
	t.Log(len(res))
	for i, v := range res {
		t.Logf("group-%v:%+v\n", i, v)
	}
}

//检查随机获取的横坐标是否已经存在于该小组中
func checkX(x int, list []person) bool {
	for _, v := range list {
		if x == v.x {
			return true
		}
	}

	return false
}

//检查横坐标对应分组是否还有成员，若无成员返回第一个有成员的横坐标
func checkGroupList(x int, pList []person) int {
	fmt.Printf("leny==%v \n", len(groupList[x]))
	if len(groupList[x]) == 0 {
		for i, v := range groupList {
			if len(v) > 0 {
				if !checkX(i, pList) {
					return i
				}
			}
		}
	}
	return x
}

func checkGroupList1(x int) int {
	fmt.Printf("leny==%v \n", len(groupList[x]))
	if len(groupList[x]) == 0 {
		for i, v := range groupList {
			if len(v) > 0 {
				return i
			}
		}
	}
	return x
}

func makeGroup1() []person {
	rand.Seed(time.Now().UnixNano())

	var group1 []person
	lastX := getRandN(x)
	y := getRandN(len(groupList[lastX]))
	p1 := person{
		name: groupList[lastX][y],
		x:    lastX,
		y:    y,
	}
	group1 = append(group1, p1)
	p1.deleteFromGroup()

	x1 := getRandN(x)
	for x1 == lastX {
		x1 = getRandN(x)
	}
	y = getRandN(len(groupList[x1]))
	p2 := person{
		name: groupList[x1][y],
		x:    x1,
		y:    y,
	}
	group1 = append(group1, p2)
	p2.deleteFromGroup()

	return group1
}

func makeGroup2() []person {
	rand.Seed(time.Now().UnixNano())

	var group []person
	lastX := getRandN(x)
	y := getRandN(len(groupList[lastX]))
	p1 := person{
		name: groupList[lastX][y],
		x:    lastX,
		y:    y,
	}
	group = append(group, p1)
	p1.deleteFromGroup()

	x1 := getRandN(x)
	for x1 == lastX {
		x1 = getRandN(x)
	}
	y = getRandN(len(groupList[x1]))
	p2 := person{
		name: groupList[x1][y],
		x:    x1,
		y:    y,
	}
	group = append(group, p2)
	p2.deleteFromGroup()

	x2 := getRandN(x)
	for x2 == lastX || x2 == x1 {
		x2 = getRandN(x)
	}
	y = getRandN(len(groupList[x2]))
	p3 := person{
		name: groupList[x2][y],
		x:    x2,
		y:    y,
	}
	group = append(group, p3)
	p3.deleteFromGroup()

	return group
}

func getRandN(n int) int {
	//if n == 0 {
	//	return 0
	//}
	r := rand.Intn(n)
	return r
}

func TestRandN(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	r := getRandN(1)
	t.Log(r)
}

func findLatestNum(arr []int, n int) []int {
	l, h := 0, len(arr)
	for l < h {
		r := partion1(arr, l, h)
		if r+1 == n {
			return arr[:n]
		}
		if r+1 > n {
			h = r - 1
		} else {
			l = r + 1
		}
	}
	return arr[:n]
}

func partion1(arr []int, l, h int) int {
	pivot := arr[h]
	i := l - 1
	for j := l; j < h; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[h] = arr[h], arr[i+1]
	return i + 1
}

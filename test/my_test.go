package test

import (
	"fmt"
	"sort"
	"testing"
	"time"
	"sync"
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
	val int  //当前台阶高度
	leftMax int  //左侧最高台阶高度
	rightMax int  //右侧最高台阶高度
}

func makeSteps(steps []int) []step {
	list := make([]step, len(steps))
	for i, v := range steps {
		s := step{val: v}
		list[i] = s
	}
	for i:=1; i<len(steps)-1; i++ {
		list[i].leftMax = Max(list[i-1].leftMax, list[i-1].val)
	}
	for i:=len(steps)-2; i>0; i-- {
		list[i].rightMax = Max(list[i+1].rightMax, list[i+1].val)
	}
	return list
}

func countStapWater (steps []int) int {
	water := 0
	stepList := makeSteps(steps)
	for i:=1; i<len(steps)-1; i++ {
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
	water := countStapWater(steps)
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
	for i:=0; i<len(str); {

	}
	
	return 0
}

// 输入一个数组和一个目标值 T，判断数组中是否存在两个数的和为 T
// 最简单方法可以直接双层循环判断和，O(n^2)
// 可以使用map的k，v分别记录元素值和差值，O（n）
func judege2Sum(arr []int, t int) (v1,v2 int){
	resMap := map[int]int{}
	for _, v := range arr{
		if _, ok := resMap[v]; !ok{
			resMap[t - v] = v
		} else{
			return v, resMap[v]
		}
	}
	return 0, 0
}

// 升级：输入一个数组和一个目标值 T，判断数组中是否存在某些数的和为 T。



func Test2Sum(t *testing.T){
	array := []int{1,2,3,4,6,5,8,9}
	// 6, 4
	v1, v2 := judege2Sum(array, 10)
	fmt.Println(v1)
	fmt.Println(v2)
}

/**
  * 输入一个数组和一个目标值 T，判断数组中是否存在两个数的和为 T，返回两数下标
  * @param numbers int整型一维数组 
  * @param target int整型 
  * @return int整型一维数组
*/
func twoSumIndex( numbers []int ,  target int ) []int {
    // write code here
    nmap := make(map[int]int,0)
    for i, v := range numbers {
        if val, ok := nmap[target-v]; ok {
            return []int{val+1, i+1}
        }else {
            nmap[v] = i
        }
    }
    return []int{}
}

func TestTwoSumIndex(t *testing.T) {
	nums := []int{3,2,4}
	target := 6
	res := twoSumIndex(nums, target)
	t.Log(res)
}

// 一个细胞的寿命是5min 他会在2min和4min 分别分裂出一个新细胞，请问n min后 ，有多少细胞 
var sumn = 1
func sum(n int) int{
    for i:=1;i<=n;i++{
        if i%2 == 0 || i%4==0 {
            // sumn += sum(n-i)
			sumn *= 2
        }
        if i%5==0 {
			div := i/5
            sumn = sumn - 2*div
        }
    }
    return sumn
}

func TestSum(t *testing.T){
	n := 15
	fmt.Println(sum(n))
}

// 判断n是否是质数
// 按素数的定义，也就是只有 1 与本身可以整除，所以可以用 2→ i-1 去除 i，如果都除不尽，i 就是素数。
// 观点对，但却笨拙。当 i>2 时，有哪一个数可以被 i-1 除尽的？没有！
// 为什么？如果 i 不是质数，那么 i=a×b，此地 a 与 b 既不是 i 又不是 1；
// 正因为 a>1，a 至少为 2，因此 b 最多也是 i/2 而已，去除 i 的数用不着是 2→ i-1，而用 2→ i/2 就可以了。
// 不但如此，因为 i=a×b，a 与 b 不能大于 sqrt(i)，为什么呢？
// 如果 a>sqrt(i)，b>sqrt(i)，于是 a×b > sqrt(i)*sqrt(i) = i，因此就都不能整除i了。
// 如果i不是质数，它的因子最大就是 sqrt(i)；换言之，用 2→ sqrt(i)去检验就行了
func judgePrime(n int) bool {
	if n==1 || n==0 {
		return false
	}
	for i:=2;i*i<=n;i++ {
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
	for i:=2;i<=n;i++ {
		numChan <- i
	}
	close(numChan)

	for i:=0;i<10;i++ {
		go findPrime(numChan, exitChan,resChan, i)
	}

	// 等待结束信号
	go func () {
		count := 0
		for v:= range exitChan {
			count++
			fmt.Printf("channel:%v is finish,count=%v;", v, count)
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
		_, ok := <- resChan
		if !ok {
			break
		} else {
			total++
		}
	}
	
	fmt.Printf("total:%v \n", total)
}

func findPrime(in,exit chan int, out chan string, num int) {
	// fmt.Printf("channel id:%v /n", num)
	for v := range in {
		if res := judgePrime(v); res {
			// fmt.Println(v)
			out <- fmt.Sprintf("%v:%v", num, v)
		}
	}
	exit <- num
}

func TestFindPrime(t *testing.T) {
	start := time.Now().Unix()
	// findPrimeAll(100000000)

	// _CalcPrimes()
    // fmt.Println(_Primes)
    // fmt.Println(100000000, "以内的素数个数为", _N)

	findPrimeBySieve(100000000)
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
    for i := 2; i<=n; i++ {
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
    go generate(ch, n)      // Start generate() as a goroutine.
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
	for i:=0;i<n;i++ {
		isPrime[i] = true
	}
	for i := 2; i * i < n; i++ {
		if isPrime[i] {
			for j := i * i; j < n; j += i {
				isPrime[j] = false;
			}
		}
	}
	
	count := 0;
	for i := 2; i < n; i++ {
		if isPrime[i] {
			// fmt.Printf("%v,", i)
			count++
		} 
	}
	
	fmt.Println("total:", count)
}


//字节一面
//一只青蛙一次可以跳上1级台阶，也可以跳上2级。求该青蛙跳上一个n级的台阶总共有多少种跳法
//递归
func jumpFloor1(N int) int {
	if N <= 0 {
	  return 0
	}
	if N == 1 || N == 2 {
	  return N
	}
	return jumpFloor1(N-1) + jumpFloor1(N-2)
}

//动态规划
func jumpFloor2(N int) int {
	if N <= 0 {
	  return 0
	}
	if N == 1 || N == 2 {
	  return N
	}
	a, b := 1, 2
	for i := 3; i <= N; i++ {
	  a, b = b, a+b
	}
	return b
}

// 使用dpTable
func jumpFloor3(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 || n == 2 {
		return n
	}

	//存储每一步的结果
	dpTable := make([]int, n+1)

	// base case
	dpTable[1] = 1
	dpTable[2] = 2

	// 状态转移
	for i := 3; i <= n; i++ {
		dpTable[i] = dpTable[i-1] + dpTable[i-2]
	}
	return dpTable[n]
}

//金山云面试题：两个协程交替打印奇数偶数，必须保证按序输出
func TestNum(t *testing.T) {
	fmt.Println("Hello, World!")
	a,b := make(chan int),make(chan int)
	go printNum2(a,b)
	go printNum1(a,b)
	//先起协程，后向channel发数据，否则死锁
	b <- 1
	time.Sleep(2*time.Second)
}

func printNum1(intChan1,intChan2 chan int) {
	//注意此处i必须为偶数，否则不会打印，也不会向channel发送信号
	for i:=2;i<=100;i+=2 {
		if _,ok := <-intChan1; ok {
			if i%2 == 0 {
				fmt.Printf("chan2:%v\n", i)
				intChan2 <- i
			}
		}
	}
}

func printNum2(intChan1,intChan2 chan int) {
	//注意此处i必须为奇数，否则不会打印，也不会向channel发送信号
	for i:=1;i<=100;i+=2 {
		if _,ok := <-intChan2; ok {
			if i%2 != 0 {
				fmt.Printf("chan1:%v\n", i)
				intChan1 <- i
			}
		}
	}
}

//血的教训，可以默写下来才叫掌握了算法思想
func reverseListNode (head *ListNode) *ListNode {
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
	
	chan1,chan2,chan3 := make(chan int),make(chan int),make(chan int)
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

func printNum11(chan1,chan2 chan int) {
	for i:=0;i<2;i++ {
		if _, ok := <- chan1; ok {
		fmt.Println(1)
		wwg.Done()
		chan2 <- 1
		}
	}
}

func printNum12(chan1,chan2 chan int) {
	for i:=0;i<2;i++ {
		if _, ok := <- chan1; ok {
		fmt.Println(2)
		wwg.Done()
		chan2 <- 1
		}
	}
}

func printNum13(chan1,chan2 chan int) {
	for i:=0;i<2;i++ {
		if _, ok := <- chan1; ok {
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
	for i:=0;i<10;i++ {
		ch <- i
		fmt.Printf("produce:%v\n", i)
		time.Sleep(2*time.Second)
	}
}

func consumer(id int, ch chan int, ch1 chan int) {
	for {
		select {
		case v1:=<- ch1:
			fmt.Printf("stop:%v\n", v1)
			break
		case v2:=<- ch:
			fmt.Printf("consumer-%v,data:%v\n",id, v2)
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

	time.Sleep(10*time.Second)

	ch1 <- 1
	ch2 <- 2
}

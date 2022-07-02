package test

import (
	"testing"
	"fmt"
	"math/rand"
	"time"
)

var out []rune
// Perm() 对 a 形成的每⼀排列调⽤ f().
func Perm(a []rune, f func([]rune)) {
    perm(a, f, 0)
}

// 对索引 i 从 0 到 len(a) - 1，实现递归函数 perm().
func perm(a []rune, f func([]rune), i int) {
    // TODO
    if i == len(a) {
        // new := []rune{}
        // new = append(new, out...)
        // f(new)
		f(out)
    }
    for _, v := range a {
        if contains(out, v) {
            continue
        }
        out = append(out, v)
        perm(a, f, i+1)
        out = out[:len(out)-1]
    }
}

func contains(a []rune, x rune) bool {
	for _, v := range a {
		if v == x {
			return true
		}
	}
	return false
}

func TestPerm(t *testing.T) {
    Perm([]rune("ABC"), func(a []rune) {
        fmt.Println(string(a))
    })
}


func rand13() int {
	// rand.Seed(time.Now().UnixNano())
	return rand.Intn(13)+1
}

func rand5() int {
	// rand.Seed(time.Now().UnixNano())
	return rand.Intn(5)+1
}

func rand13ToRand5() int {
	for {
		r := rand13()
		// fmt.Printf("rand13:%v\n", r)
		if r > 5 {
			continue
		}
		return r
	}
}

func rand5ToRand13() int {
	for {
		r := 5 * (rand5()-1) + rand5()
		// fmt.Printf("rand25:%v\n", r)
		if r > 13 {
			continue
		}
		return r
	}
}

func TestRand(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	for i:=0;i<10;i++ {
		r1 := rand13ToRand5()
		t.Log(r1)
	}
	
	for i:=0;i<10;i++ {
		r2 := rand5ToRand13()
		t.Log(r2)
	}
}


// 拼图游戏，不关⼼图⽚内容，只关注碎⽚的边缘形状是否匹配。
// 碎⽚最终需拼成完整的矩形，没有残缺，没有多余碎⽚。
// 碎⽚边缘只有平、凸起、凹陷三种形态，并且不同碎⽚的凸起和凹陷都能完美匹配、⼤⼩⼀致。
// 代码应能检查游戏是否已经完成，⽤碎⽚组成了完整矩形。

//1：平边；2：凸边；3：凹边
type chip struct {
	//碎片上边类型
	up int
	//碎片下边类型
	down int
	//碎片左边类型
	left int
	//碎片右边类型
	right int
}

//12*9矩阵
type picture [9][12]*chip

func (p *picture) put(c *chip, r, l int) {
	p[r][l] = c
}

func (p *picture) delete(r, l int) *chip{
	c := p[r][l]
	p[r][l] = nil
	return c
}

func (p *picture) getLeft(r, l int) *chip {
	if l <= 0 {
		return nil
	}
	return p[l-1][r]
}
func (p *picture) getRight(r, l int) *chip {
	if l >= 11 {
		return nil
	}
	return p[l+1][r]
}
func (p *picture) getUp(r, l int) *chip {
	if r <= 0 {
		return nil
	}
	return p[l][r-1]
}
func (p *picture) getDown(r, l int) *chip {
	if r >= 8 {
		return nil
	}
	return p[l][r+1]
}

func rmChip(p *picture, r,l int) *chip {
	c := p.delete(r, l)
	fmt.Printf("delete chip:%v from picture[%v][%v]\n", c, r, l)
	return c
}

//将碎片放置在12*9的矩阵中，平边必须位于矩阵四边，矩阵内部凸边必须和凹边相邻
func putChip(c *chip, p *picture, r,l int) bool {
	//超出矩阵边界
	if r > 8 || l > 11 || r < 0 || l < 0 {
		return false
	}
	if c == nil {
		return false
	}
	//待放置碎片左边碎片
	left := p.getLeft(r, l)
	//待放置碎片右边碎片
	right := p.getRight(r, l)
	//待放置碎片下边碎片
	down := p.getDown(r, l)
	//待放置碎片上边碎片
	up := p.getUp(r, l)

	//优先处理四个角
	// if c.up == 1 && c.left == 1 {
	// 	if r == 0 && l == 0 {
	// 		p.put(c, r, l)
	// 		return true
	// 	}
	// 	return false
	// }
	// if c.up == 1 && c.right == 1 {
	// 	if r == 0 && l == 11 {
	// 		p.put(c, r, l)
	// 		return true
	// 	}
	// 	return false
	// }
	// if c.down == 1 && c.left == 1 {
	// 	if r == 8 && l == 0 {
	// 		p.put(c, r, l)
	// 		return true
	// 	}
	// 	return false
	// }
	// if c.down == 1 && c.right == 1 {
	// 	if r == 8 && l == 11 {
	// 		p.put(c, r, l)
	// 		return true
	// 	}
	// 	return false
	// }

	// //碎片上边为平边，必须放在第一行
	// if c.up == 1 && r == 0 {
	// 	if judge(c, left, right, nil, down) {
	// 		p.put(c, r, l)
	// 		return true
	// 	}
	// 	return false
	// }
	// //碎片下边为平边，必须放在最后一行
	// if c.down == 1 && r == 8 {
	// 	if judge(c, left, right, up, nil) {
	// 		p.put(c, r, l)
	// 		return true
	// 	}
	// 	return false
	// }
	// //碎片左边为平边，必须放在第一列
	// if c.left == 1 && l == 0 {
	// 	if judge(c, nil, right, up, down) {
	// 		p.put(c, r, l)
	// 		return true
	// 	}
	// 	return false
	// }
	// //碎片右边为平边，必须放在最后一列
	// if c.right == 1 && l == 11 {
	// 	if judge(c, left, nil, up, down) {
	// 		p.put(c, r, l)
	// 		return true
	// 	}
	// 	return false
	// }

	if judge(r, l, c, left, right, up, down) {
		p.put(c, r, l)
		return true
	} else {
		fmt.Printf("put failed,chip:%v,row:%v,col:%v\n", c, r, l)
	}

	return false
}

func judge(row int, col int, c,l,r,u,d *chip) bool {
	if l != nil {
		if l.right == 3 && c.left != 2 {
			return false
		}
		if l.right == 2 && c.left != 3 {
			return false
		}
	}else{
		//判断左边界碎片
		if c.left != 1 && col == 0 {
			return false
		}
		//有平边的碎片不允许放中间
		if col != 0 && c.left == 1 {
			return false
		}
	}
	if r != nil {
		if r.left == 2 && c.right != 3 {
			return false
		}
		if r.left == 3 && c.right != 2 {
			return false
		}
	}else {
		//判断右边界碎片
		if col == 11 && c.right != 1 {
			return false
		}
		if col != 11 && c.right == 1 {
			return false
		}
	}
	if u != nil {
		if u.down == 2 && c.up != 3 {
			return false
		}
		if u.down == 3 && c.up != 2 {
			return false
		}
	}else {
		if row == 0 && c.up != 1 {
			return false
		}
		if row !=0 && c.up == 1 {
			return false
		}
	}
	if d != nil {
		if d.up == 2 && c.down != 3 {
			return false
		}
		if d.up == 3 && c.down != 2 {
			return false
		}
	} else {
		if row == 8 && c.down != 1 {
			return false
		}
		if row != 8 && c.down == 1 {
			return false
		}
	}
	return true
}

func putAndCheck (c *chip, p *picture, r,l int) {
	f := true
	res := putChip(c,p,r,l)
	if res {
		//检查矩阵是否填满，此处可以设置一个全局计数变量，记录矩阵中碎片数量，每次放置完碎片变量自增1，然后判断是否等于9*12
		for i:=0;i<len(*p);i++ {
			for j:=0;j<len((*p)[0]);j++ {
				if (*p)[i][j] == nil {
					f = false
					goto end
				}
			}
		}
	} else {
		f = false
	}
	end:
	if f {
		fmt.Println("恭喜你完成了拼图")
	} else {
		fmt.Println(p)
	}
}

func TestPicture(t *testing.T) {
	var p picture
	//1：平边；2：凸边；3：凹边
	c := &chip{
		up:1,
		down: 2,
		left: 1,
		right: 3,
	}
	putAndCheck(c,&p,0,0)

	c1 := &chip{
		up:1,
		down: 2,
		left: 3,
		right: 3,
	}
	putAndCheck(c1,&p,0,5)

	rmChip(&p, 0, 0)

	putAndCheck(nil, &p, 8, 8)
}
package img

import (
	"image"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
)

// WhiteRect 白色区域的外接矩形信息
type WhiteRect struct {
	X int // 左上角 X 坐标
	Y int // 左上角 Y 坐标
	W int // 区域宽度（像素）
	H int // 区域高度（像素）
}

// unionFind 并查集：用于高效合并等价连通域标签，带路径压缩+按秩合并
type unionFind struct {
	parent []int
	rank   []int
}

func newUnionFind(size int) *unionFind {
	parent := make([]int, size)
	rank := make([]int, size)
	for i := range parent {
		parent[i] = i
		rank[i] = 1
	}
	return &unionFind{parent: parent, rank: rank}
}

// find 查找根节点，路径压缩优化
func (uf *unionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

// union 合并两个集合，按秩合并优化
func (uf *unionFind) union(x, y int) {
	xRoot := uf.find(x)
	yRoot := uf.find(y)
	if xRoot == yRoot {
		return
	}
	if uf.rank[xRoot] < uf.rank[yRoot] {
		uf.parent[xRoot] = yRoot
	} else {
		uf.parent[yRoot] = xRoot
		if uf.rank[xRoot] == uf.rank[yRoot] {
			uf.rank[xRoot]++
		}
	}
}

// regionBounds 记录单个连通域的四边边界
type regionBounds struct {
	minX, minY int
	maxX, maxY int
}

// FindWhiteRectangles 查找图片中宽高均大于 minSize 的白色连通区域
// img: 输入图片
// minSize: 最小宽高阈值（像素），传入 50 即满足宽高>50的要求
// whiteThreshold: 白色阈值，RGB 分量均大于该值视为白色，建议 240~250
func FindWhiteRectangles(img image.Image, minSize int, whiteThreshold uint8) []WhiteRect {
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	totalPixels := width * height

	// 优化点1：统一转为RGBA格式，直接操作像素内存，规避 image.At 的类型断言开销
	rgba := image.NewRGBA(bounds)
	draw.Draw(rgba, bounds, img, bounds.Min, draw.Src)
	pix := rgba.Pix
	stride := rgba.Stride

	// 标签数组：0=背景，>0=连通域ID，一维数组保证内存连续，缓存友好
	labels := make([]int, totalPixels)
	currentLabel := 0
	// 预分配并查集空间，最坏情况每个像素一个标签
	uf := newUnionFind(totalPixels)

	// 优化点2：第一次扫描 - 分配标签 + 记录等价关系，线性时间遍历
	for y := 0; y < height; y++ {
		rowOffset := y * stride
		for x := 0; x < width; x++ {
			// 直接通过字节偏移读取像素，速度远快于 image.At()
			offset := rowOffset + x*4
			r := pix[offset]
			g := pix[offset+1]
			b := pix[offset+2]

			// 非白色像素直接跳过，减少无效计算
			if r <= whiteThreshold || g <= whiteThreshold || b <= whiteThreshold {
				continue
			}

			idx := y*width + x
			leftLabel := 0
			topLabel := 0

			// 4连通判断：只检查左、上两个已处理邻居
			if x > 0 {
				leftLabel = labels[idx-1]
			}
			if y > 0 {
				topLabel = labels[idx-width]
			}

			switch {
			case leftLabel == 0 && topLabel == 0:
				// 无相邻标签，新建连通域
				currentLabel++
				labels[idx] = currentLabel
			case leftLabel != 0 && topLabel == 0:
				labels[idx] = leftLabel
			case leftLabel == 0 && topLabel != 0:
				labels[idx] = topLabel
			default:
				// 两个邻居标签不同，合并等价标签
				labels[idx] = leftLabel
				uf.union(leftLabel, topLabel)
			}
		}
	}

	// 优化点3：第二次扫描 - 统一根标签 + 单遍统计区域边界，减少遍历次数
	regions := make([]regionBounds, currentLabel+1)
	for i := range regions {
		regions[i].minX = width
		regions[i].minY = height
		regions[i].maxX = -1
		regions[i].maxY = -1
	}

	for i := 0; i < totalPixels; i++ {
		label := labels[i]
		if label == 0 {
			continue
		}
		root := uf.find(label)
		x := i % width
		y := i / width

		// 更新连通域边界
		if x < regions[root].minX {
			regions[root].minX = x
		}
		if y < regions[root].minY {
			regions[root].minY = y
		}
		if x > regions[root].maxX {
			regions[root].maxX = x
		}
		if y > regions[root].maxY {
			regions[root].maxY = y
		}
	}

	// 过滤符合宽高条件的区域，生成最终结果
	var result []WhiteRect
	for label := 1; label <= currentLabel; label++ {
		// 只处理根标签，避免重复统计同一连通域
		if uf.find(label) != label {
			continue
		}
		rb := regions[label]
		w := rb.maxX - rb.minX + 1
		h := rb.maxY - rb.minY + 1
		if w > minSize && h > minSize {
			result = append(result, WhiteRect{
				X: rb.minX,
				Y: rb.minY,
				W: w,
				H: h,
			})
		}
	}

	return result
}

// 查找图片中宽高均大于50像素的白色区域，白色阈值240
func FindRect(img image.Image) []WhiteRect {
	// // 1. 打开图片文件
	// file, err := os.Open(imgPath)
	// if err != nil {
	// 	panic(fmt.Sprintf("打开图片失败: %v", err))
	// }
	// defer file.Close()

	// // 2. 解码图片
	// img, _, err := image.Decode(file)
	// if err != nil {
	// 	panic(fmt.Sprintf("解码图片失败: %v", err))
	// }

	// 3. 查找宽高均大于50像素的白色区域，白色阈值240
	rects := FindWhiteRectangles(img, 50, 240)

	// 4. 输出结果
	var reaultArr []WhiteRect = []WhiteRect{}

	// fmt.Printf("共找到 %d 个符合条件的白色区域:\n", len(rects))
	// for i, r := range rects {
	// 	fmt.Printf("区域 %d: 左上角(%d, %d)，宽高(%d, %d)\n",
	// 		i+1, r.X, r.Y, r.W, r.H)
	// 	reaultArr = append(reaultArr, r)
	// }

	for _, r := range rects {
		reaultArr = append(reaultArr, r)
	}

	return reaultArr
}

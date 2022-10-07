package mock_data

import (
	"fmt"
	"math/rand"
	"strings"
)

const familyNames = "赵钱孙李周吴郑王冯陈褚卫蒋沈韩杨朱秦尤许何吕施张"
const firstNames = "明国华建文平志伟东海强晓生光林小民永杰军波成荣新峰刚家龙德庆斌辉良玉俊立浩天宏子金健一忠洪江福祥中正振勇耀春大宁亮宇兴宝少剑云学仁涛瑞飞鹏安亚泽世汉达卫利胜敏群松克清长嘉红山贤阳乐锋智青跃元南武广思雄锦威启昌铭维义宗英凯鸿森超坚旭政传康继翔远力进泉茂毅富博霖顺信凡豪树和恩向道川彬柏磊敬书鸣芳培全炳基冠晖京欣廷哲保秋君劲栋仲权奇礼楠炜友年震鑫雷兵万星骏伦绍麟雨行才希彦兆贵源有景升惠臣慧开章润高佳虎根诚夫声冬奎扬双坤镇楚水铁喜之迪泰方同滨邦先聪朝善非恒晋汝丹为晨乃秀岩辰洋然厚灿卓轩帆若连勋祖锡吉崇钧田石奕发洲彪钢运伯满庭申湘皓承梓雪孟其潮冰怀鲁裕翰征谦航士尧标洁城寿枫革纯风化逸腾岳银鹤琳显焕来心凤睿勤延凌昊西羽百捷定琦圣佩麒虹如靖日咏会久昕黎桂玮燕可越彤雁孝宪萌颖艺夏桐月瑜沛杨钰兰怡灵淇美琪亦晶舒菁真涵爽雅爱依静棋宜男蔚芝菲露娜珊雯淑曼萍珠诗璇琴素梅玲蕾艳紫珍丽仪梦倩伊茜妍碧芬儿岚婷菊妮媛莲娟"

type InfoGenerator struct {
	familyNames []string
	firstNames  []string
	rand        rand.Rand
}

func NewWithSeed(seed int64) InfoGenerator {
	return InfoGenerator{
		familyNames: strings.Split(familyNames, ""),
		firstNames:  strings.Split(firstNames, ""),
		rand:        *rand.New(rand.NewSource(seed)),
	}
}

// first name with length chars
func (gen *InfoGenerator) RandomFirstNames(length int) string {
	s := ""
	var firstNamesLen = len(gen.firstNames)
	for i := 0; i < length; i++ {
		var index = gen.rand.Intn(firstNamesLen)
		s += gen.firstNames[index]
	}
	return s
}

type Names = <-chan [4]string

// Generate 4 names in a family. [father, mother, child1, child2]
func (gen *InfoGenerator) GetFamilyNames() Names {
	var ch = make(chan [4]string)
	var famLen = len(gen.familyNames)
	go func() {
		for i := 0; ; i = (i + 1) % famLen {
			var familyName = gen.familyNames[i]
			var father = familyName + gen.RandomFirstNames(2)
			famIndex := gen.rand.Intn(famLen)
			var mother = gen.familyNames[famIndex] + gen.RandomFirstNames(1)
			var child1 = familyName + gen.RandomFirstNames(1)
			var child2 = familyName + gen.RandomFirstNames(2)
			ch <- [4]string{father, mother, child1, child2}
		}
	}()
	return ch
}

func (gen *InfoGenerator) RandomNameList() <-chan string {
	var ch = make(chan string)
	var famLen = len(gen.familyNames)
	go func() {
		for i := 0; ; i = (i + 1) % famLen {
			var name = gen.familyNames[i] + gen.RandomFirstNames(2)
			ch <- name
		}
	}()
	return ch
}

func (gen *InfoGenerator) RandomPhone() <-chan string {
	var ch = make(chan string)
	go func() {
		for {
			num1, num2 := gen.rand.Intn(9999), gen.rand.Intn(9999)
			ch <- fmt.Sprintf("136%04d%04d", num1, num2)
		}
	}()
	return ch
}

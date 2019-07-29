package model

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/FengGeSe/gotest/util"
)

type CaseSuite struct {
	caseMap   map[string]*CaseFunc   `desc:"case map"`
	caseArray map[string][]*CaseFunc `desc:"case array"`
}

func (s *CaseSuite) AddCaseFuncs(path string, fs []*CaseFunc) {
	for _, f := range fs {
		s.caseMap[f.FuncName] = f
	}
	s.caseArray[path] = fs
}

func (s *CaseSuite) GetFunc(name string) *CaseFunc {
	return s.caseMap[name]
}

func (s *CaseSuite) GetCaseArray() map[string][]*CaseFunc {
	return s.caseArray
}

// 新建一个case的集合
// 遍历根目录root下所有测试方法
func NewCaseSuite(root string) *CaseSuite {
	caseSuite := &CaseSuite{
		caseMap:   map[string]*CaseFunc{},
		caseArray: map[string][]*CaseFunc{},
	}
	testFiles := util.GoTestFiles(root)
	for _, path := range testFiles {
		caseFuncSlc := ScanCaseFromFile(path)
		caseSuite.AddCaseFuncs(path, caseFuncSlc)
	}

	return caseSuite
}

// 测试方法
type CaseFunc struct {
	Path     string `desc:"文件路径"`
	FuncName string `desc:"方法名"`
	Desc     string `desc:"描述"`
}

func (c *CaseFunc) String() string {
	return c.FuncName + "\t" + util.WrapBlue("// "+c.Desc)
}

func ScanCaseFromFile(path string) []*CaseFunc {
	file, err := os.Open(path)
	if err != nil {
		msg := fmt.Sprintf("Error: 打开_test.go文件出错！%v", err)
		fmt.Println(util.WrapRed(msg))
		return []*CaseFunc{}
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	slc := []*CaseFunc{}
	comment := ""
	for scanner.Scan() {
		current := scanner.Text()
		if len(current) == 0 {
			continue
		}
		if strings.HasPrefix(current, "//") {
			comment = current
		}
		isMatch, err := regexp.MatchString(`^func Test[a-zA-Z0-9_]+\(t \*testing\.T\)`, current)
		if err != nil {
			msg := fmt.Sprintf("Error: 正则匹配Test方法出错！", err)
			fmt.Println(util.WrapRed(msg))
			continue
		}
		if isMatch {
			start := strings.Index(current, "Test")
			end := strings.Index(current, "(")
			funcName := current[start:end]
			desc := strings.TrimPrefix(comment, "//")
			temp := &CaseFunc{
				Path:     path,
				FuncName: funcName,
				Desc:     strings.TrimSpace(desc),
			}
			slc = append(slc, temp)
		}
	}
	return slc
}

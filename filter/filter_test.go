package filter

import (
	"testing"
	"time"
)

func BenchmarkFilter(b *testing.B) {
	// 运行该测试前需修改config.go里面的文件夹路径为"../conf/sensitive_words"
	// 执行语句：在filter文件夹下：go test -bench=. -run=none
	begin := time.Now()
	err := Init()
	if err != nil {
		b.Errorf("%v", err)
	}
	b.Logf("%dms", time.Now().Sub(begin).Milliseconds())
	cases := []string{
		"猜十二生肖?(奔逸绝尘打一个生肖)(一句真言:奔逸绝尘)猜什么动物?<p><p>###040+#加·威·∨·信：【 W855～329 】 为您解--答",
		"13mknfokg12ahrsoevw26",
		"谁有毒品卖啊，大约多少钱一克",
		"你是不是练法轮功的",
		"几个点提一下，咨询.1.4.7～0.0.0.6～6.3.7.2。1.【配套齐全】 共享城东大配套，周边4大综合体齐聚:中骏世界城、星光耀广场、华大泰禾广场、润柏 香港城;",
		"☞cherry.然然♛☜ 你被移出了群@☞cherry.然然♛☜ 骗你的😁😁😁😁 ",
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, c := range cases {
			b.Logf("%s", Inst.Filter(c))
		}
	}
}

package inner

import (
	"fmt"
	_ "unsafe"
)

/*
此处的link，是将innerFunc2和OuterFunc2绑定，
其中只要有一个方法定义了实现，
另外的方法只要申明即可。

但是和innerFunc1的略显不同的地方，OuterFunc2如果不加上outer2.s空文件。
会出现：missing function body 的报错。
go build无法编译go:linkname,必须用单独的compile命令进行编译，
因为go build会加上-complete参数，这个参数会检查到没有方法体的方法，并且不通过，报错missing function body。
或者，在这个没有方法体的包下添加一个空的xxx.s 的汇编文件标示，就可以编译了。

此用法在runtime包中经常使用
如：
time.Now()的实现实在runtime包中定义的，使用link将time_now()方法导出使用
//go:linkname time_now time.now
func time_now() (sec int64, nsec int32, mono int64) {
	sec, nsec = walltime()
	return sec, nsec, nanotime()
}

PS:这种实现只是为了更好的理解go源码，平时没必要不要瞎折腾
*/

//go:linkname innerFunc2 golinkname/outer.OuterFunc2
func innerFunc2() {
	fmt.Println("innerFunc2")
}

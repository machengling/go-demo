# Go实用操作demo

## 1.golinkname
- //go:linkname funcname1 path/funcname2
- go源码中经常会看到的骚操作。
- 该操作一般使用在，把内部包的方法，在外部包引用调用。

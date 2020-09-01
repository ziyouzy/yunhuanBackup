package physicalnode
//这里是否需要依赖注入实现Evolve()的包实体对象指针？
//似乎是可以的，因为只需要注入包的指针，不会占用内存
//Evolue包如果和dbreflector，dao，entity一样最为yunhuanfactory独立的功能模块，那么他就需要在main函数中实例化
//然后先依赖注入dao实体，再dao实体内部依赖注入对应的entity实体

//然而
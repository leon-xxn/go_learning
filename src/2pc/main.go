package _pc

//一阶段：资源协调者对每个mysql实例调用prepare命令,让所有mysql实例准备好，如果其中有mysql没有准备好，协调者就让所有实例调用rollback命令回滚。
//如果所有mysql都prepare完成，进入第二阶段

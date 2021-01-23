# 生成器模式（Builder Pattern）

通过一个 Director 指挥者来使用不同的 Builder 创建出对外统一的对象。

> Director 指挥者依赖于抽象 Builder 与 抽象 Result，与工厂方法同理，符合依赖倒置原则

# 动机

一个父类，有多个创建过程相同的子类，子类可能会一直增加，并且使用者只关心生成之后的父类的功能。
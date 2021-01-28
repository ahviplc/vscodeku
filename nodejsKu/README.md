# nodejsKu

> 一些nodejs小练习

```
npm rum app
```

## node中exports和module.exports的关系及使用
```javascript

1.在node中,需要记住,在使用exports和module.exports的时候,实际输出的是module.exports.

2.exports指向module.exports,是module.exports的引用,所以,当使用 exports.a = x 的时候,通过引用关系，造成了module.exports.a = x.当使用 exports = x 的时候,造成了exports不再指向module.exports,所以,仅改变了exports,并没有改变module.exports,也就并没有对输出起作用.

3.当 exports.a = x 和 module.exports.a = xx 一起使用时,最后输出出来的 a = xx；因为一般我们会把module.exports.a == xx放到后面来写，实际上等效于 module.exports.a = x ,然后又执行了module.exports.a = xx,只是进行了一个重新赋值.

4.同样,当 exports.a = x 和 module.exports = xx一起使用时,最后的输出只有 xx,我们再来进行一次转化， exports.a = x 等效于module.exports.a = x, 等效于 module.exports = {a: x}, 然后又执行了module.exports = xx, 实际也是进行了一个重新赋值.

5.如果你觉得，exports与module.exports之间的区别很难分清，一个简单的处理方法，就是放弃使用exports，只使用module.exports
```

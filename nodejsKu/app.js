const path = require('path')
const rf = require('./util/utils')

// 输出模块
console.log('rf -> ', rf) // { AUTHOR: 'LC', readFile: [Function: readFile] }
// 输出作者
console.log('rf.AUTHOR -> ', rf.AUTHOR)

// 使用自定义工具utils
async function app() {
    let res = await rf.readFile(path.join(__dirname, "./data/data.json"))
    console.log(res)
    return 'done'
}

// 执行app
// 函数前面多了一个aync关键字。await关键字只能用在aync定义的函数内。async函数会隐式地返回一个promise，该
// promise的reosolve值就是函数return的值。(示例中reosolve值就是return的字符串'done')
app().then((res) => {
    console.log('app then -> ', res)
}).catch((err) => {
    console.log('app catch -> ', err)
}).finally(() => {
    console.log('app finally -> ', 'Over')
})

const path = require('path')
const rf = require('./util/utils')

// 使用自定义工具utils
async function app() {
    let res = await rf.readFile(path.join(__dirname, "./data/data.json"))
    console.log(res)
}

// 执行
app()

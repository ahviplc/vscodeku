const fs = require('fs')
const path = require('path')

function readFile(path, isSetError) {
    let promise = new Promise((resolve, reject) => {
        fs.readFile(path, 'utf8', function (err, data) {
            if (err || isSetError) {
                if (isSetError) {
                    reject('主动reject...')
                }
                reject(err)
            }
            const resData = JSON.parse(data);

            resolve({
                resData
            })
        })
    });
    return promise;
}

// 方法1
// readFile('./data/data.json')
//     .then((res) => {
//         console.log(res)
//     })
//     .catch((err) => {
//         console.log(err)
//     })
//     .finally(() => {
//         console.log('Over')
//     })

// 方法2
async function run() {
    let res = await readFile('./data/data.json')
    console.log(res)
}

// 方法3 优化文件获取 使用绝对路径
async function run2() {
    let res = await readFile(path.join(__dirname, "./data/data.json"))
    console.log(res)
}

// 执行
run2()

// console.log('我直接过')
//
// console.log(__filename) // C:\_developSoftKu\ideaIU-2019.1.3.win\#CodeKu\vscodeku\nodejsKu\Learn_Promise.js
//
// console.log(__dirname) // C:\_developSoftKu\ideaIU-2019.1.3.win\#CodeKu\vscodeku\nodejsKu (__dirname表示当前文件的目录名)
// console.log(path.dirname(__filename)) // C:\_developSoftKu\ideaIU-2019.1.3.win\#CodeKu\vscodeku\nodejsKu
//
// console.log(path.resolve(__dirname, '..')) // C:\_developSoftKu\ideaIU-2019.1.3.win\#CodeKu\vscodeku
// console.log(path.dirname(__dirname)) // C:\_developSoftKu\ideaIU-2019.1.3.win\#CodeKu\vscodeku
//
// // 这种写法适配各个操作系统
// console.log(path.join(__dirname, "./data/data.json")) // C:\_developSoftKu\ideaIU-2019.1.3.win\#CodeKu\vscodeku\nodejsKu\data\data.json

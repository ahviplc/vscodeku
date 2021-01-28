const fs = require('fs')

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

// 执行
run()

const fs = require('fs')

// 作者
const AUTHOR = 'LC'

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
    console.log('...来自uitl...方法...readFile...')
    return promise;
}

module.exports = {
    AUTHOR,
    readFile
}

// exports指向module.exports,是module.exports的引用.
// 所以，当使用 exports.a = x 的时候，通过引用关系，造成了 module.exports.a = x .
// 也就是说 exports.a = x 和 module.exports.a = x

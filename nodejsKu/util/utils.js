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
    console.log('...来自uitl...方法...readFile...')
    return promise;
}

module.exports = {
    readFile
}

// promptly - npm
// https://www.npmjs.com/package/promptly

const promptly = require('promptly');

// (async () => {
//     const name = await promptly.prompt('Name: ',options.timeout=5000);
//     console.log(name);
// })();

(async () => {
    const answer = await promptly.confirm('Are you really sure? ');

    console.log('Answer:', answer);
})();

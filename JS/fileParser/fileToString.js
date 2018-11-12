var fs = require('fs');

module.exports = function (file) { 
    return fs.readFileSync(`./${file}`, "utf8"); 
}
var fs = require('fs');

module.exports = function (file) { 
    return fs.readFileSync(`./${file}`, "utf8").split("\r\n").map(item => parseInt(item)); 
}
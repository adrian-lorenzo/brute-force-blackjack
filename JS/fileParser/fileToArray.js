const fs = require('fs');

// Returns a file parsed into an Array of ints
module.exports = function (file) { 
    return fs.readFileSync(`${file}`, "utf8").split("\n").map(item => parseInt(item)); 
}
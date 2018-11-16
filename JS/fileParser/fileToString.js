const fs = require('fs');

// Returns a file parsed into a String
module.exports = function (file) { 
    return fs.readFileSync(`./${file}`, "utf8"); 
}
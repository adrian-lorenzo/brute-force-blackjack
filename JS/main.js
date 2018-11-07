const { iteratorClass} = require("./iterator");

function main () {
    const Iterator = iteratorClass();
    Iterator.init(100,95);
    console.time("Loop");
    while (Iterator.hasNext()) {        
        Iterator.next();
    }
    console.timeEnd("Loop");
       
}

main();
const { iteratorClass } = require("./iterator/iterator");

function main () {
    const Iterator = iteratorClass();
    Iterator.init(20,10);
    console.time("Loop");
    while (Iterator.hasNext()) {        
        // console.log(Iterator.next());
        
        Iterator.next();
    }
    console.timeEnd("Loop");
       
}

main();
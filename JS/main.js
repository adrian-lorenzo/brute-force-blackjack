const { iteratorClass} = require("./iterator");

function main () {
    const Iterator = iteratorClass();
    Iterator.init(5,3);
    while (Iterator.hasNext()) {        
        console.log(Iterator.next());
    }
}

main();
const { iteratorClass} = require("./iterator");

function main () {
    const Iterator = iteratorClass();
    Iterator.init(7,3);
    let i = 0;
    const start = new Date().getMilliseconds();
    while (Iterator.hasNext()) {        
        console.log(Iterator.next());
        i++
    }

    console.log(new Date().getMilliseconds()-start);
    
}

main();
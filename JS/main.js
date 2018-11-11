const Iterator = require("./iterator/iterator");
const readFrom = require("./fileParser/fileToArray");
const blackjack = require("./card/blackjack");

function main () {
    const givenCard = parseInt(process.argv[2]);
    const handSize = parseInt(process.argv[3]);

    const deck = blackjack.shuffle(blackjack.draw(readFrom("deck.dat"), givenCard));
    const iterator = new Iterator(deck.length, handSize);

    const { probability, benchmark } = blackjack.getProbabilityToHit(givenCard, iterator, deck);

    console.log(`The probability of hitting blackjack with a ${givenCard} choosing ${handSize} cards is of ${probability}%. TIME: ${benchmark} ms`);
    
}

main();
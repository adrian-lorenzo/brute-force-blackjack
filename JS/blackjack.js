#! /usr/bin/env node

const instructions = require("./fileParser/fileToString")("instructions.txt")
const readCards = require("./fileParser/fileToArray");
const Iterator = require("./iterator/iterator");
const Blackjack = require("./bj/blackjack");

function main () {
    const { givenCard, numberOfDraws } = getParameters();     
    numberOfDraws.forEach((handSize) => {
        const { probability, benchmark } = play(givenCard, handSize);
        console.log(`The probability of hitting blackjack having a ${givenCard} choosing ${handSize} cards is of ${probability}%. TIME: ${benchmark} ms`);
    });  
}

function getParameters () {
    process.argv.shift();
    process.argv.shift();
    
    if (process.argv.includes("-help")) {
        console.log(instructions);
        process.exit();
    }

    if (process.argv.includes("-intense")) {
        process.argv.splice(process.argv.indexOf("-intense"), 1)
        if (process.argv.includes("-nopick")) {
            return { givenCard: 0, numberOfDraws: [...Array(7).keys()].slice(1) }
        }
        return { givenCard: parseInt(process.argv[0]), numberOfDraws: [...Array(7).keys()].slice(1) }
    } 
    if (process.argv.includes("-nopick")) {
        process.argv.splice(process.argv.indexOf("-intense"), 1);
        return { givenCard: 0, numberOfDraws: [parseInt(process.argv[1])] }
    }
    if (process.argv.length === 2) {
        return { givenCard: parseInt(process.argv[0]) , numberOfDraws: [parseInt(process.argv[1])] }
    }
    console.log("ERROR: Bad function invocation, run \"node .\\main.js -help\" for instructions");
    process.exit(1)
    

}

function play(givenCard, handSize) {
    const deck = readCards("../deck.dat");
    const play = new Blackjack (deck, givenCard);
    const iterator = new Iterator(play.deck.length, handSize);
    return play.getProbabilityToHit(iterator);
}

main();
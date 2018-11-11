const BLACKJACK = 21;

function shuffle (deck) {
    for (let i = deck.length - 1; i > 0; i--) {
      const j = Math.floor(Math.random() * (i + 1));
      [deck[i], deck[j]] = [deck[j], deck[i]];
    }
    return deck;
}

function draw(deck, card) {
    deck.splice(deck.indexOf(card), 1);
    return deck;
}


function getHand(combination, deck) {
    return combination.map(item => deck[item]);
}

function isBlackjack (given, hand) {
    return hand.reduce(((value, handCard) => (value + handCard)), given) === BLACKJACK;
}

function getProbabilityToHit(givenCard, iterator, deck) {
    let wins = 0;
    let plays = 0;
    const startTime = parseInt(process.hrtime.bigint())/1e6; 
    while (iterator.hasNext()) {
        const hand = getHand(iterator.next(), deck);
        if (isBlackjack(givenCard, hand)) {
            wins++;
        }
        plays++;
    }
    const endTime = parseInt(process.hrtime.bigint())/1e6;
    
    return { probability: ((wins / plays) * 100).toFixed(2), benchmark: (endTime - startTime).toFixed(2) };
}

module.exports = { shuffle, draw, getHand, isBlackjack, getProbabilityToHit };
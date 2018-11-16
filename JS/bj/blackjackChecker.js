const BLACKJACK = 21;

module.exports = class {

    // Inits the instance of a Blackjack play
    constructor (deck, givenCard) {
        this.deck = deck;
        this.givenCard = givenCard;
        this.reset()
    }

    // Draws from the deck the given card already in hand and shuffles a deck  
    reset() {
        this.draw(this.givenCard)
        this.shuffle();
    }
    
    // Shuffles the deck of the BlackJack Object in a random way
    shuffle () {
        for (let i = this.deck.length - 1; i > 0; i--) {
          const j = Math.floor(Math.random() * (i + 1));
          [this.deck[i], this.deck[j]] = [this.deck[j], this.deck[i]];
        }
    }
    
    // Brute force Algorithm that calculates the probabilities to have equal or 
    // less total hand value (addition of the values of all the cards currently in hand)
    // than 21 (Not Lose a BlackJack Play)
    getProbabilityToHit(iterator) {
        let wins = 0;
        let plays = 0;
        const startTime = parseInt(process.hrtime.bigint())/1e6; 
        while (iterator.hasNext()) {
            const hand =this.getHand(iterator.next());
            if (this.isBlackjack(hand)) {
                wins++;
            }
            plays++;
        }
        const endTime = parseInt(process.hrtime.bigint())/1e6;
        return { probability: ((wins / plays) * 100).toFixed(2), benchmark: (endTime - startTime).toFixed(2) };
    }

    // Removes an element "card" from an array "deck"
    draw(card) { return this.deck.splice(this.deck.indexOf(card), 1); }
    
    // Gets a hand drawn from the deck given a combination 
    getHand(combination) { return combination.map(item => this.deck[item]); }
    
    // Compares if the value of the hand is less or equal to 21
    isBlackjack (hand) { return hand.reduce(((value, handCard) => (value + handCard)), this.givenCard) <= BLACKJACK }   

}



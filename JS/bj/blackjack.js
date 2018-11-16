const BLACKJACK = 21;

module.exports = class {

    constructor (deck, givenCard) {
        this.deck = deck;
        this.givenCard = givenCard;
        this.reset()
    }

    reset() {
        this.draw(this.givenCard)
        this.shuffle();
    }
    
    shuffle () {
        for (let i = this.deck.length - 1; i > 0; i--) {
          const j = Math.floor(Math.random() * (i + 1));
          [this.deck[i], this.deck[j]] = [this.deck[j], this.deck[i]];
        }
    }
    
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

    draw(card) { return this.deck.splice(this.deck.indexOf(card), 1); }
    getHand(combination) { return combination.map(item => this.deck[item]); }
    isBlackjack (hand) { return hand.reduce(((value, handCard) => (value + handCard)), this.givenCard) <= BLACKJACK }   

}



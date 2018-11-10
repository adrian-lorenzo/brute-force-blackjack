# brute-force
### Project for algorithm analysis practice of brute-force strategies in Go, Javascript and Python

CLI program that calculates the probability of not passing if you have one card when taking more in blackjack.

```
.\blackjack <num-first-card> <num-to-take>/-intense

Calculates the probability of not passing if you have one card when taking more in blackjack.

Options:

-help           Prints information about the command.
-intense        Generates the probability of not passing from 1 to 8 cards more taken.
-nopick         The command ignores <num-first-card> argument and calculates the probability
                without taking any card. It works with -intense option.
```


#### Go
Remember to set the project at `$GOPATH/src/github.com/adrianlorenzodev/brute-force/GO/`. You can also get it with the following command `go get github.com:adrianlorenzo/brute-force/GO`.

More info in [How To Write Go Code](https://golang.org/doc/code.html)

We have provided a build script (`builder.sh`) to build the program easily, generating the program binary (`blackjack`).


#### Javascript
We recommend you to use [Node.js](https://nodejs.org/en/) to run the script. Use the command: `node main.js`at the directory of the script.


#### Python
In this case, we also recommend you to use [Python 3](https://www.python.org/download/releases/3.0/) to run the script. At the directory of the script, the command to run it is `python3 main.py`.

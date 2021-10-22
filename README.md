# Data upload scheduler

* [The Problem](#the-problem) 
* [Acceptance criteria](#acceptance-criteria)  
* [Implementation](#implementation) 
* [Running in CLI](#running-in-cli)
* [Running locally](#running-locally) 
* [API Protocol](#API-protocol) 

### The Problem
Faced with falling out of the public consciousness, IBM have decided to re-ignite the
famous Deep Blue vs Gary Kasparov chess matches of the 1990’s. Unfortunately, original 
repository is lost and this repository contains the parts of original repositories logic 
as a rewrite.

### Acceptance Criteria
Given a standard 8x8 chessboard, accept two squares identified by algebraic chess
notation. The first square is the starting position, and the second square is the ending
position. Find the shortest sequence of valid moves to take a Knight piece from the
starting position to the ending position. Each move must be a legal move by a Knight.
For any two squares there may be more than one valid solution. There are no pieces
other than the Knight on the board.

### Implementation
This implementation runs the DFS algorithm as it runs and approaches to closest path each time
with trying the possible moves . After it hits the destination as a result of attempts , it makes
a reverse tree shaking , and removes the unused ways , and failed attempts .

### Running in CLI
Under the /cmd/cli folder , there is a command line app for the path generator . with --parsejson flag
it parses output square notations array into JSON . 

### Running Locally
There is a multi-staged docker file inside of the repository , just build it and run , no external dependencies are needed.

### API Protocol
It contains one simple URL for send a GET request , contains 2 square annotations case sensitively as a source and destination square as a query parameter, and returns the JSON
message contains the array of the squares on the path.

    {{path}}/deep-blue/knight/calculate-route?start_square=a3&dst_square=h1

And returns a JSON (as example):

    {"steps":["a3","c3","d7","e5","d6"]}


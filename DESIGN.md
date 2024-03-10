
1. Tic-Tac-Toe rules should be implemented as described in the [Appendix](#appendix).
2. There can be more than one game occuring at the same time.
3. Joining a game requires a payment of 1 TTT (the native token of the protocol).
4. It is to be assumed that only two players can play a single game. **NOTE**: 
Only players that signalled their intent to join a game must be able to play.
5. Each player should be rewarded 3 points for a win, 1 point for a draw and 
0 points for a loss.
6. A ranking system based on points should be stored. A player which manages 
to score 15 points based on this leaderboard gets 5 newly minted WIN tokens 
and their points tally is set back to 0.

## Tech Requirements

Apart from the game requirements described above, the Tech Lead listed these 
requirements:

1. Unless otherwise stated in the interview the task should be coded using the
[Cosmos-SDK](https://github.com/cosmos/cosmos-sdk). Feel free to use the 
[Ignite CLI](https://github.com/ignite/cli) to scaffold your blockchain.
2. Provide functionality for frontends or third-party programs to query important 
information about the game. At the very least queries should be defined for the 
leaderboard, points tally, and game board.
3. Game board visualisation as the game progresses. You can make use of the functionality
implemented in 2 to display the board on the CLI.
4. Unit testing

## What do we look out for

1. Clean and readable code.
2. Good design choices.
3. Requirements specified above must be adhered to.
4. Assumptions must be reasonable and well explained.
5. Clear inline documentation.
6. Bonus points for features not specified above. Even though not required, we do 
give extra points to candidates that go over and beyond the task requirements.

## Appendix

Rules for Tic-Tac-Toe:

1. The game is played on a grid that's 3 squares by 3 squares.
2. Two players join a game where one player is assigned **X** and the other is assigned **O**. Players take turns putting their marks in empty squares.
3. The first player to get 3 of their marks in a row (up, down, across, or diagonally) is the winner.
4. When all 9 squares are full, the game is over. If no player has 3 marks in a row, the game ends in a tie.

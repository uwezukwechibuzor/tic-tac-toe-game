# Run the chain locally

```
chmod +x start.sh
./start.sh

```




# Transactions commands for the tictactoe module

Usage:

  tic-chaind tx tictactoe [flags]
  
  tic-chaind tx tictactoe [command]

Available Commands:

  create-game --Send a create-game tx
  
  join-game-o --Send a join-game as opponent tx
  
  join-game-x --Send a join-game tx
  
  play-move   --Send a play-move tx
  

# Querying commands for the tictactoe module

Usage:

  tic-chaind query tictactoe [flags]
  
  tic-chaind query tictactoe [command]
  

Available Commands:

  list-game        --List all game
  
  list-game-boards --List all game boards
  
  list-leaderboard --List all leaderboard
  
  show-game        --Shows a game by id
  
  show-game-board  --Shows a game board by id

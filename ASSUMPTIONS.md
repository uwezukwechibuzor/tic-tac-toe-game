# List of Assumptions

1. Game can only be created by an admin set at Genesis using this command

   ```
   tic-chaind add-genesis-game-admin $(tic-chaind keys show admin -a)
   ```

2. Admin can not participate in games

3. The First Player is assigned "X", the first player is the player that joins the game first
4. Opponent is assigned "O"
5. Players make a move using numbers from 1 to 9
6. Board sampe

     1 | 2 | 3
     ---------
     4 | 5 | 6
     ---------
     7 | 8 | 9
    -----------

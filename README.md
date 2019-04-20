
# Valakut Model

A model in Go adapted from the Python model described [here](http://charles.uno/titan-breach-simulation/). Games are played exhaustively, independently tracking all possible sequences of legal plays. For a given deck list, we're able to determine an upper bound for how often it can "go off" with Through the Breach, Primeval Titan, or Scapeshift.

## Shuffling

Suppose we have an Explore in hand and a fetch land in play. Normally, we would have to choose whether to fetch before or after drawing. But the model tries both ways and keeps the better outcome -- essentially double dipping on luck.

The order of the deck does not change over the course of the game. If we would pull a card out of our deck, we instead create a new one out of thin air. This means we neglect the effects of deck thinning. Our estimate is that this is a percent-level effect.

## Lands and Mana

Other than the basic Forests (and maybe a Blighted Woodland), all the lands in the deck can make red mana. And the spells are pretty much all green. In fact, if we have five mana on the table, we're guaranteed to be able to cast Through the Breach. Red mana is never a constraint, so we don't track it. We just track total available mana and available green mana.

This has significant performance implications. We never have to spend cycles figuring out the optimal way to tap for something, like we would if we were modeling something like Bring to Light. We also never have to worry about what to fetch. We always fetch basic Forest. It comes into play untapped, taps for green, and helps with Cinder Glade.

## Game States

A `game_state` object keeps track of about what you would expect: cards in hand, cards on the board, deck order, what turn it is, whether we've played a land, and so on.

At each point, we identify all legal plays, then copy the game state that many times. If you have two lands in play and an Explore in hand, one copy plays it. If you have a Sakura-Tribe Elder in hand too, a copy tries that as well. Another copy passes the turn without playing anything.





## Unique Game States




## Tracking Results





Let's see if we can implement the Valakut model in Go instead of Python


Data files are CSV with fields:

- Turn: what turn did you go off?
- Play: 1 for play, 0 for draw
- Mulls: how many mulligans?
- Fast: 1 for Breach/Shift, 0 for hard-casting Titan

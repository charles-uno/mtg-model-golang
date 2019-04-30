
# Valakut Model

A model in Go adapted from the Python model described [here](http://charles.uno/titan-breach-simulation/). For a given deck list, we're able to determine an upper bound for how often it can "go off" with Through the Breach, Primeval Titan, or Scapeshift.

It's an "upper bound," not exactly an estimate, because the model cheats a little bit. It shuffles up, draws an opening hand, then attempts all possible legal plays and keeps the best result. For example, the model might play Explore, not like what it draws, then go back and play Sakura-Tribe Elder instead.

The effect is small, in general. There aren't that many choices to be made when playing a Valakut deck, so the play patterns that emerge look like what a human would do. But it's important to keep the distinction in mind especially when considering effects that draw cards, like Explore or Sheltered Thicket.

## Usage

To run, use:

```
go run main.go N DECKNAME
```

It'll then load up the deck listed in `lists/DECKNAME.txt` and run `N` games with it. The outcome of each game is saved in `data/DECKNAME.out`. The last run is also printed out verbosely, for example:

```
Deck has 60 cards, 26 lands
On the draw
Shuffling

[T1] Hand: ? Explore Forest SearchforTomorrow ShefetMonitor Shock Valakut, drawing Mountain
Forest
Suspending SearchforTomorrow

[T2] Hand: ? Explore Mountain ShefetMonitor Shock Valakut, drawing ShelteredThicket
[T2] Board: Forest
[T2] Exile: ..SearchforTomorrow, ticking down
Mountain
Explore, drawing Shock

[T3] Hand: ? ShefetMonitor ShelteredThicket Shock*2 Valakut, drawing Pact
[T3] Board: Forest Mountain
[T3] Exile: .SearchforTomorrow, ticking down
SearchforTomorrow from exile
Shock
Cycling ShefetMonitor, drawing Breach

[T4] Hand: ? Breach Pact ShelteredThicket Shock Valakut, drawing Fetch
[T4] Board: Forest*3 Mountain Taiga
Valakut
Breach
```

The program can also be invoked "naked," leaving off all arguments, to see a summary of all existing data:

```
go run main.go
```

## Mulligans

The computer is too good at mulligans. It basically looks at its seven, six, and five in parallel and keeps whichever it likes best. This leads to some non-human play patterns. The logic is in there for Vancouver mulligans, but it's currently turned off.

## Shuffling

Shuffling is a problem, so the model doesn't do it. Playing all possible lines exhaustively means the computer will always find the optimal sequence of plays. But shuffling to blind-draw the optimal card isn't luck or skill -- it's cheating.

The order of the deck does not change over the course of the game. If we would pull a card out of our deck, we instead create a new one out of thin air. This means we neglect the effects of deck thinning. Our estimate is that this is a percent-level effect.

## Lands and Mana

Other than the basic Forests (and maybe a Blighted Woodland), all the lands in the deck can make red mana. And the spells are pretty much all green. In fact, if we have five mana on the table, we're guaranteed to be able to cast Through the Breach. Red mana is never a constraint, so we don't track it. We just track total available mana and available green mana.

This has significant performance implications. We never have to spend cycles figuring out the optimal way to tap for something, like we would if we were modeling something like Bring to Light. We also never have to worry about what to fetch. We always fetch basic Forest. It comes into play untapped, taps for green, and helps with Cinder Glade.

## Game States

A `game_state` object keeps track of about what you would expect: cards in hand, cards on the board, deck order, what turn it is, whether we've played a land, and so on.

At each point, we identify all legal plays, then copy the game state that many times. If you have two lands in play and an Explore in hand, one copy plays it. If you have a Sakura-Tribe Elder in hand too, a copy tries that as well. Another copy passes the turn without playing anything.

Each turn, we clone and iterate each game state until all possible states reach the next turn. If any have found a line to put Primeval Titan on the table, we're done. Otherwise, we play out another turn.

## Unique Game States

A typical game includes thousands of bifurcated game states, many of which involve playing the same cards in a different order. To keep computation time under control, we purge these duplicates. Essentially, each game state packs all its information into a key string, and we eliminate states with duplicate keys.

## Tracking Results

Data files under `data/` are CSV with fields:

- Turn: what turn did you go off?
- Play: 1 for play, 0 for draw
- Mulls: how many mulligans?
- Fast: 1 for Breach/Shift, 0 for hard-casting Titan

After each game completes, a line is added to the file for that deck name. Running with no arguments causes that data to be loaded and summarized. 

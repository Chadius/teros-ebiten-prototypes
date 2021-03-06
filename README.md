# Demos
To run a specific demo, open `main.go` and uncomment the demo you want to run.

Demo descriptions are below:

## mapclicker
Clicking on the screen shows you where you clicked in 3 coordinate systems:
- Screen coordinates (where you clicked)
- World coordinates (the center of the screen is (0,0))
- Map coordinates (grid coordinates)

Screen and World coordinates use the x and y axes, where positive values move towards the right and bottom of the screen.

This is a hex grid, so a neighbor can have up to 6 tiles:
- Left, Right
- Up-left, up-right
- Down-left, down-right

Hex grids have two axes:
- q direction is from left to right. Positive values move right.
- r direction is diagonal, from up-left to down-right. Positive values move down-right.

## mapscroll
Scroll the screen around the map by moving the mouse around the edge.
There are two scroll speeds based on getting close to 10% near the edge and again near 5%.
Also loads all tiles onto the screen.

# Planned demo notes

### 3) UI Widgets

I need to build a copy or something a very **basic** UI. Make a page full of all the widgets I expect to use.

- A button
- A label

Both should be clickable and highlighted.

### 4) Paged UI

Don't bother with layering windows, too complicated!

Just make a UI that can be paged. There should be enough to let the user interact with a map and to display some information.

### 5) Dialogue, full screen

Make a prototype for dialogue:

- A couple of portraits
- And a dialogue window with text
- The two are having a back and forth talk

### 5b) Advanced Dialog

- Scrolling Text - the text is typing out slowly
- Click once to get to the bottom of the message and skip the scrolling.
- Play sound effects while the text is scrolling, to represent them talking.

### 6) Map Loader
Get some kind of input that lets you load the entire map.
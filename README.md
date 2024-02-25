# Picasso
<div align="center">
  <img width="400" alt="image" src="https://github.com/LoyalPotato/picasso/assets/10759876/ab0e976e-f321-409b-bc32-1ab74f3c19ac">
</div>
<br>

Painting app made with go, using [gtk4 bindings](https://github.com/diamondburned/gotk4).

## Current status

At the moment of writing, this is really in an initial state. I'm still learning GTK and all that it encompasses. But the goal is to have a basic painting app.

## MVP

The MVP will require:

- Tools (toolbar)
  - [ ] Eraser
  - [ ] Pen/Pencil
  - [ ] Bucket
  - [ ] Drag/panning
- [ ] Basic Swatch Color Picker
- [ ] Good mouse tracking (in the very first trial, there's some spacing between painting)

## Features that I'm interested in implementing

- Infinite canvas support
- Canvas size configuration
- Line thickness configuration
- Different types of painting tools (airbrush, soft brush, square brush, round brush, etc.)
- Pen pressure support
- Zoom in/out
- Better color picker (full color wheel, hsl, rgb, rgba, etc)

I also would like to be able to use this on the web. At the moment, the only possible way I see it is WASM, but even that I'd still need to confirm.

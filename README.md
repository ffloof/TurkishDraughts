# TurkishDraughts

![](docs/preview.jpg)

## Table of Contents
1. [Introduction](#introduction)
2. [Game Rules](#game-rules)
3. [Move Evaluation](#move-evaluation)
4. [Current Optimizations](#current-optimizations)
5. [Future Roadmap](#future-roadmap)

## Introduction

I wanted to write an ai for a board game but seeing that most like chess or checkers have been done before a million times over, I opted for a variant of checkers/draughts known as [Turkish Draughts](https://en.wikipedia.org/wiki/Turkish_draughts). The rules are fairly similar to traditional checkers/draughts with the main difference being pieces move along the axis as opposed to diagonally. This project includes the ai, and a frontend ui, to play the game against it or watch a match between the two ai.

## Game Rules

The game is similar to checkers except the pieces move along the axis. Non kinged pieces are not allowed to move backwards. Kinged pieces are able to move in any amount of moves in any direction.

![](docs/pawnmoves.jpg) | ![](docs/kingmoves.jpg)
:-: | :-: 

If a piece has the opportunity to take another piece it has to take that piece, if there are multiple combinations of takes you must choose the one that leads to the maximum amount of pieces captured.

//Show image of piece that has to take and post take side by side

Players win when a player has no pieces remaining, no possible moves, or one kinged piece against a non kinged piece

//Show winning boards side by side

## Move Evaluation

//Minmax Explanation

//Tree diagram

//Gif of ai playing against itself

## Current Optimizations

//AB Pruning

//Transposition Table

## Future Roadmap

//Reduce King Move Depth (#16)
//Change Board Value Heuristic To Promote Certain Piece Structures Or Levels Of Agression (#12)
//Various Minor Optimizations (#11)
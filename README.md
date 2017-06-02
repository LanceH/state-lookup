# Overview

This contains two programs.  One to breakdown a shapefile of the US into triangles.  The second will load those triangles and run a server.  It will accept json requests with latitude & longitude and return the state at that location.

# Installation

`go get gitlab.com/LanceH/state-lookup`

# Usage

`statebreakdown <shapefile.shp>`

`statelookupd <statesfile>`

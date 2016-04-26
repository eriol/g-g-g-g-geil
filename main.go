// Copyright © 2016 Daniele Tricoli <eriol@mornie.org>.
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main // import "eriol.xyz/g-g-g-g-geil"

import (
	"fmt"
	"math/rand"
	"time"
)

var phrases = [...]string{
	"On Friday the thirteenth of December",
	"Bruce & Bongo discovered Germany's most successful word: Geil",

	"(G-g-g-g-g)",
	"(G-g-g-g-g)",
	"(G-g-g-g-g)",
	"(G-g-geil, g-g-geil)",

	"The disc-jockey's geil (G-g-g-g-geil)",
	"The disc-jockey's geil (G-g-g-g-geil)",
	"I said the disc-jockey's geil (G-g-g-g-geil)",
	"G-g-geil, g-g-geil",

	"Everybody's geil (G-g-g-g-geil)",
	"Everybody's geil (G-g-g-g-geil)",
	"I said everybody's geil (G-g-g-g-geil)",
	"G-g-geil, g-g-geil",

	"We're so geil",
	"Bruce is geil",

	"Everybody's geil, g-g-g-g-geil",
	"Everybody's geil, g-g-g-g-geil",
	"Everybody's geil",
	"G-g-g-g-geil",

	"Everybody's geil, g-g-g-g-geil",
	"Everybody's geil, g-g-g-g-geil",
	"Everybody's geil",
	"G-g-g-g-geil",

	"Forty, love",
	"Forty, fifteen",
	"Forty, thirty",
	"D-d-d-deuce, d-d-d-deuce",

	"Advantage B-Be",
	"Advantage B-Be",
	"Advantage B-B-B-Becker",
	"Advantage B-B-B-Becker",

	"Advantage B-B-B-Becker",
	"Advantage B-B-B-Becker",

	"B-B-B-Boris is geil (G-g-g-g-geil)",
	"B-B-B-Boris is geil (G-g-g-g-geil)",
	"B-B-B-Boris is geil (G-g-g-g-geil)",
	"B-B-B-Boris, B-B-B-Boris",

	"(G-g-g-g-geil)",
	"Affen sind geil (G-g-g-g-geil)",
	"Affen sind geil (G-g-g-g-geil)",
	"Affen sind geil",

	"Affen affen geil (G-g-g-g-geil)",
	"Affen affen geil (G-g-g-g-geil)",
	"I said Affen affen geil (G-g-g-g-geil)",
	"G-g-geil, g-g-geil",

	"Ich bin geil (G-g-g-g-geil)",
	"Du bist geil (G-g-g-g-geil)",
	"Wir sind geil (G-g-g-g-geil)",
	"G-g-geil, g-g-geil",

	"Everybody's geil (G-g-g-g-geil)",
	"I said everybody's geil (G-g-g-g-geil)",
	"Everybody's geil (G-g-g-g-geil)",
	"G-g-geil, g-g-geil",

	"We're so geil",
	"You're so geil",

	"Everybody's geil, g-g-g-g-geil",
	"Everybody's geil, g-g-g-g-geil",
	"Everybody's geil",
	"G-g-g-g-geil",

	"Everybody's geil, g-g-g-g-geil",
	"Everybody's geil, g-g-g-g-geil",
	"Everybody's geil",
	"G-g-g-g-geil",

	"We're so geil",
	"You're so geil",
	"I'm so geil",
	"Geil",

	"We're so geil",
	"You're so geil",
	"I'm so geil",
	"Geil",

	"I lose control 'cause I'm a creature of the night",
	"Hahahaha",

	"We're so geil",
	"You're so geil",
	"I'm so geil",
	"Geil",

	"Everybody's geil, g-g-g-g-geil",
	"Everybody's geil, g-g-g-g-geil",
	"Everybody's geil",
	"G-g-g-g-geil",

	"Everybody's geil, g-g-g-g-geil",
	"Everybody's geil, g-g-g-g-geil",
	"Everybody's geil",
	"G-g-g-g-geil",

	"I'm so geil",
	"You're so geil",
}

func main() {

	rand.Seed(time.Now().UnixNano())

	fmt.Println(phrases[rand.Intn(len(phrases))])

}

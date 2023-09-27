package pathing_test

import (
	"fmt"
	"testing"

	"github.com/quasilyte/pathing"
)

func BenchmarkGreedyBFS(b *testing.B) {
	l := pathing.MakeGridLayer([4]uint8{1, 0, 1, 1})
	for i := range bfsTests {
		test := bfsTests[i]
		if !test.bench {
			continue
		}
		numCols := len(test.path[0])
		numRows := len(test.path)
		b.Run(fmt.Sprintf("%s_%dx%d", test.name, numCols, numRows), func(b *testing.B) {
			parseResult := testParseGrid(b, test.path)
			bfs := pathing.NewGreedyBFS(pathing.GreedyBFSConfig{
				NumCols: uint(parseResult.numCols),
				NumRows: uint(parseResult.numRows),
			})
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				bfs.BuildPath(parseResult.grid, parseResult.start, parseResult.dest, parseResult.grid.Cost(l))
			}
		})
	}
}

func TestGreedyBFS(t *testing.T) {
	for i := range bfsTests {
		runPathfindTest(t, bfsTests[i], func(cols, rows uint) pathBuilder {
			return pathing.NewGreedyBFS(pathing.GreedyBFSConfig{
				NumCols: cols,
				NumRows: rows,
			})
		})
	}
}

var bfsTests = []pathfindTestCase{
	{
		name: "trivial_short",
		path: []string{
			"..........",
			"...A   $..",
			"..........",
		},
		bench: true,
	},

	{
		name: "trivial_short2",
		path: []string{
			"..........",
			"...A......",
			"... ......",
			"... ......",
			"...  $....",
			"..........",
		},
		bench: true,
	},

	{
		name: "trivial",
		path: []string{
			".A..........",
			". ..........",
			". ..........",
			". ..........",
			". ..........",
			". ..........",
			".          $",
		},
		bench: true,
	},

	{
		name: "trivial_long",
		path: []string{
			".......................x........",
			"                               $",
			"A...............................",
			"..........................x.....",
		},
		bench: true,
	},

	{
		name: "simple_wall1",
		path: []string{
			"........",
			"...A....",
			"...   ..",
			"....x ..",
			"....x $.",
		},
		bench: true,
	},

	{
		name: "simple_wall2",
		path: []string{
			"...   ..",
			"...Ax ..",
			"....x ..",
			"....x ..",
			"....x $.",
		},
		bench: true,
	},

	{
		name: "simple_wall3",
		path: []string{
			"..........x.....................",
			"..........x.....................",
			"..........x.....................",
			"..........x.....................",
			".............   ................",
			"..            x          $......",
			".. ...........x.................",
			"..A...........x.................",
			"....x...........................",
			"....x...........................",
			"....x...........................",
			"....x...........................",
		},
		bench: true,
	},

	{
		name: "simple_wall4",
		path: []string{
			"..........x.....................",
			"..........x.....................",
			"..........x.....................",
			"..........x.....................",
			"................................",
			"..............x.................",
			"..............x.................",
			"..A...........x.................",
			".. .x...........................",
			".. .x...........................",
			".. .x...........................",
			".. .x...........................",
			".. .............................",
			".. .............................",
			".. ..................xxxxxxxx...",
			".. .............................",
			".. .............................",
			".. ...........x.................",
			".. ...........x.................",
			"..    ........x.................",
			"....x ..........................",
			"....x                      $....",
			"....x...........................",
			"....x...........................",
		},
		bench: true,
	},

	{
		name: "zigzag1",
		path: []string{
			"........",
			"   A....",
			" xxxxxx.",
			" .......",
			" .xxxxxx",
			" .......",
			" $......",
		},
		bench: true,
	},

	{
		name: "zigzag2",
		path: []string{
			"........",
			"...A    ",
			".xxxxxx ",
			".....   ",
			"..xxx xx",
			"..... ..",
			".....  $",
		},
		bench: true,
	},

	{
		name: "zigzag3",
		path: []string{
			"...   ....x.....",
			"..A x ....x.....",
			"....x ....x.....",
			"....x ....x.....",
			"....x        $..",
			"....x...........",
		},
		bench: true,
	},

	{
		name: "zigzag4",
		path: []string{
			"...   .x.   x...",
			"... x .x. x x...",
			"... x .x. x x...",
			"... x .x. x   ..",
			"..A x  x  x.x  $",
			"....x.   .x.x...",
		},
		bench: true,
	},

	{
		name: "zigzag5",
		path: []string{
			".A     ..",
			"xxxxxx ..",
			"..     ..",
			".. xxxxxx",
			"..   ....",
			"xxxx x...",
			"....    .",
			"...xxxx x",
			".......$.",
		},
		bench: true,
	},

	{
		name: "double_corner1",
		path: []string{
			".   .x  A.",
			". x .x ...",
			"x x .x ...",
			"  x .x ...",
			" xx    ...",
			" .xxxxxxxx",
			"   $......",
		},
		bench: true,
	},

	{
		name: "double_corner2",
		path: []string{
			".    x..A.",
			". x. x.. .",
			"x x. x.. .",
			"  x. x.. .",
			" xx.     .",
			" .xxxxxxxx",
			"        $.",
			"..........",
		},
	},

	{
		name: "double_corner3",
		path: []string{
			"    x..A.",
			" x. x.. .",
			" x. x.. .",
			" x.     .",
			" xxxxxxxx",
			"       $.",
		},
	},

	{
		name: "labyrinth1",
		path: []string{
			".........x.....",
			"xxxxxxxx.x.  $.",
			"x.     x.x. ...",
			"x. xxx x.x. ...",
			"x.   x x.x. ...",
			"x...Ax   xx .xx",
			"x....x.x x  ...",
			"xxxxxx.x x xxxx",
			"x......x x    .",
			"xxxxxxxx xxxx x",
			"........ x    .",
			"........   ....",
		},
		bench: true,
	},

	{
		name: "labyrinth2",
		path: []string{
			".x......x.......x............",
			".x......x.......x............",
			".x......x.......x............",
			".x......x.......xxxxxxxxxx...",
			".x....       ...x.....    ...",
			".x     xxx.x    x.....$.x  xx",
			"   .x..x...xxx. x.......x.  .",
			"A...x..x...x... xxxxxxxxxxx .",
			"..x.x..x.......     x       .",
			"..x.x..x....x...... x .......",
			"..x.x..x..xxxx...x.   .......",
			"..x.x.......x....x...........",
		},
		bench: true,
	},

	{
		name: "labyrinth3",
		path: []string{
			"...x......x........x............",
			"..Ax......x........x............",
			".. x......x........xxxxxxxxxx...",
			".. x...............x............",
			".. x.....xxx..x....x.......x..xx",
			".. ...x..x....xxx..x.......x....",
			".. ...x..x....x....xxxxxxxxxxx..",
			".. .x.x..x.....x...   .x........",
			".. .x.x..x...xxxx.  x         ..",
			"..        x....... xxxxxxxxxx ..",
			"xxxx.....       .. x........  ..",
			"...x.....xxx..x    x.......x .xx",
			"......x..x....xxx..x.......x   .",
			"......x..x....x....xxxxxx.xxxx .",
			"....x.x........x....x..x...$   .",
		},
		bench: true,
	},

	{
		// This is unfortunate.
		// TODO: can we adjust anything to make it better?
		name: "depth1",
		path: []string{
			"........................",
			".xxxxxxxxxxxxxxxxxxxx...",
			"....................x...",
			".xxxxxxxxxxxxxxxxxx.x...",
			"....................x...",
			".x.xxxxxxxxxxxxxxxxxx...",
			"..................A x.B.",
			".x.xxxxxxxxxxxxxxxxxx...",
			"....................x...",
			".xxxxxxxxxxxxxxxxxx.x...",
			"....................x...",
			".xxxxxxxxxxxxxxxxxxxx...",
			"........................",
		},
		partial: true,
		bench:   true,
	},

	{
		name: "depth2",
		path: []string{
			"...................   ..",
			"..                  x ..",
			".x xxxxxxxxxxxxxxxxxx ..",
			"..                A.x $.",
			".x.xxxxxxxxxxxxxxxxxx...",
			"....................x...",
			".xxxxxxxxxxxxxxxxxx.x...",
			"....................x...",
			".xxxxxxxxxxxxxxxxxxxx...",
			"........................",
		},
		bench: true,
	},

	{
		name: "nopath1",
		path: []string{
			"A    x.B",
			".....x..",
		},
		partial: true,
		bench:   true,
	},

	{
		name: "nopath2",
		path: []string{
			"....Ax.B",
			".....x..",
		},
		partial: true,
		bench:   true,
	},

	{
		name: "nopath3",
		path: []string{
			"........",
			".xxxxx..",
			".x...x..",
			".x.A.x..",
			".x.  x..",
			".xxxxx..",
			".......B",
		},
		partial: true,
		bench:   true,
	},

	{
		name: "nopath4",
		path: []string{
			".....x.....x..",
			".xxxxx.   .x..",
			".x...x. x .x..",
			".x.A.x. x  x..",
			".x.     xxxx..",
			".xxxxxxxx.....",
			".............B",
		},
		partial: true,
		bench:   true,
	},

	{
		name: "nopath5",
		path: []string{
			".B...x.....x..",
			".xxxxx.....x..",
			".x  .x..x..x..",
			".x.A.x..x..x..",
			".x......xxxx..",
			".xxxxxxxx.....",
			"..............",
		},
		partial: true,
		bench:   true,
	},

	{
		name: "tricky1",
		path: []string{
			"               $",
			" .xxxxxxxxxxxx..",
			" ............x..",
			" ............x..",
			" ............x..",
			" ............x..",
			" ............x..",
			"A..xxxxxxxxxxx..",
			"................",
		},
		bench: true,
	},
	{
		name: "tricky2",
		path: []string{
			"...............",
			".             .",
			"  xxxxxxxxxxx $",
			" ...........x..",
			" ...........x..",
			" ...........x..",
			" ...........x..",
			" ...........x..",
			" ...........x..",
			" ...........x..",
			" ...........x..",
			" ...........x..",
			"A.xxxxxxxxxxx..",
			"...............",
			"...............",
		},
		bench: true,
	},

	{
		name: "tricky3",
		path: []string{
			"...............",
			"...............",
			"..xxxxxxxxxxx A",
			"............x .",
			"............x .",
			"............x .",
			"............x .",
			"............x .",
			"............x .",
			"............x .",
			"............x .",
			"............x .",
			"$ xxxxxxxxxxx .",
			".             .",
			"...............",
		},
		bench: true,
	},

	{
		name: "tricky4",
		path: []string{
			"...............",
			".             .",
			". xxxxxxxxxxx $",
			".     ......x..",
			"..... ......x..",
			"..... ......x..",
			"..... ......x..",
			"..... ......x..",
			"..... ......x..",
			"..... ......x..",
			"..... ......x..",
			".....A......x..",
			"..xxxxxxxxxxx..",
			"...............",
			"...............",
		},
		bench: true,
	},

	{
		name: "tricky5",
		path: []string{
			"...............",
			"...............",
			"A.xxxxxxxxxxx..",
			" ...........x..",
			" ...........x..",
			" ...........x..",
			" ...........x..",
			" ...........x..",
			" ...........x..",
			" ...........x..",
			" ...........x..",
			" ...........x..",
			"  xxxxxxxxxxx $",
			".             .",
			"...............",
		},
	},

	{
		name: "tricky6",
		path: []string{
			"............$ .",
			"............. .",
			"..xxxxxxxxxxx .",
			"..x.........x .",
			"..x.........x .",
			"..x.........x .",
			"..x.........x .",
			"..x.........x .",
			"..x.........x .",
			"..x.........x .",
			"..x.........x .",
			"..x.........x .",
			"..x.........x .",
			"..            .",
			"..A............",
		},
	},

	{
		name: "tricky7",
		path: []string{
			"..          A..",
			".  ............",
			". xxxxxxxxxxx..",
			". x.........x..",
			". x.........x..",
			". x.........x..",
			". x.........x..",
			". x.........x..",
			". x.........x..",
			". x.........x..",
			". x.........x..",
			". x.........x..",
			". x.........x..",
			". .............",
			". $............",
		},
	},

	{
		name: "tricky8",
		path: []string{
			". $............",
			". .............",
			". xxxxxxxxxxx..",
			". x.........x..",
			". x.........x..",
			". x.........x..",
			". x.........x..",
			". x.........x..",
			". x.........x..",
			". x.........x..",
			". x.........x..",
			". x.........x..",
			". x.........x..",
			".            ..",
			"............A..",
		},
	},

	{
		name: "tricky9",
		path: []string{
			". $............",
			". .............",
			". xxxxxxxxxxx..",
			". x.........x..",
			". x.........x..",
			". x.........x..",
			". x.........x..",
			". x         x..",
			". x .......Ax..",
			". x ........x..",
			". x ........x..",
			". x ........x..",
			". x ........x..",
			".   ...........",
			"...............",
		},
	},

	{
		name: "tricky10",
		path: []string{
			". $............",
			". .............",
			". xxxxxxxxxxx..",
			". x.........x..",
			". x.........x..",
			". x.........x..",
			". x.........x..",
			". x.........x..",
			". xA........x..",
			". x ........x..",
			". x ........x..",
			". x ........x..",
			". x ........x..",
			".   ...........",
			"...............",
		},
	},

	{
		name: "tricky11",
		path: []string{
			".    $.........",
			". .............",
			". xxxxxxxxxxx..",
			". x.........x..",
			". x.........x..",
			". x.........x..",
			". x.........x..",
			". x.........x..",
			". x.........x..",
			". x.........x..",
			". x.........x..",
			". x.        x..",
			". x  ...... x..",
			".   .......  ..",
			"............A..",
		},
	},

	{
		name: "tricky12",
		path: []string{
			"..........$   .",
			"............. .",
			"..xxxxxxxxxxx .",
			"..x.........x .",
			"..x.........x .",
			"..x.........x .",
			"..x.........x .",
			"..x.........x .",
			"..x.........x .",
			"..x.........x .",
			"..x.........x .",
			"..x.........x .",
			"..x.........x .",
			"............  .",
			"............A..",
		},
	},

	{
		name: "tricky13",
		path: []string{
			"...............",
			"           $...",
			" .....xxxxxxx..",
			" ...........x..",
			" ...........x..",
			" ...........x..",
			" ...........x..",
			" ...........x..",
			" ...........x..",
			" ...........x..",
			" ...........x..",
			" ...........x..",
			"A.xxxxxxxxxxx..",
			"...............",
			"...............",
		},
	},

	{
		name: "distlimit1",
		path: []string{
			"A                                                        ..........B",
		},
		bench:   true,
		partial: true,
	},

	{
		name: "distlimit2",
		path: []string{
			"A.............x......   ....            ......x.....x.....x....",
			" .............x...... x      xxxxxxxxxx ......x..x..x..x..x....",
			" ...xxxxxxxxxxx...... x...............x ......x..x..x..x..x....",
			"                      x...............x       ...x.....x......B",
		},
		bench:   true,
		partial: true,
	},
}

package login

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

func showLogo(v *gocui.View) {
	fmt.Fprintln(v, "                       ....                                          ")
	fmt.Fprintln(v, "                       ....                                          ")
	fmt.Fprintln(v, "                       ....                                          ")
	fmt.Fprintln(v, "                    ....... ....                                     ")
	fmt.Fprintln(v, "                    ..............                                   ")
	fmt.Fprintln(v, "                    ....... ...  ..      ..                          ")
	fmt.Fprintln(v, "                    .....  ..    .    .....                ..        ")
	fmt.Fprintln(v, "                        ... .. .....   .. ...   .    ..    ... .. .  ")
	fmt.Fprintln(v, "                        .... .......  ..      ..... .... ..... .. .  ")
	fmt.Fprintln(v, "                        ..... ..   .  ...  .. .. .. .... .. .. .. .  ")
	fmt.Fprintln(v, "                        .    . ..  .   ...... ..... .... ..... ....  ")
	fmt.Fprintln(v, "                        ..  .. .. .    .....  ..... ....  .... ....  ")
	fmt.Fprintln(v, "                         ..... ........                          ..  ")
	fmt.Fprintln(v, "                          ...  ........                         ...  ")
	fmt.Fprintln(v, "                            ............                             ")
	fmt.Fprintln(v, "                                .......                              ")
	fmt.Fprintln(v, "                                ....                                 ")
	fmt.Fprintln(v, "                                ....                                 ")
	fmt.Fprintln(v, "                                 ...                                 ")

}

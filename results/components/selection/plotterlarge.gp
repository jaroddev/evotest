set title 'Comparaison des selections avec une large population 200N'
set xlabel 'Generation'
set ylabel 'Best fitness'

set grid
set key center right vertical noreverse enhanced autotitle box dashtype solid
set tics out nomirror
set border 3 front linetype black linewidth 3.0 dashtype solid

set xrange [1:4000]

set terminal png enhanced

set datafile separator ","

set output 'Comparaisons200.png'

plot '200-best-three-clone-elitist.csv' using 1:2 with lines linewidth 2 title 'best', \
     '200-tournament-three-clone-elitist.csv' using 1:2 with lines linewidth 2 title 'tournament', \
     '200-random-three-clone-elitist.csv' using 1:2 with lines linewidth 2 title 'random',
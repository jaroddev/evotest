set title 'Comparaison des croisements'
set xlabel 'Generation'
set ylabel 'Best fitness'

set grid
set key center right vertical noreverse enhanced autotitle box dashtype solid
set tics out nomirror
set border 3 front linetype black linewidth 3.0 dashtype solid

set xrange [1:4000]
set style line 1 linecolor rgb '#0060ad' linetype 1 linewidth 1

set terminal png enhanced

set datafile separator ","

set output 'Comparaisons.png'

plot 'best-three-mono-elitist.csv' using 1:2 with lines title 'mono', \
    'best-three-uniform-elitist.csv' using 1:2 with lines title 'uniform'
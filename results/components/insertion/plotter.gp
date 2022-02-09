set title 'Comparaison des insertions'
set xlabel 'Generation'
set ylabel 'Best fitness'

set grid
set key center right vertical noreverse enhanced autotitle box dashtype solid
set tics out nomirror
set border 3 front linetype black linewidth 1.0 dashtype solid

set xrange [1:4000]
set style line 1 linecolor rgb '#0060ad' linetype 1 linewidth 1

set terminal png enhanced

set datafile separator ","

set output 'Comparaisons.png'

plot 'best-three-clone-elitist.csv' using 1:2 with lines linewidth 4 title 'elitist', \
    'best-three-clone-age.csv' using 1:2 with lines linewidth 2 title 'age'
set title 'Comparaison des configs selon la taille de population'
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

plot '50-cp1-mp1.csv' using 1:2 with lines linewidth 3 title 'N=50', \
     '200-cp1-mp1.csv' using 1:2 with lines linewidth 3 title 'N=200', \
     '500-cp1-mp1.csv' using 1:2 with lines linewidth 3 title 'N=500', \
     '1000-cp1-mp1.csv' using 1:2 with lines linewidth 3 title 'N=1000'
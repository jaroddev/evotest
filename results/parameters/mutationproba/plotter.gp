set title 'Comparaison des probas de mutations'
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

plot '200-cp1-mpquarter.csv' using 1:2 with lines linewidth 3 title 'mutate with 1/4 chances', \
     '200-cp1-mphalf.csv' using 1:2 with lines linewidth 3 title 'mutate half time', \
     '200-cp1-mp3quarter.csv' using 1:2 with lines linewidth 3 title 'often mutate (3/4)', \
     '200-cp1-mp1.csv' using 1:2 with lines linewidth 3 title 'always mutate'
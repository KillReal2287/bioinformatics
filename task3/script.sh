rm -f ./ali.sam
minimap2 -a -t 6 ./et_ecoli.fna ./SRR24487960.fasta > ali.sam
QUALITY=$(samtools flagstat ali.sam | python3 -c 'from sys import stdin; d=stdin.read(); import re; print(re.findall(r"\d+\s+\+\s+\d+\s+mapped\s+\((\d+\.\d+)%", d)[0])')
COMPARATION=$(python3 -c "print(int(float($QUALITY) > 90))")
echo $COMPARATION

if [ "$COMPARATION" -eq "1" ]; then
    echo "QUALITY is ok ($COMPARATION > 90%)"
else
    echo "QUALITY is not ok ($COMPARATION < 90%)"
    exit 1
fi

rm -f ./ali.bam
samtools view -S -b ./ali.sam > ./ali.bam

rm -f ./ali.sorted.bam
samtools sort ali.bam  -o ali.sorted.bam

freebayes -f ./et_ecoli.fna ./ali.sorted.bam > var.vcf
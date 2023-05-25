package main

import (
	sp "github.com/scipipe/scipipe"
	spc "github.com/scipipe/scipipe/components"
)

func main() {
	wf := sp.NewWorkflow("workflow", 4)

	src_ref := spc.NewFileSource(wf, "in_ref", "input/ref_e_coli.fna.gz")
	src_source := spc.NewFileSource(wf, "in_source", "input/e_coli_fastq.gz")
	src_parser := spc.NewFileSource(wf, "in_parser", "input/parser.py")

	index := wf.NewProc("index", "minimap2 -d {o:ref} {i:in_ref}")
	index.In("in_ref").From(src_ref.Out())
	index.SetOut("ref", "output/ref.mmi")

	fastqc := wf.NewProc("fastqc", "fastqc {i:in_source}")
	fastqc.In("in_source").From(src_source.Out())

	alignment := wf.NewProc("alignment", "minimap2 -a {i:ref} {i:in_source} > {o:ares}")
	alignment.In("ref").From(index.Out("ref"))
	alignment.In("in_source").From(src_source.Out())
	alignment.SetOut("ares", "output/alignment.sam")

	samtools := wf.NewProc("samtools", "samtools flagstat {i:ares} > {o:sres}")
	samtools.In("ares").From(alignment.Out("ares"))
	samtools.SetOut("sres", "output/result.txt")

	parse := wf.NewProc("parse", "python3 {i:in_parser} {i:sres} > {o:parsed}")
	parse.In("in_parser").From(src_parser.Out())
	parse.In("sres").From(samtools.Out("sres"))
	parse.SetOut("parsed", "output/parsed.txt")

	view := wf.NewProc("view", "echo {i:parsed} > {o:out}")
	view.In("parsed").From(parse.Out("parsed"))

	wf.PlotGraph("my_workflow_graph.dot")
	wf.PlotGraphPDF("my_workflow_graph.dot")

	// Run workflow
	wf.Run()
}

package main

import "github.com/spf13/cobra"

func commands() *cobra.Command {
	squash := &cobra.Command{
		Use:   `squash`,
		Short: `Squashing revisions`,
		Long: `squash: usage: repocutter [-q] [-r SELECTION] [-m mapfile] [-f] [-c] squash

The 'squash' subcommand merges adjacent commits that have the same
author and log text and were made within 5 minutes of each other.
This can be helpful in cleaning up after migrations from file-oriented
revision control systems, or if a developer has been using a pre-2006
version of Emacs VC.

With the -m (or --mapfile) option, squash emits a map to the named
file showing how old revision numbers map into new ones.

With the -e (or --excise) option, the specified set of revisions in
unconditionally removed.  The tool will exit with an error if an
excised remove is part of a clique eligible for squashing.  Note that
repocutter does not perform any checks on whether the repository
history is afterwards valid; if you delete a node using this option,
you won't find out you have a problem until you attempt to load the
resulting dumpfile.

repocutter attempts to fix up references to Subversion revisions in log
entries so they will still be correct after squashing.  It considers
anything that looks like the regular expression \br[0-9]+\b to be
a comment reference (this is the same format that Subversion uses
in log headers).

Every revision in the file after the first omitted one gets the property
'repocutter:original' set to the revision number it had before the
squash operation.

The option --f (or --flagrefs) causes repocutter to wrap its revision-reference
substitutions in curly braces ({}).  By doing this, then grepping for 'r{'
in the output of 'repocutter log', you can check for false conversions.

The -c (or --compressmap) option changes the mapfile format to one
that is easier for human browsing, though less well suited for
interpretation by other programs.`,
		Run: func(cmd *cobra.Command, args []string) {},
	}

	select_ := &cobra.Command{
		Use:   `select`,
		Short: `Selecting revisions`,
		Long: `select: usage: repocutter [-q] [-r SELECTION] select

The 'select' subcommand selects a range and permits only revisions in
that range to pass to standard output.  A range beginning with 0
includes the dumpfile header.`,
		Run: func(cmd *cobra.Command, args []string) {},
	}

	propdel := &cobra.Command{
		Use:   `propdel PROPNAME`,
		Short: `Deleting revision properties`,
		Long: `propdel: usage: repocutter [-r SELECTION] propdel PROPNAME...

Delete the property PROPNAME. May be restricted by a revision
selection. You may specify multiple properties to be deleted.`,
		Run: func(cmd *cobra.Command, args []string) {},
	}

	propset := &cobra.Command{
		Use:   `propset PROPNAME=PROPVAL`,
		Short: `Setting revision properties`,
		Long: `propset: usage: repocutter [-r SELECTION] propset PROPNAME=PROPVAL...

Set the property PROPNAME to PROPVAL. May be restricted by a revision
selection. You may specify multiple property settings.`,
		Run: func(cmd *cobra.Command, args []string) {},
	}

	proprename := &cobra.Command{
		Use:   `proprename OLDNAME->NEWNAME`,
		Short: `Renaming revision properties`,
		Long: `proprename: usage: repocutter [-r SELECTION] proprename OLDNAME->NEWNAME...

Rename the property OLDNAME to NEWNAME. May be restricted by a
revision selection. You may specify multiple properties to be renamed.`,
		Run: func(cmd *cobra.Command, args []string) {},
	}

	log := &cobra.Command{
		Use:   `log`,
		Short: `Extracting log entries`,
		Long: `log: usage: repocutter [-r SELECTION] log

Generate a log report, same format as the output of svn log on a
repository, to standard output.`,
		Run: func(cmd *cobra.Command, args []string) {},
	}

	setlog := &cobra.Command{
		Use:   `setlog`,
		Short: `Mutating log entries`,
		Long: `setlog: usage: repocutter [-r SELECTION] --logentries=LOGFILE setlog

Replace the log entries in the input dumpfile with the corresponding entries
in the LOGFILE, which should be in the format of an svn log output.
Replacements may be restricted to a specified range.`,
		Run: func(cmd *cobra.Command, args []string) {},
	}

	strip := &cobra.Command{
		Use:   `strip PATTERN`,
		Short: `Replace content with unique cookies, preserving structure`,
		Long: `strip: usage: repocutter [-r SELECTION] strip PATTERN...

Replace content with unique generated cookies on all node paths
matching the specified regular expressions; if no expressions are
given, match all paths.  Useful when you need to examine a
particularly complex node structure.`,
		Run: func(cmd *cobra.Command, args []string) {},
	}

	expunge := &cobra.Command{
		Use:   `expunge PATTERN`,
		Short: `Expunge operations by Node-path header`,
		Long: `expunge: usage: repocutter [-r SELECTION ] expunge PATTERN...

Delete all operations with Node-path headers matching specified
Python regular expressions (opposite of 'sift').  Any revision
left with no Node records after this filtering has its Revision
record removed as well.`,
		Run: func(cmd *cobra.Command, args []string) {},
	}

	sift := &cobra.Command{
		Use:   `sift PATTERN`,
		Short: `Sift for operations by Node-path header`,
		Long: `sift: usage: repocutter [-r SELECTION ] sift PATTERN...

Delete all operations with Node-path headers *not* matching specified
Python regular expressions (opposite of 'expunge').  Any revision left
with no Node records after this filtering has its Revision record
removed as well.`,
		Run: func(cmd *cobra.Command, args []string) {},
	}

	pathrename := &cobra.Command{
		Use:   `pathrename FROM TO`,
		Short: `Transform path headers with a regexp replace`,
		Long: `pathrename: usage: repocutter [-r SELECTION ] pathrename FROM TO

Modify Node-path headers, Node-copyfrom-path headers, and
svn:mergeinfo properies matching the specified Python regular
expression FROM; replace with TO.  TO may contain Pyton backreferences
to parenthesized portions of FROM.`,
		Run: func(cmd *cobra.Command, args []string) {},
	}

	renumber := &cobra.Command{
		Use:   `renumber`,
		Short: `Renumber revisions so they're contiguous`,
		Long: `renumber: usage: repocutter renumber

Renumber all revisions, patching Node-copyfrom headers as required.
Any selection option is ignored. Takes no arguments.`,
		Run: func(cmd *cobra.Command, args []string) {},
	}

	reduceCommand := &cobra.Command{
		Use:   `reduce INPUT-FILE`,
		Short: `Topologically reduce a dump.`,
		Long: `reduce: usage: repocutter reduce INPUT-FILE

Strip revisions out of a dump so the only parts left those likely to
be relevant to a conversion problem. A revision is interesting if it
either (a) contains any operation that is not a plain file
modification - any directory operation, or any add, or any delete, or
any copy, or any operation on properties - or (b) it is referenced by
a later copy operation. Any commit that is neither interesting nor
has interesting neighbors is dropped.

Because the 'interesting' status of a commit is not known for sure
until all future commits have been checked for copy operations, this
command requires an input file.  It cannot operate on standard input.
The reduced dump is emitted to standard output.`,
		Run: func(cmd *cobra.Command, args []string) {},
	}

	seeCommand := &cobra.Command{
		Use:   `see`,
		Short: `Report only essential topological information`,
		Long: `see: usage: repocutter [-r SELECTION] see

Render a very condensed report on the repository node structure, mainly
useful for examining strange and pathological repositories,  File content
is ignored.  You get one line per repository operation, reporting the
revision, operation type, file path, and the copy source (if any).
Directory paths are distinguished by a trailing slash.  The 'copy'
operation is really an 'add' with a directory source and target;
the display name is changed to make them easier to see.`,
		Run: func(cmd *cobra.Command, args []string) {},
	}

	root := &cobra.Command{
		Use:   `repocutter [-q] -[r SELECTION] SUBCOMMAND`,
		Short: ``,
		Long: `repocutter - stream surgery on SVN dump files
general usage: repocutter [-q] [-r SELECTION] SUBCOMMAND

In all commands, the -r (or --range) option limits the selection of revisions
over which an operation will be performed. A selection consists of
one or more comma-separated ranges. A range may consist of an integer
revision number or the special name HEAD for the head revision. Or it
may be a colon-separated pair of integers, or an integer followed by a
colon followed by HEAD.

Normally, each subcommand produces a progress spinner on standard
error; each turn means another revision has been filtered. The -q (or
--quiet) option suppresses this.

Type 'repocutter help <subcommand>' for help on a specific subcommand.

Available subcommands:
   squash
   select
   propdel
   propset
   proprename
   log
   setlog
   strip
   expunge
   sift
   pathrename
   renumber
   reduce
   see
`,
		Run: func(cmd *cobra.Command, args []string) {},
	}

	root.AddCommand(squash, select_, propdel, propset, proprename, log, setlog,
		strip, expunge, sift, pathrename, renumber, reduceCommand, seeCommand)

	return root
}

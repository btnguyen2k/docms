#!/bin/sh

## Utility script to release project with a tag
## Usage:
##   ./release.sh <tag-name>

usage() {
	echo "Usage: $0 <component> <tag-name>"
	echo "where <component> is one of:"
	echo "	cli     - DO CMS CLI tool"
	echo "	runtime - DO CMS runtime"
	echo "	all     - all components"
	exit -1
}

release() {
	component="$1"
	tag="$2"
	echo "Tagging $component-$tag..."
	git tag -f -a "$component-$tag" -m "$component-$tag"
	git push origin "$component-$tag" -f
}

releaseCli() {
	release cli "$1"
}

releaseRuntime() {
	release rt "$1"
}

releaseAll() {
	tag="$1"
	echo "Tagging $tag..."
	git tag -f -a "$tag" -m "$tag"
	git push origin "$tag" -f
}

if [ "$#" -ne 2 ]; then
	usage
fi

case "$1" in
	"cli")
		releaseCli "$2"
		;;
	"runtime"|"rt")
		releaseRuntime "$2"
		;;
	"all")
		releaseAll "$2"
		;;
	*)
		usage
		;;
esac

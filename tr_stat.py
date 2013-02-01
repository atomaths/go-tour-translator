#!/usr/bin/python

import os
import sys

def extractTitles(src):
    tmpFn = '/tmp/go-tour.titles'
    cmdFmt = 'grep "^\* " %s | nl -s ":" > %s'
    os.system(cmdFmt%(src, tmpFn))
    return open(tmpFn).readlines()


if __name__ == '__main__':
    enTitles = extractTitles('tour.article.orig')
    koTitles = extractTitles('tour.article')

    if len(koTitles) != len(enTitles):
        print "Something wrong?!"
        sys.exit(0)

    notTranslatedCnt = 0
    for i in range(len(koTitles)):
        tKo = koTitles[i]
        tEn = enTitles[i]

        if tKo == tEn:
            notTranslatedCnt += 1
            print tEn.rstrip()

    print "="*79
    print "%d%% translated"%(notTranslatedCnt * 100 / len(enTitles))

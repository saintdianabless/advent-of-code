#ifndef _UTILS_H_
#define _UTILS_H_

#include <vector>
#include <string>
#include <fstream>
#include <ios>
#include <stdio.h>

#include "folly/String.h"
#include "folly/FBVector.h"

using folly::fbvector;
using folly::fbstring;

fbvector<fbstring> readLines(folly::StringPiece filename)
{
    std::ifstream infile(filename.data(), std::ios::in);
    if(!infile.is_open()) {
        printf("error when open file.\n");
        exit(1);
    }
    fbvector<fbstring> lines;
    for (fbstring line; getline(infile, line);)
    {
        lines.emplace_back(std::move(line));
    }
    return lines;
}

fbvector<folly::StringPiece> split(folly::StringPiece s, folly::StringPiece delim) {
    fbvector<folly::StringPiece> v;
    folly::split(delim, s, v);
    return v;
}

#endif
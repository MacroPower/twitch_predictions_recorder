g++ -fPIC -lm -std=c++11 -shared \
    jaroWinkler.cpp pylcs.cpp dldist.cpp lcsubstr.cpp perm.cpp subseq.cpp RegistExt.cpp utf8_unicode.cpp \
    -o distlib_64.so
mv distlib_64.so distlib_64.dylib

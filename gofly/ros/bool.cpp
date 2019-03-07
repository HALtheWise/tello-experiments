#include "bool.h"

const Bool openCVVersion() {
    auto b = Bool();
    b.data = true;
    return b;
}

const bool Bool_data(Bool b){
    return b.data;
}
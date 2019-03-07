#include "vins.h"

Estimator* NewEstimator() {
    return new Estimator;
}

bool EstimatorInitialStructure(Estimator* e){
    return e->initialStructure();
}
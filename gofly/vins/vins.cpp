#include "vins.h"

Estimator* NewEstimator() {
    return new Estimator;
}

bool EstimatorInitialStructure(){
    readParameters(std::string());
    return Estimator::INITIAL;
}
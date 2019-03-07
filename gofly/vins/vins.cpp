#include "vins.h"

// Estimator estimator;

// Estimator NewEstimator() {
//     return estimator;
// }

bool EstimatorInitialStructure(){
    readParameters(std::string());
    return Estimator::INITIAL;
}
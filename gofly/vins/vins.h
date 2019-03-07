#include <stdbool.h>

#ifdef __cplusplus
#include <../../third_party/vins-fusion/vins_estimator/src/estimator/estimator.h>

extern "C" {
#endif

#ifdef __cplusplus
// typedef Estimator Estimator;
#else
typedef void* Estimator;
#endif

Estimator* NewEstimator();
bool EstimatorInitialStructure(Estimator* e);

#ifdef __cplusplus
}
#endif
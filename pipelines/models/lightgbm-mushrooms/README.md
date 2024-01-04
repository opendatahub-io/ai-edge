# Mushroom model
## Description
This model classifies mushrooms based on the input image.

## Dataset

This model was trained using the [Mushroom Classification](https://www.kaggle.com/datasets/uciml/mushroom-classification) dataset.
## Test
### Run Seldon MLServer with Mushrooms model
Execute the following command from the [pipelines](../../) folder:
```
docker run -d -u $(id -u) --rm -v ${PWD}/models:/opt/models:Z -p 8080:8080 -p 8081:8081 -p 8082:8082 -ti seldonio/mlserver:1.3.5-lightgbm mlserver start /opt/models/lightgbm-mushrooms
```
### Test call

```
curl -s -d '{"inputs": [{"name": "predict", "shape": [1, 126], "datatype": "FP32", "data": [[1.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 1.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 1.0, 0.0, 0.0, 1.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 1.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0]]}]}' -H 'Content-Type: application/json' -X POST http://localhost:8080/v2/models/mushroom-lgb/versions/v0.1.0/infer
```
## Credits
 - https://www.kaggle.com/code/stpeteishii/mushroom-predict-and-visualize-importance
# Iris model
## Description
This model classifies Iris flowers in order to distinguish between the three Irish species (Iris setosa, Iris virginica and Iris versicolor).

## Dataset

This model was trained using the well known [Iris](https://archive.ics.uci.edu/dataset/53/iris) dataset.

## Test
### Run Seldon MLServer with IRIS model
Execute the following command from the [examples](../../) folder:
```
podman run -d --rm -v ${PWD}/models:/opt/models:Z -p 8080:8080 -p 8081:8081 -p 8082:8082 -ti seldonio/mlserver:1.3.5-lightgbm mlserver start /opt/models/lightgbm-iris
```

### Test call
```
curl -s -d '{"inputs": [{"name": "predict-prob", "shape": [1, 4], "datatype": "FP32", "data": [[4.9, 3.0, 1.4, 0.2]]}]}' -H 'Content-Type: application/json' -X POST http://localhost:8080/v2/models/iris/versions/v0.1.0/infer
```

## Credits
- https://www.kaggle.com/code/ajsherlock/iris-classification-using-lightgbm
- https://github.com/SeldonIO/MLServer/blob/master/docs/examples/lightgbm/README.md

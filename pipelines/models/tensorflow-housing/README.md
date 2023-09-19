# How was this model generated?

## Data source

The model uses
[scikit-learn California Housing Dataset](https://scikit-learn.org/stable/datasets/real_world.html#california-housing-dataset),
see [this course notebook](https://inria.github.io/scikit-learn-mooc/python_scripts/datasets_california_housing.html)
for more information about the data.

## Generating model

A command similar to the following was used to generate the model:
```
mlflow run 'https://github.com/mlflow/mlflow#examples/tensorflow'
```
It used the training script
https://github.com/mlflow/mlflow/blob/master/examples/tensorflow/train.py
present in that repo.

Running it on a checkout of the repository also works:
```
mlflow run mlflow/examples/tensorflow
```

The result gets stored in a `mlruns/0/<run-id>/artifacts/model/` directory,
and was copied over to this repository.

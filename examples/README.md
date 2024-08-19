# Pipelines Examples

This directory contains resources used to successfully run pipelines and tasks from the [manifests](../manifests) directory
using the [Pipelines Setup](../manifests/README.md) guide.

## Container files

These files are used for building images with models during the pipeline run.

## Models

This directory contains the following trained examples.

- [bike-rentals-auto-ml](models/bike-rentals-auto-ml/) is using MLFlow format and can run in [Seldon MLServer](https://github.com/SeldonIO/MLServer).
- [tensorflow-housing](models/tensorflow-housing/) is using MLFlow format and wraps a TensorFlow model. It can run in [Seldon MLServer](https://github.com/SeldonIO/MLServer), but can also run in [OVMS](https://github.com/openvinotoolkit/model_server) by loading the [tf2model](models/tensorflow-housing/tf2model) artifacts.
- [MNIST](models/onnx-mnist) is using ONNX format that can run on [OVMS](https://github.com/openvinotoolkit/model_server).
- [Face Detection](models/tensorflow-facedetection) is using OpenVino IR propietary format and would run only on  [OVMS](https://github.com/openvinotoolkit/model_server).
- [Iris](models/lightgbm-iris) is using Booster format which can run on [Seldon MLServer](https://github.com/SeldonIO/MLServer).
- [Mushrooms](models/lightgbm-mushrooms) is using Booster format which can run on [Seldon MLServer](https://github.com/SeldonIO/MLServer).

```plaintext
bike-rentals-auto-ml/
├── conda.yaml
├── MLmodel
├── model.pkl
├── python_env.yaml
└── requirements.txt

tensorflow-housing/
├── conda.yaml
├── MLmodel
├── model.pkl
├── python_env.yaml
├── requirements.txt
└── tf2model/
    ├── saved_model.pb
    └── ...

onnx-mnist/
├── 1
│   ├── mnist.onnx
│   └── schema
│       └── schema.json
└── README.md

tensorflow-facedetection/
├── 1
│   ├── face-detection-retail-0004.bin
│   └── face-detection-retail-0004.xml
└── README.md

lightgbm-iris/
├── iris-lightgbm.bst
├── model-settings.json
├── README.md
└── settings.json

lightgbm-mushrooms/
├── model-settings.json
├── mushroom-lightgbm.bst
├── README.md
└── settings.json

```

 ## Tekton
This directory contains example Tekton Pipeline Runs, templates for secrets and test data.
